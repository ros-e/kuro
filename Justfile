set shell := ["powershell.exe", "-c"]

name := "kuro"

test:
    go test ./...

build-windows:
    $env:GOOS="windows"; $env:GOARCH="amd64"; go build -o build/{{name}}-windows-amd64.exe

build-linux:
    $env:GOOS="linux"; $env:GOARCH="amd64"; go build -o build/{{name}}-linux-amd64


build-all:
    just build-windows
    just build-linux
