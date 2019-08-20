build:
	GOOS=windows GOARCH=amd64 go build -o wdshot.exe
	GOOS=darwin  GOARCH=amd64 go build -o wdshot.osx
	GOOS=linux   GOARCH=amd64 go build -o wdshot
