#!/bin/sh

rm -r ./bin
sleep 1

go build src/dog.go && go build src/cf.go && go build src/now.go && go build src/ok.go && go build src/user.go && go build src/usrid.go && go build src/wf.go && go build src/crdir.go

mkdir bin
mv dog cf now ok user usrid wf crdir ./bin
