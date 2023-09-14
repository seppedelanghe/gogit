build:
	go build -ldflags "-s -w" -o bin/gogit main.go

install: build
	cp bin/gogit /usr/local/bin/
