APP ?= lgen
ENV ?= .env

clean:
	@rm -f bin/${APP}

build: clean
	@go build \
		-tags go_json \
		-o bin/${APP} \
		./cmd/

gorun: build
	@bin/${APP}

docker-compose-up:
	@docker-compose \
		--env-file ${ENV} \
		up -d

docker-compose-down:
	@docker-compose \
		--env-file ${ENV} \
		down
