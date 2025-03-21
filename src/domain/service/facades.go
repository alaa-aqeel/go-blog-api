package service

import "github.com/alaa-aqeel/govalid/src/database"

func User() *UserService {
	return &UserService{
		db: database.DB,
	}
}
