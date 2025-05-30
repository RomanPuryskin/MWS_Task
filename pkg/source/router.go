package source

import (
	"os"
)

// возможные команды
const (
	PROFILE = "profile"
	HELP    = "help"
)

// возможные подкоманды
const (
	CREATE = "create"
	GET    = "get"
	LIST   = "list"
	DELETE = "delete"
)

func RouterCommands(args []string) {

	if len(args) < 2 {
		logUsage()
		os.Exit(1)
	}

	command := args[1]

	switch command {
	case PROFILE:
		profileCommandHandler(args[1:])
	case HELP:
		helpCommandHandler(args[1:])
	default:
		logUnknownCommandError(command)
		os.Exit(1)
	}

}
