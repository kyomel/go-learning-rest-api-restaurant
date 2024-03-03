package rest

import (
	"errors"
	"net/http"
	"rest-api-restaurant/internal/model"
	"strings"
)

func GetSessionData(r *http.Request) (model.UserSession, error) {
	authString := r.Header.Get("Authorization")
	splitString := strings.Split(authString, " ")
	if len(splitString) != 2 {
		return model.UserSession{}, errors.New("unauthorized")
	}

	accessString := splitString[1]
	return model.UserSession{
		JWTToken: accessString,
	}, nil
}
