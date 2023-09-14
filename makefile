build:
	go build -v -o bin/gogit main.go

install: build
	cp bin/gogit /usr/local/bin/
	gogit init
