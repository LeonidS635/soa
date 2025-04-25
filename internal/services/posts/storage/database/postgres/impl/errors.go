package impl

import "errors"

var PostAccessError = errors.New("operation with post is available only for author")
var PostNotFoundError = errors.New("post not found")
