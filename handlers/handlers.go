package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
	"go.auth.playground/models"
)

// PENDING - Jangan lupa taruh di config~
var secret = []byte(os.Getenv("SERVICE_SECRET"))

// Auth ...
type Auth struct {
	UsersRepository models.UsersAdapter
}

// Index ...
type Index struct {
	UsersRepository models.UsersAdapter
}

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type indexRequest struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// Home ...
func (a *Auth) Home(w http.ResponseWriter, r *http.Request) {
	indexRequest := &indexRequest{}
	err := json.NewDecoder(r.Body).Decode(indexRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	claims := &claims{}

	token, err := jwt.ParseWithClaims(indexRequest.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil && err.Error() != "signature is invalid" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var res []byte
	if (err != nil && err.Error() == "signature is invalid") || (indexRequest.Username != claims.Username) || !token.Valid {
		res, err = json.Marshal(map[string]string{"message": "You aren't authorized."})
	} else {
		res, err = json.Marshal(map[string]string{"message": "You are authorized."})
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// SignUp ...
func (a *Auth) SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.Users{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	user.Password = string(hashedPassword)
	a.UsersRepository.Create(*user)
}

// SignIn ...
func (a *Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	requestedUser := &models.Users{}
	err := json.NewDecoder(r.Body).Decode(requestedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	userFromDatabase, err := a.UsersRepository.GetByName(requestedUser.Username)
	err = bcrypt.CompareHashAndPassword([]byte(userFromDatabase.Password), []byte(requestedUser.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	willBeExpiredAt := time.Now().Add(60 * time.Minute)

	claims := &claims{
		Username: userFromDatabase.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: willBeExpiredAt.Unix(),
		},
	}

	signature := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := signature.SignedString(secret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	res, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
