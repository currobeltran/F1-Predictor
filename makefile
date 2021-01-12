build:
	go build -o bin/src -v ./src/

install:
	
start:
	./bin/src

start-local: build
	heroku local

test:
	go test ./src/api ./src/f1predictor 

travis:
	docker run -t -v `pwd`:/test currobeltran/f1-predictor:latest

clean:
	rm -r ./bin
