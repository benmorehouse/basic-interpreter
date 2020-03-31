package main

import (
	"golang.org/x/crypto/bcrypt"
)

/*
NOTE:
	In order to hash for each user and each of their files, i want to
	create filenames which are filepath_username, and then get a number hash for an id based on that.
*/

// HashFileID will return a full fileID for the sql database
func (a *App) HashFileName(filename string) []byte {

	h := sha1.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

// GetFileFromFilestore will get a file given a hash from the database.
// the ~filestore~ as you will :)
func (d *DBcxn) GetFileFromFilestore(fileHash []byte) (*File, error) {

	return nil, nil
}
