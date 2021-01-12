build:
	go build -o bin/f1-predictor -v ./src/main/

install:
	
start:
	./bin/src

test:
	go test ./src/api ./src/f1predictor 

travis:
	docker run -t -v `pwd`:/test currobeltran/f1-predictor:latest

clean:
	rm f1predictor
