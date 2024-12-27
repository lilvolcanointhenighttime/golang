package modules

import (
	"errors"

)

type AuthRequest struct {
	Email     string     `json:"email"`
	Password  string     `json:"password"` 
}

func Authentication(password1, password2 string) (bool, error) {
	if password1 != password2 {
		return false, errors.New("password is incorect")
	}
	return true, nil
}

