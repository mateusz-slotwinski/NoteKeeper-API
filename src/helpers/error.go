package helpers

import (
	context "context"
	log "log"

	Config "NoteKeeperAPI/src/config"
)

func PrintError(err error) {
	if Config.Mode == "DEV" && err != nil {
		log.Print(err)
	}
}

func PrintCancel(err context.CancelFunc) {
	if Config.Mode == "DEV" && err != nil {
		log.Fatal(err)
	}
}
