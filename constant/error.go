package constant

import "errors"

var ErrInsertDatabase error = errors.New("Invalid Add Data")
var ErrEmptyInput error = errors.New("field cannot be empty")
var ErrDataNotFound error = errors.New("Data not found")
var ErrLoginFailed error = errors.New("Login Failed")
var ErrInputData error = errors.New("Invalid Input Data")
