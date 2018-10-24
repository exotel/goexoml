build: test
	mkdir -p bin
	go build -o ./bin/buildergenerator ./buildergenerator/buildergenerator.go

test:
	go test ./...
