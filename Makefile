fmt:
	gofmt -s -w .	
test:
	cd interpreter && go test ./lexer 
run:
	cd interpreter && go run main.go