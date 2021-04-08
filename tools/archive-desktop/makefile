default:
	echo ""

build:
	go build -ldflags="-s -w" -o dist/rrm src/main.go

buildd:
	env GOOS=linux GOARCH=arm go build -ldflags="-s -w" -tags prod -o ../dist/linux-arm32/rrm ./src
	env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -tags prod -o ../dist/linux-arm64/rrm ./src
	env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -tags prod -o ../dist/linux-386/rrm ./src
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags prod -o ../dist/linux-amd64/rrm ./src
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -tags prod -o ../dist/darwin-amd64/rrm ./src
	env GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -tags prod -o ../dist/darwin-arm64/rrm ./src
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -tags prod -o ../dist/windows-amd64/rrm.exe ./src
	env GOOS=windows GOARCH=386 go build -ldflags="-s -w" -tags prod -o ../dist/windows-386/rrm.exe ./src
