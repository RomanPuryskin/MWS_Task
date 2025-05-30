package source

import (
	"flag"
	"os"
)

// возможные флаги
const (
	NAME_FLAG    = "name"
	USER_FLAG    = "user"
	PROJECT_FLAG = "project"
)

func profileCommandHandler(args []string) {
	if len(args) < 2 {
		logUsage()
		os.Exit(1)
	}

	subcommand := args[1]

	switch subcommand {
	case CREATE:
		createSubcommandHandler(args[1:])
	case GET:
		getSubcommandHandler(args[1:])
	case LIST:
		listSubcommandHandler(args[1:])
	case DELETE:
		deleteSubcommandHandler(args[1:])
	default:
		logUnknownCommandError(subcommand)
		os.Exit(1)
	}
}

func helpCommandHandler(args []string) {
	// для help не нужны никакие доп параметры
	if len(args) > 1 {
		logUsage()
		os.Exit(1)
	}
	logHelp()
}

func createSubcommandHandler(args []string) {

	// создаем новый набор флагов
	createCmd := flag.NewFlagSet(CREATE, flag.ContinueOnError)

	// устанавливаем возможные флаги в набор
	name := createCmd.String(NAME_FLAG, "", "Profile name")
	user := createCmd.String(USER_FLAG, "", "User name")
	project := createCmd.String(PROJECT_FLAG, "", "Project name")

	// парсим флаги
	createCmd.Parse(args[1:])

	// флаги должны быть указаны
	if *name == "" || *user == "" || *project == "" {
		logEmptyParametrsError()
		os.Exit(1)
	}

	createCommand(*name, *user, *project)
}

func getSubcommandHandler(args []string) {
	// создаем новый набор флагов
	createCmd := flag.NewFlagSet(GET, flag.ContinueOnError)

	// устанавливаем возможные флаги в набор
	name := createCmd.String(NAME_FLAG, "", "Profile name")

	// парсим флаги
	createCmd.Parse(args[1:])

	// флаги должны быть указаны
	if *name == "" {
		logEmptyParametrsError()
		os.Exit(1)
	}

	getCommand(*name)
}

func listSubcommandHandler(args []string) {
	// для profile list не нужны никакие доп параметры
	if len(args) > 1 {
		logUsage()
		os.Exit(1)
	}

	listCommand()
}

func deleteSubcommandHandler(args []string) {
	// создаем новый набор флагов
	createCmd := flag.NewFlagSet(DELETE, flag.ContinueOnError)

	// устанавливаем возможные флаги в набор
	name := createCmd.String(NAME_FLAG, "", "Profile name")

	// парсим флаги
	createCmd.Parse(args[1:])

	// флаги должны быть указаны
	if *name == "" {
		logEmptyParametrsError()
		os.Exit(1)
	}

	deleteCommand(*name)
}
