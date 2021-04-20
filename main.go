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
		Address: "0x1f81df95c5478059e0e85f7594467bbfe511792a", //&pb.Address{Data: []byte("contract-address")},
		Method:  "get_fee",
		Input: &pb.Input{
			Msg: &pb.Msg{
				Sender:    "0x215D76a620De5D2e9dC552278048C4dA22aA7AD9", //&pb.Address{Data: []byte("sender")},
				Recipient: "0x1f81df95c5478059e0e85f7594467bbfe511792a", //&pb.Address{Data: []byte("recipient")},
			},
			Arguments: args,
		},
	}
}

func SendQueryRequest(endpoint string, address string, method string, input string) string {
	endpoint_cstr := C.CString(endpoint)
	defer C.free(unsafe.Pointer(endpoint_cstr))

	address_cstr := C.CString(address)
	defer C.free(unsafe.Pointer(address_cstr))

	method_cstr := C.CString(method)
	defer C.free(unsafe.Pointer(method_cstr))

	input_cstr := C.CString(input)
	defer C.free(unsafe.Pointer(input_cstr))

	res_cstr := C.ffi_send_query_request(endpoint_cstr, address_cstr, method_cstr, input_cstr)
	defer C.free(unsafe.Pointer(res_cstr))

	return C.GoString(res_cstr)
}

func main() {
	in := makeSampleQueryInput()
	out := SendQueryRequest(rinkebyUrl, "0x1f81df95c5478059e0e85f7594467bbfe511792a", "get_fee", input)
	fmt.Println("Fee:", out.Fee)
}
