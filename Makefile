all: clean setup test

setup:
	chmod +x ./lib/bash/unit-test.sh

clean:
	go clean

build:
	go build

test: clean
	./lib/bash/unit-test.sh

all-travis: setup test
