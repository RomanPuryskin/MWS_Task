package source

import "fmt"

func logUsage() {
	fmt.Println(`Usage: mws.exe <command> [subcommand] [options])
	Print ./mws.exe help
	To get more information`)
}

func logUnknownCommandError(command string) {
	fmt.Printf("Unknown command %s:\n ", command)
	fmt.Println(`Print ./mws.exe help
	To get more information`)
}

func logEmptyParametrsError() {
	fmt.Println("Flag parametrs are required")
}

func logProfileNotExistsError(name string) {
	fmt.Printf("Reading error of profile %s do not exist\n", name)
}

func logProfile(name string, profile *Profile) {
	fmt.Printf("Name: %s\n\tUser: %s\n\tProject: %s\n", name, profile.User, profile.Project)
}

func logHelp() {
	fmt.Println(`Usage: ./mws.exe <command> [subcommand] [flags]
	Commands:
  	profile - (profile management)
		Subcommands:
		create - (create a new profile)
			--name     profile name (required)
			--user     user name (required)
			--project  project name (required)
		get  -	 (get profile with current name)
			--name     profile name (required)
		list  -  (get all profiles)
		delete - (delete profile with current name)
			--name     profile name (required)
	help  -  (show usage information)`)
}
