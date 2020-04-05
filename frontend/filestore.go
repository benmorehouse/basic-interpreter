package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

/*
NOTE:
	In order to hash for each user and each of their files, i want to
	create filenames which are filepath_username, and then get a number hash for an id based on that.
*/

// HashFileID will return a full fileID for the sql database
// the fileID will be in the form of <filepath>#filename
func ValidateFileName(filename, filepath string) ([]byte, error) {

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

	fs = append(fs, byte('#'))

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

	return fs, nil
}

// GetFileFromFilestore will get a file given a hash from the database.
// the ~filestore~ as you will :)
func (d *DBcxn) GetFileFromFilestore(fileHash []byte, currentUserID string) (*File, error) {

	if err := d.PingContext(); err != nil {
		log.Error(err)
		return nil, err
	}

	query := fmt.Sprintf(
		"select * from %s where id=%s;",
		d.FileTable,
		string(fileHash),
	)

	row := d.cxn.QueryRowContext(*d.context, query)

	var fileBuffer, userID string
	var contentBuffer []byte

	row.Scan(
		&fileBuffer,
		&userID,
		&contentBuffer,
	)

	filepath, filename := "", ""
	cursor := 0
	for key, value := range fileBuffer {
		if value == '#' {
			cursor = key + 1
			break
		}

		filepath += string(value)
	}

	if cursor > len(fileBuffer)-1 {
		return nil, PostgresError(FileIsNil, nil)
	}

	for _, value := range fileBuffer[cursor:] {
		filename += string(value)
	}

	newFile := &File{
		Name: filename,
		Path: filepath,
	}

	newFile.CheckIfBasicFile()
	newFile.ReadFileFromFilestore(contentBuffer)

	return nil, nil
}

// Save is used to save a file into the database
func (d *DBcxn) Save(file *File, userId string) error {

	if err := d.PingContext(); err != nil {
		return err
	}

	if file == nil {
		err := PostgresError(FileIsNil, nil)
		return err
	}

	if exists, err := d.FileAlreadyExists(file, userId); err != nil {
		log.Error(err)
		return err
	} else if exists {
		err := NewFileStoreError(FileAlreadyExists, nil)
		log.Error(err)
		return err
	}

	fileHash, err := ValidateFileName(file.Path, file.Name)
	if err != nil {
		log.Error(err)
		return err
	}

	//fileHash is the ida
	query := `
		insert into ? (
			id,
			userid, 
			file
		) values (
			?,
			?,
			?	
		);	
	`

	if _, err = d.cxn.ExecContext(
		*d.context,
		query,
		d.FileTable,
		fileHash,
		userId,
		file.WriteFileForSaving(),
	); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// FileAlreadyExists will tell whether or not a file exists in the database already
func (d *DBcxn) FileAlreadyExists(file *File, userId string) (bool, error) {

	if err := d.PingContext(); err != nil {
		return true, err
	}

	if file == nil {
		err := PostgresError(FileIsNil, nil)
		return true, err
	}

	fileHash, err := ValidateFileName(file.Path, file.Name)
	if err != nil {
		log.Error(err)
		return true, err
	}

	query := `
		select * from ?
		where id=?
		and userid=?;
	`

	if result := d.cxn.QueryRowContext(
		*d.context,
		query,
		d.FileTable,
		fileHash,
		userId,
	); result != nil {
		log.Error(NewFileStoreError(FileAlreadyExists, nil))
		return true, nil
	}

	return false, nil
}
