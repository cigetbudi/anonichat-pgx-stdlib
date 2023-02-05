package repos

import (
	"anonichat-pgx-stdlib/db"
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/utils"
	"errors"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(u *models.User) error {
	defer utils.Timer(time.Now(), "register")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Phone = html.EscapeString(strings.TrimSpace(u.Phone))
	u.Phone = strings.Replace(u.Phone, "-", "", -1)
	dob, err := time.Parse("2006-01-02", u.DOB)
	if err != nil {
		return err
	}

	res, err := db.DBN.Exec("INSERT INTO USERS (username,fullname,password,email,created_at,dob, phone, gender_code) values($1,$2,$3,$4,$5,$6,$7,$8)", u.Username, u.Fullname, hashedPassword, u.Email, time.Now(), dob, u.Phone, u.GenderCode)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row != 1 {
		return errors.New("data terinsert ada lebih dari 1")
	}
	return nil

}

func VerifyPassword(pass, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}

func AddLastLogin(username string) error {
	defer utils.Timer(time.Now(), "Add Last Login")

	res, err := db.DBN.Exec("UPDATE users SET last_login = $1 , WHERE username = $2", time.Now(), username)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row != 1 {
		return errors.New("data terinsert ada lebih dari 1")
	}
	return nil
}

func CheckLoginAttemp(username string) (uint, error) {
	defer utils.Timer(time.Now(), "Check Login Attempt")
	var logCount uint

	err := db.DBN.QueryRow("SELECT login_attempt FROM users where username = $1 ", username).Scan(&logCount)
	if err != nil {
		return 4, err
	}
	return logCount, nil
}

func LoginCheck(username, password string) (string, error) {
	defer utils.Timer(time.Now(), "Check Login by ID")
	var (
		un  string
		pa  string
		err error
		id  uuid.UUID
	)

	err = db.DBN.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&id, &un, &pa)
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, pa)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(id)
	if err != nil {
		return "", err
	}
	defer utils.Timer(time.Now(), "Update Last Login")

	_, err = db.DBN.Exec("UPDATE users SET last_login = $1 WHERE id = $2 ", time.Now(), id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func AddLoginAttemp(username string) error {
	defer utils.Timer(time.Now(), "Add Login Attempt")

	_, err := db.DBN.Exec("UPDATE users SET login_attempt = login_attempt + 1 WHERE username = $1", username)
	if err != nil {
		return err
	}
	return nil
}
