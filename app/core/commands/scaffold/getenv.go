package scaffold

import (
	"log"
	"os"
)

func GetGOPATH() string {
	GOPATH := os.Getenv("GOPATH")
	if GOPATH == "" {
		log.Fatal("Oops, GOPATH is empty, have you set your GOPATH env variable?")
	}
	return GOPATH
}

// GetPKGPath returns the template project path
func GetPKGPATH() string {
	return "github.com/fernandochristyanto/retsgo"
}

func GetVersion() string {
	return "1.1.0"
}
