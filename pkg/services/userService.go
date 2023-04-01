package services

import (
	"ChatProgramming/pkg/errors"
	"ChatProgramming/pkg/models"
	"ChatProgramming/pkg/session"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"sync"
	"time"
)

type UserService struct {
	db    *sql.DB
	mutex sync.Mutex
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db:    db,
		mutex: sync.Mutex{},
	}
}

func (u *UserService) Register(username,
	firstName, lastName, password string) (string, error) {
	if u.checkUserExists(username) {
		return "", errors.ErrUserAlreadyExists
	}
	salt := session.RandomString(20)
	hashedPassword := shaPasswordHashing(password, salt)
	sqlStatement := `
		insert into users (username, first_name, last_name, password_hash,
		                          password_salt, registration_date)
		values ($1, $2, $3, $4, $5, $6)
	`
	_, err := u.db.Exec(sqlStatement, username, firstName, lastName,
		hashedPassword, salt, time.Now())
	if err != nil {
		log.Fatalf("Failed to save user info to db: %v", err)
		return "", err
	}
	return username, nil
}

func (u *UserService) Login(username string, password string) (string, error) {
	if !u.checkUserExists(username) {
		return "", errors.ErrUserNotExists
	}
	salt := u.getSaltForUser(username)
	currentHashedPassword := shaPasswordHashing(password, salt)
	correctPassword := u.getHashedPassword(username)

	if correctPassword == currentHashedPassword {
		return username, nil
	}
	return "", errors.ErrWrongPassword
}

func (u *UserService) GetInfo(username string) (*models.UserPublicInfo, error) {
	if !u.checkUserExists(username) {
		return nil, errors.ErrUserNotExists
	}
	var userInfo models.UserPublicInfo
	sqlStatement := `
		select username, first_name, last_name, registration_date from users where username = $1
	`

	err := u.db.QueryRow(sqlStatement, username).Scan(
		&userInfo.Username, &userInfo.FirstName, &userInfo.LastName, &userInfo.RegistrationDate)
	if err != nil {
		log.Fatalf("Failed to get user info by username %s: %v", username, err)
		return nil, err
	}
	return &userInfo, nil
}

func (u *UserService) checkUserExists(username string) bool {
	var userExists bool
	err := u.db.QueryRow(
		"select exists (select 1 from users where username = $1)",
		username,
	).Scan(&userExists)
	if err != nil {
		log.Fatalf("Failed to check if user with username %s exists: %v", username, err)
	}
	return userExists
}

func (u *UserService) getSaltForUser(username string) string {
	var passwordSalt string
	sqlStatement := `
		select password_salt from users where username = $1
	`
	err := u.db.QueryRow(sqlStatement, username).Scan(&passwordSalt)
	if err != nil {
		log.Fatalf("Failed to get password salt for user with username %s: %v", username, err)
		return ""
	}
	return passwordSalt
}

func (u *UserService) getHashedPassword(username string) string {
	var passwordHash string
	sqlStatement := `
		select password_hash from users where username = $1
	`

	err := u.db.QueryRow(sqlStatement, username).Scan(&passwordHash)
	if err != nil {
		log.Fatalf("Failed to get password hash for user with username %s: %v", username, err)
		return ""
	}
	return passwordHash
}

func shaPasswordHashing(password, salt string) string {
	passwordToHash := password + salt
	bytePassword := []byte(passwordToHash)
	shaHash := sha256.Sum256(bytePassword)
	return hex.EncodeToString(shaHash[:])
}
