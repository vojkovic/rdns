# get from .env file i.e.
OPENWEATHERMAP_API_KEY ?= $(shell cat .env | grep OPENWEATHERMAP_API_KEY | cut -d '=' -f2)

all:
	export OPENWEATHERMAP_API_KEY=$(OPENWEATHERMAP_API_KEY) && go run .

docker:
	docker build -t rdns .
	docker run -p 8080:8080 -e OPENWEATHERMAP_API_KEY=$(OPENWEATHERMAP_API_KEY) rdns

.PHONY: all clean docker
