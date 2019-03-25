NAME=app
USER_ID=$(shell id -u ${USER})
GROUP_ID=$(shell id -g ${USER})

setup:
	docker build -t node-8-15.1-yarn .

run:
	docker run -it --rm \
	       	-u ${USER_ID}:${GROUP_ID} \
	       	-v $(shell pwd)/:/build \
		node-8-15.1-yarn \
	       	/bin/bash

start:
	docker run -it --rm \
		--network host \
	       	-u ${USER_ID}:${GROUP_ID} \
		-v $(shell pwd)/:/build \
	       	node-8-15.1-yarn \
		/bin/bash -c "yarn start"

test: 
	docker run -it --rm \
		-v $(shell pwd)/:/build \
	       	node-8-15.1-yarn \
		/bin/bash -c \
	       	"yarn test --watchAll --coverage --collectCoverageFrom=src/app/**/*js"
	
build:
	docker run -it --rm \
		-u ${USER_ID}:${GROUP_ID} \
		-v $(shell pwd)/:/build \
	       	node-8-15.1-yarn \
		/bin/bash -c "yarn build"
