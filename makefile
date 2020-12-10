test:
	go test src/main/*
	go test ./src/f1predictor 

build:
	mkdir -p functions
	GOOS=linux GOARCH=amd64 go build -o functions/first netlify/main.go

start:
	go build src/main/main.go
	./main

travis:
	docker run -t -v `pwd`:/test currobeltran/f1-predictor:latest

clean: 
	rm ./main