package serializer_test

import (
	"testing"

	"github.com/Arpeet-gupta/go-grpc-protobuf/v4/pb"
	"github.com/Arpeet-gupta/go-grpc-protobuf/v4/sample"
	"github.com/Arpeet-gupta/go-grpc-protobuf/v4/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryfile := "./testdata/laptop.bin"
	jsonfile := "./testdata/laptop.json"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryfile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryfile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonfile)
	require.NoError(t, err)

}
