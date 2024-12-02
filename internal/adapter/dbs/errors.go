package dbs

import "errors"

var (
	ErrorRecordNotFound      = errors.New("record not found")
	ErrorRecordAlreadyExists = errors.New("record already exist")
)
