# Go supported OSs and Archs: https://github.com/golang/go/blob/master/src/go/build/syslist.go

export CGO_ENABLED=0

build: clear linux windows

linux:
	GOOS=linux GOARCH=386 go build -o build/openapi2confluence_linux_386 main.go
	GOOS=linux GOARCH=amd64 go build -o build/openapi2confluence_linux_amd64 main.go

windows:
	GOOS=windows GOARCH=386 go build -o build/openapi2confluence_windows_386 main.go
	GOOS=windows GOARCH=amd64 go build -o build/openapi2confluence_windows_amd64 main.go

clear:
	rm -rf build