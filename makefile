build:
	go build -o f1predictor -v ./src/main

install:
	

test:
	go test ./src/main ./src/f1predictor 

travis:
	docker run -t -v `pwd`:/test currobeltran/f1-predictor:latest

clean:
	rm f1predictor

run-heroku:
	./f1predictor