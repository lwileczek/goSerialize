package main

import (
	"fmt"
	"myfbe/fbe"
	"myfbe/proto"
	"os"
)

func main() {
	submarine := proto.NewSubStructFromFieldValues("linux", "Spciy")
	emptyPayload := proto.NewPbPayload()
	emptyPayload.SubShop = *submarine

	buf := fbe.NewEmptyBuffer()
	// Serialize the account to the FBE stream
	writer := proto.NewPbPayloadModel(buf)
	if _, err := writer.Serialize(emptyPayload); err != nil {
		fmt.Println("Error serializing the data", err)
		panic("serialization error")
	}
	if ok := writer.Verify(); !ok {
		panic("verify error")
	}

	// Show the serialized FBE size
	fmt.Printf("FBE size: %d\n", writer.Buffer().Size())
	//writeData(writer)

	// Deserialize the account from the FBE stream
	reader := proto.NewPbPayloadModel(writer.Buffer())
	if ok := reader.Verify(); !ok {
		panic("verify error")
	}
	rngPayload, _, err := reader.Deserialize()
	if err != nil {
		panic("deserialization error")
	}

	// Show account content
	fmt.Println("\n\n")
	fmt.Println(rngPayload)
}

func writeData(data *proto.PbPayloadModel) {
	f, err := os.Create("./data.fbe.bin")
	if err != nil {
		panic("error opening file to write")
	}
	defer f.Close()
	b := data.Buffer().Data()
	if _, err = f.Write(b); err != nil {
		fmt.Println("err:", err)
		panic("could not write data to file")
	}
}
