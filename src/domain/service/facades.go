package service

import "github.com/alaa-aqeel/govalid/src/pkgs/database"

func User() *UserService {
	return &UserService{
		db: database.DB,
	}
}
