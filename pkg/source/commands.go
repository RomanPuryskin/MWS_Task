package source

import (
	"bufio"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	User    string `yaml:"user"`
	Project string `yaml:"project"`
}

func createCommand(profileName, userName, projectName string) {

	newProfile := Profile{
		User:    userName,
		Project: projectName,
	}

	fileName := profileName + ".yaml"

	yamlData, _ := yaml.Marshal(&newProfile)

	// создаем файл и записываем в него структуру Profile с заданными полями
	file, _ := os.Create(fileName)
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(yamlData)
	writer.Flush()
}

func getCommand(profileName string) {
	fileName := profileName + ".yaml"
	// попытаемся открыть файл, если файла с таким именем нет, то выйдем с напечатанной ошибкой
	file, err := os.Open(fileName)
	if err != nil {
		logProfileNotExistsError(profileName)
		os.Exit(1)
	}
	defer file.Close()

	// если файл с таким именем есть, прочитаем его содержимое в структуру Profile и напечатаем
	yamlData, _ := os.ReadFile(fileName)
	var curProfile Profile
	yaml.Unmarshal(yamlData, &curProfile)

	logProfile(profileName, &curProfile)
}

func deleteCommand(profileName string) {

	fileName := profileName + ".yaml"

	// удалим файл с заданным именем, если его нет, то выйдем с напечатанной ошибкой
	err := os.Remove(fileName)
	if err != nil {
		logProfileNotExistsError(profileName)
		os.Exit(1)
	}
}

func listCommand() {
	dir := "./"

	// по условия все файлы хранятся в директории бинарника
	files, _ := os.ReadDir(dir)

	// пройдемся по содержимому папки, нас интересуют только файлы с расширением .yaml
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".yaml" {

			// для каждого файла прочитаем его содержимое в структуру Profile и напечатаем
			yamlData, _ := os.ReadFile(file.Name())
			var curProfile Profile
			yaml.Unmarshal(yamlData, &curProfile)

			// Получим имя профиля, для этого из строки с названием файла "вычтем" строку с расширением (.yaml)
			profileName := file.Name()[:len(file.Name())-len(filepath.Ext(file.Name()))]

			logProfile(profileName, &curProfile)
		}
	}

}
