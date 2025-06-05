package utils

import (
	"log"
	"os/exec"
)

func GetUUID() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(newUUID)
}
