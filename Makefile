APP	?= lgen

clean:
	@rm -f bin/${APP}

build: clean
	@go build			\
		-tags go_json	\
		-o bin/${APP}	\
		./cmd/

run: build
	@bin/${APP}
