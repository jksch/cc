build:
	go build

build_web:
	$(MAKE) -C internal/webapp assable

run:
	./cc

all: build_web build run

test:
	go test -cover -count=1 ./...
