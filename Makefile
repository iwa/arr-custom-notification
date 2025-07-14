build: go.mod
	go build -o bin/arr-custom-notification cmd/arr-custom-notification/main.go

run:
	go run cmd/arr-custom-notification/main.go

clean:
	rm -r bin

.PHONY: build run clean