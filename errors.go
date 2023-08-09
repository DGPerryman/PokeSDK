package pokesdk

import "errors"

// ErrNotFound: returned when an item was requested that was not found in the PokeAPI database
var ErrNotFound = errors.New("not found")

// ErrReadFailed: returned when reading from PokeAPI failed
var ErrReadFailed = errors.New("read failed")

// ErrUnexpectedResponse: returned when the PokeAPI response cound not be processed
var ErrUnexpectedResponse = errors.New("unexpected response from PokeAPI")
