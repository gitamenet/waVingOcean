package main

import (
	"bytes"
	"io"
	"os"

	"github.com/gitamenet/waVingOcean/configure"
	"github.com/golang/protobuf/proto"
	wavingocean "github.com/xiaokangwang/waVingOcean"
)

func main() {
	var inbuf bytes.Buffer
	io.Copy(&inbuf, os.Stdin)
	conffile := new(configure.WaVingOceanConfigure)
	err := proto.Unmarshal(inbuf.Bytes(), conffile)
	if err != nil {
		panic(err)
	}
	wavingocean.Ignite(*conffile)
}
