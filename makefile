APPLICATION_NAME=pomogoro

build:
	go build -o ${APPLICATION_NAME} -ldflags="-s -w" main.go

run:
	go run main.go

clean:
	go clean
	rm -f ${APPLICATION_NAME}