default:
	rm -rf dist
	rm -rf ./src/version.txt
	date -u --iso-8601=minutes > ./src/version.txt
	env GOOS=linux GOARCH=arm go build -tags prod -o ./dist/alshx-linux-arm32 ./src
	env GOOS=linux GOARCH=arm64 go build -tags prod -o ./dist/alshx-linux-arm64 ./src
	env GOOS=linux GOARCH=386 go build -tags prod -o ./dist/alshx-linux-386 ./src
	env GOOS=linux GOARCH=amd64 go build -tags prod -o ./dist/alshx-linux-amd64 ./src
	env GOOS=darwin GOARCH=amd64 go build -tags prod -o ./dist/alshx-darwin-amd64 ./src
	env GOOS=darwin GOARCH=arm64 go build -tags prod -o ./dist/alshx-darwin-arm64 ./src
	env GOOS=windows GOARCH=amd64 go build -tags prod -o ./dist/alshx-windows-amd64.exe ./src
	env GOOS=windows GOARCH=386 go build -tags prod -o ./dist/alshx-windows-386.exe ./src

run:
	go run -tags prod ./src

build:
	go build -o ./dist/alshx ./src

