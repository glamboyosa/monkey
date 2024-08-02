fmt:
	gofmt -s -w .	
test:
	cd interpreter && go test ./lexer 