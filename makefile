test:
	go test ./src/f1predictor

build: 
	go build ./src/f1predictor

travis:
	docker run -t -v `pwd`:/test currobeltran/f1-predictor:latest