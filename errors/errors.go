package errors

import "errors"

var ErrCannotFindZipcode = errors.New("can not find zipcode")
var ErrWeatherNotFound = errors.New("weather not found")
