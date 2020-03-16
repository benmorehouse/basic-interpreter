package bos

import (
	"sort"
	errors "github.com/benmorehouse/basicInterpreter/err"
	gen "github.com/benmorehouse/basicInterpreter/err"
	log "github.com/sirupsen/logrus"
)

// Directory is a struct that will hold all files and subdirs in a single interface.
//
type Directory struct {
	Name           string
	SubDirectories map[string]*Directory
	SubFiles       map[string]*File
}

// List is a method which lists of contents of a directory.
func (d Directory) List() []string {
	if d == nil {
		return nil
	}

	var arrayOfNames []string
	for name := range d.SubDirectories {
		arrayOfNames = append(arrayOfNames, name)
	}

	for name := range d.SubFiles {
		arrayOfNames = append(arrayOfNames, name)
	}

	sort.Strings(arrayOfNames)
	return arrayOfNames
}

func (d Directory) AddDirectory(subDir *Directory) error
	if d == nil {
		return nil
	}

	if subDir == nil {
				
	}

	return nil
}

// File is a struct which contains the name as well as the method to get the file contents from the database.
type File struct {
	Name string
}
