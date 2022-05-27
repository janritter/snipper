prepare:
	go mod download

build: prepare
	go build -o ./bin/snipper -v -ldflags "-X github.com/janritter/snipper/cmd.gitSha=`git rev-parse HEAD` -X github.com/janritter/snipper/cmd.buildTime=`date +'%Y-%m-%d_%T'` -X github.com/janritter/snipper/cmd.version=LOCAL_BUILD"

tests:
	go test ./... -v  --cover

run:
	go run main.go
