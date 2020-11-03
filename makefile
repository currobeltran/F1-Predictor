test:
	go test ./src/f1predictor

build: 
	go build ./src/f1predictor

travis:
	docker build -t f1-predictor-test .
	docker run -t -v `pwd`:/test f1-predictor-test