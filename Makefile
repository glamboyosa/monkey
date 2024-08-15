fmt:
	gofmt -s -w .	
test:
	cd interpreter && go test -v ./lexer &&  go test ./parser && go test ./ast
run:
	cd interpreter && go run main.go