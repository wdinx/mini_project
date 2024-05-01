package constant

import "errors"

var ErrInsertDatabase error = errors.New("Invalid Add Data in Database")
var ErrEmptyInput error = errors.New("name, email and password cannot be empty")
var ErrDataNotFound error = errors.New("Data not found in database")
var ErrLoginFailed error = errors.New("Login Failed")
