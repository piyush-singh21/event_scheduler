package auth

import (
	"errors"
	"event_scheduler/database"
	"event_scheduler/model"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	UserID int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(User model.User) (string, error) {
	expirationTime := time.Now().Add(2 * time.Minute) //Token expires in 2 minutes
	claims := &Claims{
		UserID: User.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   strconv.Itoa(User.ID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
func verifyUser(enteredPass, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(enteredPass))
}
func AuthenticateUser(email, password string) (string, error) {
	var user model.User
	err := database.DB.QueryRow("SELECT id,password FROM users WHERE email=?", email).Scan(&user.ID, &user.Password)

	if err != nil {
		return "", errors.New("email not found,please register")
	}
	if err = verifyUser(password, user.Password); err != nil {
		return "", err
	}
	token, err := GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
