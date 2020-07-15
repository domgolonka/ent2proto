build:
	go build -o main main.go
	chmod +x main

run:
	./start.sh
	./main generate ./test/schema/

add:
	./start.sh
	go get github.com/domgolonka/ent2proto

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...
	#go mod vendor
help:
	@echo "Help commands: build (build it). run (start, test). add (./start, add to gopath). install_deps (install deps)"