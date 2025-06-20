package utils

import (
	"log"
	"os/exec"
	"strings"
)

func GetUUID() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	// Remove trailing newline
	return strings.TrimSpace(string(newUUID))
}
