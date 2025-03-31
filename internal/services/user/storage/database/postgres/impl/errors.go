package impl

import "errors"

var UserNotFoundError = errors.New("user not found")
var UserAlreadyExistsError = errors.New("user already exists")
var IncorrectPasswordError = errors.New("incorrect password")
