build: bin clean
	mkdir -p bin
	go build -o bin/oza ./cmd/oza

test:
	go test -v ./...

clean:
	rm -f bin/oza
