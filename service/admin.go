package service

import (
	"golang.org/x/crypto/bcrypt"
	"zealthy-helpdesk-backend/dao"
)

func CheckAdminExists(email string) (bool, error) {
	SQL := `SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)`
	args := []any{email}
	var exists bool
	err := dao.DB.Get(&exists, SQL, args...)
	return exists, err
}

func CheckAdminPassword(email, password string) (bool, error) {
	SQL := `SELECT password FROM admins WHERE email = $1`
	args := []any{email}
	var passwordHash string
	err := dao.DB.Get(&passwordHash, SQL, args...)
	if err != nil {
		return false, err
	}
	return doPasswordMatch(password, passwordHash), nil
}

func doPasswordMatch(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
