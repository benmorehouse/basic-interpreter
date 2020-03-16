package bos

import (
	log "github.com/sirupsen/logrus"
)

type OperatingSystem struct {
	CommandMap       map[string]int
	CurrentDirectory *Directory
}
