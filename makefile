test:
	go test ./src/f1predictor

build:
	mkdir -p functions
	GOOS=linux GOARCH=amd64 go build -o functions/first ./main.go

travis:
	docker run -t -v `pwd`:/test currobeltran/f1-predictor:latest