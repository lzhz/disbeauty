package disbeauty

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/lzhz/disbeauty/proto"
)

func NewClient(addr string, size int) (*proto.DisBeautyClient, error) {
	socket, err := thrift.NewTSocket(addr)
	if err != nil {
		return nil, err
	}

	if err = socket.Open(); err != nil {
		return nil, err
	}

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTBufferedTransportFactory(size)
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	client := proto.NewDisBeautyClientFactory(transportFactory.GetTransport(socket), thrift.NewTBinaryProtocolFactoryDefault())
	return client, nil
}

type ProtoHandler struct {
}

func (self *ProtoHandler) Display(cmd string) (string, error) {
	return Exec(cmd), nil
}

func NewServer(addr string, size int) (*thrift.TSimpleServer, error) {
	processor := proto.NewDisBeautyProcessor(&ProtoHandler{})
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTBufferedTransportFactory(size)
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	serverSocket, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return nil, err
	}

	if err = serverSocket.Open(); err != nil {
		return nil, err
	}

	server := thrift.NewTSimpleServer4(processor, serverSocket, transportFactory, protocolFactory)
	return server, nil
}
