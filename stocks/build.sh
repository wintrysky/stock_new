#/bin/bash

# set project name
EXENAME=stocks

mkdir -p /data/src
mkdir -p /data/bin
cd /data/src
# clean
rm -rf /data/bin/*

# build
#CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build $1 -o ../bin/$EXENAME main.go
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $1 -o ../bin/$EXENAME main.go
#CGO_ENABLED=0 GOOS=linux GOARCH=386 go build $1 -o ../bin/$EXENAME main.go
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $1 -o ../bin/$EXENAME main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /data/bin/stocks /data/src/main.go
#CGO_ENABLED=0 GOOS=windows GOARCH=386 go build $1 -o ../bin/$EXENAME.exe  main.go
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $1 -o ../bin/$EXENAME.exe  main.go

# dependency
if [ -f "./config.yaml" ]
then
	cp ./config.yaml /data/bin/config.yaml
fi

# clean
rm -f /data/bin/**/.gitkeep
echo "ALL DONE" >> /data/bin/log.txt



