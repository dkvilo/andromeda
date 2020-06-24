clean:
	rm -rf ./andromeda

build:
	go build andromeda.go

run:
	./andromeda

.PHONY: build run
