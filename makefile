run:
	go main.go

test:
	go test -v ./...

migrate-linux-build:
	env GOOS=linux GOARCH=amd64 go build -o ./migrate database/main.go
	@echo "###SUCCESS BUILD FOR LINUX ENVIRONMENT###"

migrate-mac-build:
	env GOOS=darwin GOARCH=amd64 go build -o ./migrate database/main.go
	@echo "###SUCCESS BUILD FOR MAC ENVIRONMENT###"

linux-build:
	env GOOS=linux GOARCH=amd64 go build -o ./migrate database/main.go
	env GOOS=linux GOARCH=amd64 go build -o ./govel main.go
	@echo "###SUCCESS BUILD FOR LINUX ENVIRONMENT###"

mac-build:
	env GOOS=darwin GOARCH=amd64 go build -o ./migrate database/main.go
	env GOOS=darwin GOARCH=amd64 go build -o ./govel main.go
	@echo "###SUCCESS BUILD FOR MAC ENVIRONMENT###"