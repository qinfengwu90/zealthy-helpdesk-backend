package service

import (
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
	"zealthy-helpdesk-backend/dao"
)

func CheckAdminExists(email string) (bool, error) {
	return dao.CheckAdminExists(email)
}

func CheckAdminPassword(email, password string) (bool, error) {
	passwordHash, err := dao.GetAdminPasswordHash(email)
	if err != nil {
		return false, err
	}
	return doPasswordMatch(password, passwordHash), nil
}

func doPasswordMatch(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

func CreateAdmin(email string, password string, firstName null.String, lastName null.String) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return dao.CreateAdmin(email, passwordHash, firstName, lastName)
}

func ChangeAdminPassword(email, newPassword string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return dao.ChangeAdminPassword(email, passwordHash)
}
