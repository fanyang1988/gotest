package main

import (
	"encoding/binary"
	"flag"
	"github.com/cihub/seelog"
	"net"
)

var address = flag.String("a", "127.0.0.1:80", "address to send tcp packet")
var str = flag.String("s", "", "string to send")

func main() {
	defer seelog.Flush()

	flag.Parse()

	conn, err := net.Dial("tcp", *address)
	if err != nil {
		seelog.Errorf("err %s", err.Error())
		panic(err)
	}

	var endian = binary.BigEndian

	// | 4-length |  string  |
	length := len(*str)
	binary.Write(conn, endian, uint32(length))
	binary.Write(conn, endian, []byte(str))

	var readLength uint32
	err = binary.Read(conn, endian, &readLength)
	if err != nil {
		seelog.Errorf("err %s", err.Error())
		panic(err)
	}

	datas := make([]byte, readLength)
	err = binary.Read(conn, endian, datas)
	if err != nil {
		seelog.Errorf("err %s", err.Error())
		panic(err)
	}

	seelog.Infof("read %d : %s", readLength, string(datas))
}
