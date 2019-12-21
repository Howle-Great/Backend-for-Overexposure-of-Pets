package main

import "database/sql"

type ItemRepository struct {
	DB *sql.DB
}

func (repo *ItemRepository) GetUsersPassword(nicknameOrMail string) (string, error) {
	var hashedPassword string
	err := repo.DB.
		QueryRow("select password from items where nickname = $1 OR mail = $1", nicknameOrMail).
		Scan(&hashedPassword)

	if err != nil {
		return "", err
	}

	return hashedPassword, nil
}
