
.PHONY:
.SILENT:
# Имя исполняемого файла
TARGET = .bin/mws.exe

# Исходные файлы
SOURCES = cmd/main.go

# Команда для компиляции
build:
	go build -o $(TARGET) $(SOURCES)

run: build
	$(TARGET)