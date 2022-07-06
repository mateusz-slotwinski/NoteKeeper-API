package src

import (
	Config "NoteKeeperAPI/src/config"
	fmt "fmt"
	runtime "runtime"
)

var Logs = func() {
	if Config.Mode == "DEV" {
		fmt.Printf("\x1b[4m\n___________________________________________________\x1b[0m\n\n")
		fmt.Printf("\x1b[36mNoteKeeper v2.0.1 \x1b[0m\n")
		fmt.Printf("\x1b[32mGo Version: %v \x1b[0m \n", runtime.Version())
		fmt.Printf("by: Mateusz Słotwiński\n\n")
		fmt.Printf("App running at: \x1b[36mhttp://localhost%v/\x1b[0m", Config.Port)
		fmt.Printf("\x1b[4m\n___________________________________________________\x1b[0m\n\n\n")
	}
}
