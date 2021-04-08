default:
	go build -ldflags="-s -w" -o dist/archive-desktop src/main.go

run:
	go run src/main.go

windows:
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -tags prod -o "dist/archive-desktop.exe" ./src
