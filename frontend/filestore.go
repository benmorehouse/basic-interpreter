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
func (a *App) ValidateFileName(filename, filepath string) ([]byte, error) {

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
	
	fs = append(fs, byte("#"))

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

	if err := d.PingContext(); err != nil {
		log.Error(err)
		return nil, err
	}
	
	query := fmt.Sprintf(
		"select * from %s where id=%s;",
		d.FileTable, 
		string(fileHash)
	)

	buffer, err := d.cxn.QueryRowContext(*d.context, query)
	if err != nil {
		log.Error(err)
		return err
	}

	filepath, filename := "", ""	
	cursor := 0
	for key, value := range buffer {
		if value == "#" {
			cursor = key + 1
			break
		}

		filepath += value
	}

	for _, value := range buffer[cursor:] {
		
	}

	newFile := &File {
			
	}
	
	return nil, nil
}

// Save is used to save a file into the database
func (d *DBcxn) Save(file *File) error {

	if err := d.PingContext(); err != nil {
		return err
	}

	if file == nil {
		err := PostgresError(FileIsNil, nil)
		return err
	}

	return nil
}
