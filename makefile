build:
	

install:
	

test:
	go test ./src/main ./src/f1predictor 

travis:
	docker run -t -v `pwd`:/test currobeltran/f1-predictor:latest
