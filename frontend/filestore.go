package main

import (
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

/*
NOTE:
	In order to hash for each user and each of their files, i want to
	create filenames which are filepath_username, and then get a number hash for an id based on that.
*/

// HashFileID will return a full fileID for the sql database
// the fileID will be in the form of <filepath>#filename
func (a *App) HashFileName(filename, filepath string) ([]byte, error) {

	fs := []byte{}
	for _, char := range filename {
		switch char {
		case '#', '&', ':', ' ':
			err := NewFileStoreError(InvalidFileName, nil)
			log.Error(err)
			return nil, err

		default:
			fs = append(fs, byte(char))
		}
	}

	for _, char := range filepath {
		switch char {
		case '#', '&', ':', ' ':
			err := NewFileStoreError(InvalidFileName, nil)
			log.Error(err)
			return nil, err

		default:
			fs = append(fs, byte(char))
		}
	}

	newFileStore, err := bcrypt.GenerateFromPassword(fs, 10)
	if err != nil {
		err := NewFileStoreError(FileTranslationFailed, err)
		log.Error(err)
		return nil, err
	}

	return newFileStore, nil
}

// GetFileFromFilestore will get a file given a hash from the database.
// the ~filestore~ as you will :)
func (d *DBcxn) GetFileFromFilestore(fileHash []byte) (*File, error) {

	return nil, nil
}

// #################################################
// ############## sql ##############################

// NOTE: now would be a great time to have implemented an error interface. Then we just call
// a.Save(whatever) all over the place.
func (a *App) Save(interface{}) error {

}

func (d *DBcxn) Save(file *File) error {

	if err := d.PingContext(); err != nil {
		return err
	}

	if file == nil {
		err := PostgresError(FileIsNil, nil)
		return err
	}

}
