package main

/*
#cgo LDFLAGS: -L./lib -lzinctx
#include <stdlib.h>
#include "./lib/zinctx.h"
*/
import "C"

import (
	"fmt"
	"log"
	"unsafe"

	pb "github.com/motxx/zinctx-go/protos"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func makeSampleQueryRequest() *pb.QueryRequest {
	args, err := anypb.New(&pb.GetFeeArguments{})
	if err != nil {
		log.Fatalln("Cannot marshal arguments to any")
	}
	return &pb.QueryRequest{
		Address: "contract-example", //&pb.Address{Data: []byte("contract-address")},
		Method:  "get_fee",
		Input: &pb.Input{
			Msg: &pb.Msg{
				Sender:    "sender-address",    //&pb.Address{Data: []byte("sender")},
				Recipient: "recipient-address", //&pb.Address{Data: []byte("recipient")},
			},
			Arguments: args,
		},
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
	out := &pb.GetFeeOutput{}
	if err := response.Output.UnmarshalTo(out); err != nil {
		log.Fatalln("Cannot unmarshal output")
	}
	fmt.Println("Fee:", out.Fee)
}
