package main

import (
	log "github.com/sirupsen/logrus"
)

// HashFileID will return a full fileID for the sql database
// the fileID will be in the form of <filepath>#filename
func ValidateFileName(filename, filepath string) ([]byte, error) {
	log.Info(filename)
	log.Info(filepath)
	fs := []byte{}
	for _, char := range filename {
		switch char {
		case '#', '&', ':', ' ':
			err := NewFileStoreError(InvalidFileName, nil)
			log.Error(err)
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
			log.Error(err)
			return nil, err

		default:
			fs = append(fs, byte(char))
		}
	}

	log.Warning(string(fs))
	log.Warning(string(fs))
	log.Warning(string(fs))
	return fs, nil
}

// GetFileFromFilestore will get a file given a hash from the database.
// the ~filestore~ as you will :)
func (d *DBcxn) GetFileFromFilestore(fileHash []byte, currentUserID string) (*File, error) {

	if err := d.PingContext(); err != nil {
		log.Error(err)
		return nil, err
	}

	query := `
		select * from ?
		where id=?
		and userid=?;
	`

	row := d.cxn.QueryRowContext(
		*d.context,
		query,
		d.FileTable,
		string(fileHash),
		currentUserID,
	)

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

	return newFile, nil
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

	fileHash, err := ValidateFileName(file.Path, file.Name)
	if err != nil {
		log.Error(err)
		return err
	}

	if exists, err := d.FileAlreadyExists(fileHash, userId); err != nil {
		log.Error(err)
		return err
	} else if exists {
		log.Debug("Overwriting a user file...")
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
func (d *DBcxn) FileAlreadyExists(fileHash []byte, userId string) (bool, error) {

	if err := d.PingContext(); err != nil {
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
		string(fileHash),
		userId,
	); result != nil {
		log.Error(NewFileStoreError(FileAlreadyExists, nil))
		return true, nil
	}

	return false, nil
}
