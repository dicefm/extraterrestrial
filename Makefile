all: clean setup test

setup:
	chmod +x ./lib/bash/unit-test.sh

clean:
	go clean

build:
	go build

test: clean
	go test -v "github.com/dicefm/extraterrestrial/phone"

all-travis: setup test
