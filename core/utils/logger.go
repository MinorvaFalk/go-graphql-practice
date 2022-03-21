package utils

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	l := log.New(os.Stdout, "[Book API] ", log.LstdFlags)
	return l
}
