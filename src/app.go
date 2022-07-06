package src

import (
	Config "NoteKeeperAPI/src/config"
	Router "NoteKeeperAPI/src/router"
)

func App() {

	App := Router.CreateServer()

	App.Run(Config.Port)
}

func Init() {
	Config.Init()

	Logs()
	App()

}
