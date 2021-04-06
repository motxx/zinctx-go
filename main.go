package main

/*
#cgo LDFLAGS: -L./lib -lzinctx
#include <stdlib.h>
#include "./lib/zinctx.h"
*/
import "C"

import (
	"log"
	"unsafe"

	proto "github.com/golang/protobuf/proto"
	pb "github.com/motxx/zinctx-go/protos"
)

func main() {

	args := map[string]*pb.Value{
		"hoge": &pb.Value{
			ValueOneOf: &pb.Value_Intval{123},
		},
		"fuga": &pb.Value{
			ValueOneOf: &pb.Value_Stringval{"bar"},
		},
	}
	sq := &pb.SendQuery{
		Address:   "1234567890abcdef",
		Arguments: args,
	}

	raw, err := proto.Marshal(sq)
	if err != nil {
		log.Fatalln("Failed to encode send_query:", err)
	}

	ptr := C.CString(string(raw))
	defer C.free(unsafe.Pointer(ptr))

	C.ffi_send_query(ptr)
}
