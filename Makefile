build: go.mod
	go build -o bin/arr-custom-notification cmd/arr-custom-notification/main.go

build-linux: go.mod
	GOOS=linux GOARCH=amd64 go build -o bin/arr-custom-notification-linux-x86 cmd/arr-custom-notification/main.go

build-linux-static: go.mod
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/arr-custom-notification-linux-x86-static cmd/arr-custom-notification/main.go

run:
	go run cmd/arr-custom-notification/main.go

clean:
	rm -r bin

.PHONY: build run clean