all: run

clean:
	@rm -rf main main.exe

run:
	@go run *.go

build: clean
	@go build -o main -ldflags="-w -s" *.go
	@strip -x main
	@upx -9 main