package repositorys

import (
	"API/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) error {
	statement, err := repository.db.Prepare(
		"INSERT INTO USERS (name, nick, email, password) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repository Users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := repository.db.Query(
		"SELECT ID, NAME, NICK, EMAIL, CREATEDATE FROM USERS WHERE NAME LIKE ? OR NICK LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateDate,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) GetById(userId uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT ID, NAME, NICK, EMAIL, CREATEDATE FROM USERS WHERE id = ?",
		userId,
	)
	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()
	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateDate,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE USERS SET NAME = ?, NICK = ?, EMAIL = ? WHERE ID = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM USERS WHERE ID = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository Users) GetByEmail(email string) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT ID, PASSWORD FROM USERS WHERE EMAIL = ?",
		email,
	)
	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) GetPassword(userId uint64) (string, error) {
	lines, err := repository.db.Query("SELECT PASSWORD FROM USERS WHERE ID = ?", userId)
	if err != nil {
		return "", err
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository Users) ChangePassword(userId uint64, password string) error {
	statement, err := repository.db.Prepare("UPDATE USERS SET PASSWORD = ? WHERE ID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userId); err != nil {
		return err
	}

	return nil
}
