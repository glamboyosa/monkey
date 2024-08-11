fmt:
	gofmt -s -w .	
test:
	cd interpreter && go test ./lexer &&  go test ./parser && go test ./ast
run:
	cd interpreter && go run main.go