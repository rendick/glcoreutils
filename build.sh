#!/bin/sh

rm -r ./bin
sleep 1

go build src/dog.go && go build src/now.go && go build src/ok.go && go build src/user.go && go build src/usrid.go && go build src/wf.go && go build src/crdir.go && go build src/vd.go && go build src/cwd.go && go build src/rf.go

sleep 1

mkdir bin
mv dog now ok user usrid wf crdir vd cwd rf ./bin
