package handlers

import "errors"

var JWTError = errors.New("JWT error")
var ClaimsError = errors.New("claims in JWT error")
var SigningError = errors.New("error signing JWT")
var ReadBodyError = errors.New("error reading request body")
