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

const (
	localhostUrl = "http://localhost:4001"
	rinkebyUrl   = "https://rinkeby3-zandbox.zksync.dev"
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

func SendQueryRequest(endpoint string, req *pb.QueryRequest) *pb.QueryResponse {
	endpoint_cstr := C.CString(endpoint)
	defer C.free(unsafe.Pointer(endpoint_cstr))

	raw, err := proto.Marshal(req)
	if err != nil {
		log.Fatalln("Failed to encode req:", err)
	}

	req_cstr := C.CString(string(raw))
	defer C.free(unsafe.Pointer(req_cstr))

	res_cstr := C.ffi_send_query_request(endpoint_cstr, req_cstr)
	defer C.free(unsafe.Pointer(res_cstr))

	res := &pb.QueryResponse{}
	err = proto.Unmarshal([]byte(C.GoString(res_cstr)), res)
	if err != nil {
		log.Fatalln("Failed to encode res:", err)
	}

	return res
}

func main() {
	req := makeSampleQueryRequest()
	response := SendQueryRequest(rinkebyUrl, req)
	out := &pb.GetFeeOutput{}
	if err := response.Output.UnmarshalTo(out); err != nil {
		log.Fatalln("Cannot unmarshal output")
	}
	fmt.Println("Fee:", out.Fee)
}
