package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// WriteProtobufToJSONFile writes protocol buffer message to JSON file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	marshaler := protojson.MarshalOptions{
		Multiline:       true,
		Indent:          " ",
		AllowPartial:    false,
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		Resolver:        nil,
	}
	data, err := marshaler.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal protobuf message to JSON: %w", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}
	return nil
}

//WriteProtobufToBinaryFile writes protocol buffer message to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}
	ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}
	return nil
}

//ReadProtobufFromBinaryFile reads protocol buffer from binary file
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}
	return nil
}
