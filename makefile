BinariesDirectory = ./bin
WindowsBinariesDirectory = bin
MainDirectory = ./cmd/main
ConfigDirectory = ./configs
ConfigDirectoryForWindows = configs

all: create-windows-bin win-copy-config build run
run-l: run-linux
run-linux: create-bin copy-config build linux-run

create-windows-bin:
	if not exist "$(BinariesDirectory)" mkdir "$(BinariesDirectory)"

create-bin:
	mkdir -p "$(BinariesDirectory)"

copy-config:
	cp -rfRT "$(ConfigDirectory)" "$(BinariesDirectory)/"

win-copy-config:
	xcopy "$(ConfigDirectoryForWindows)" "$(WindowsBinariesDirectory)" /e /h /c /y /s

build:
	go build -o "$(BinariesDirectory)" "$(MainDirectory)/main.go"

run:
	cd "$(BinariesDirectory)" && main

linux-run:
	cd "$(BinariesDirectory)" && ./main

git-clean-get:
	git reset --hard
	git clean -df
	git status
	git pull