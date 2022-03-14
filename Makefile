cwd = $(shell pwd)

all:
	go build -o build/casing cmd/casing/main.go
	ln -i -s -t ~/.local/bin/ $(cwd)/build/casing
