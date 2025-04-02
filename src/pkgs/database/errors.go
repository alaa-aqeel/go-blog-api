package database

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// DBError represents a custom database error type
type DBError struct {
	Message   string
	Original  error
	ErrorCode string
}

func (e *DBError) Error() string {
	return e.Message
}

func (e *DBError) Unwrap() error {
	return e.Original
}

func NewDBError(message string, original error, code string) *DBError {
	return &DBError{
		Message:   message,
		Original:  original,
		ErrorCode: code,
	}
}

// HandleSqlError handles all types of database errors
func HandleSqlError(err error) error {
	if err == nil {
		return nil
	}

	// Handle GORM errors
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewDBError("The requested record was not found.", err, "NOT_FOUND")
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return NewDBError("This value already exists and must be unique.", err, "DUPLICATE_KEY")
	}

	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return NewDBError("The database transaction is invalid.", err, "INVALID_TX")
	}

	// Handle PostgreSQL errors
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			column := extractColumnFromConstraint(pgErr.ConstraintName)
			msg := "This data already exists."
			if column != "" {
				msg = fmt.Sprintf("The value for '%s.%s' already exists.", pgErr.TableName, column)
			}

			return NewDBError(msg, err, pgErr.Code)

		case pgerrcode.ForeignKeyViolation:
			return NewDBError("This action violates a relationship rule in the database.", err, pgErr.Code)

		case pgerrcode.NotNullViolation:
			column := extractColumnFromMessage(pgErr.Message)
			msg := "A required field is missing."
			if column != "" {
				msg = fmt.Sprintf("The field '%s%s' is required.", pgErr.TableName, column)
			}
			return NewDBError(msg, err, pgErr.Code)

		case pgerrcode.CheckViolation:
			return NewDBError("The provided data does not meet the required conditions.", err, pgErr.Code)

		case pgerrcode.DeadlockDetected:
			return NewDBError("A database deadlock was detected. Please try again.", err, pgErr.Code)

		case pgerrcode.InsufficientResources:
			return NewDBError("The database has reached its resource limit. Please try later.", err, pgErr.Code)

		case pgerrcode.ConnectionException:
			return NewDBError("There was a problem connecting to the database.", err, pgErr.Code)

		default:
			return NewDBError(
				fmt.Sprintf("A database error occurred: %s (Error code: %s)", pgErr.Detail, pgErr.Code),
				err,
				pgErr.Code,
			)
		}
	}

	// Handle generic SQL errors
	if strings.Contains(err.Error(), "duplicate") {
		return NewDBError("This entry already exists.", err, "DUPLICATE_ENTRY")
	}

	if strings.Contains(err.Error(), "timeout") {
		return NewDBError("The database operation took too long and timed out.", err, "TIMEOUT")
	}

	if strings.Contains(err.Error(), "connection") {
		return NewDBError("There was an issue connecting to the database.", err, "CONNECTION_ERROR")
	}

	// Fallback for unhandled errors
	return NewDBError(fmt.Sprintf("An unexpected database error occurred: %v", err), err, "UNKNOWN_ERROR")
}

// Helper function to extract column name from constraint
func extractColumnFromConstraint(constraint string) string {
	if constraint == "" {
		return ""
	}

	// Common constraint patterns
	parts := strings.Split(constraint, "_")
	if len(parts) > 1 {
		// Handle patterns like: uix_users_email, idx_email, etc.
		lastPart := parts[len(parts)-1]
		if len(lastPart) > 0 {
			return lastPart
		}
	}

	return ""
}

// Helper function to extract column name from error message
func extractColumnFromMessage(message string) string {
	if message == "" {
		return ""
	}

	// PostgreSQL messages often contain: "null value in column \"column_name\" violates not-null constraint"
	start := strings.Index(message, "\"")
	if start == -1 {
		return ""
	}

	end := strings.Index(message[start+1:], "\"")
	if end == -1 {
		return ""
	}

	return message[start+1 : start+1+end]
}
