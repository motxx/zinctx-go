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

	pb "github.com/motxx/zinctx-go/protos"
	"google.golang.org/protobuf/proto"
)

func makeSampleQueryRequest() *pb.QueryRequest {
	args := map[string]*pb.Value{
		"hoge": &pb.Value{
			ValueOneOf: &pb.Value_Intval{123},
		},
		"fuga": &pb.Value{
			ValueOneOf: &pb.Value_Stringval{"bar"},
		},
	}
	return &pb.QueryRequest{
		Address:   "1234567890abcdef",
		Arguments: args,
	}
}

func SendQueryRequest(req *pb.QueryRequest) *pb.QueryResponse {
	raw, err := proto.Marshal(req)
	if err != nil {
		log.Fatalln("Failed to encode req:", err)
	}

	req_cstr := C.CString(string(raw))
	defer C.free(unsafe.Pointer(req_cstr))

	res_cstr := C.ffi_send_query_request(req_cstr)
	defer C.free(unsafe.Pointer(res_cstr))

	res := &pb.QueryResponse{}
	err = proto.Unmarshal([]byte(C.GoString(res_cstr)), res)
	if err != nil {
		log.Fatalln("Failed to encode res:", err)
	}

	return res
}

func main() {
	proto := makeSampleQueryRequest()
	response := SendQueryRequest(proto)
	println(response.Output.GetStringval())
}
