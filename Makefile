PROG_NAME=gopcb

build:
	@go build -o ${PROG_NAME} ./cmd/

clean:
	@go clean
	@rm -f ${PROG_NAME}

test:
	@go test -v ./... -cover