package auth

import (
	"errors"
	"event_scheduler/database"
	"event_scheduler/model"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	UserID int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(User model.User) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour) //Token expires in 2 hours
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
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unable to login")
	}
	if !token.Valid {
		return nil, errors.New("unable to login")
	}
	return claims, nil
}

//	func verifyUser(enteredPass, hashedPass string) error {
//		return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(enteredPass))
//	}
func AuthenticateUser(email, password string) (string, error) {
	var user model.User
	err := database.DB.QueryRow("SELECT id,password FROM users WHERE email=?", email).Scan(&user.ID, &user.Password)

	if err != nil {
		return "", errors.New("email not found,please register")
	}
	if password != user.Password {
		return "", errors.New("unable to login")
	}
	token, err := GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
