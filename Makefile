cross_build:
	go install
	env GOOS=linux GOARCH=386 go build -o bin/shorty-linux
	env GOOS=windows GOARCH=386 go build -o bin/shorty-windows

build:
	go install
	cp env.example .env
	go build -o bin/shorty .

run:
	./bin/shorty


