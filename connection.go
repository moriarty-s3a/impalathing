package impalathing

import (
	"errors"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/chenjingping/impalathing/services/beeswax"
	impala "github.com/chenjingping/impalathing/services/impalaservice"
)

type Options struct {
	PollIntervalSeconds float64
	BatchSize           int32
}

var (
	DefaultOptions = Options{PollIntervalSeconds: 0.1, BatchSize: 1024}
)

type Connection struct {
	client    *impala.ImpalaServiceClient
	handle    *beeswax.QueryHandle
	transport thrift.TTransport
	options   Options
}

func Dail(name, host, port string) (interface{}, error) {
	return Connect(host, port, DefaultOptions)
}

func Connect(host, port string, options Options) (*Connection, error) {
	socket, err := thrift.NewTSocket(fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		return nil, err
	}

	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := transportFactory.GetTransport(socket)
	if err != nil {
		return nil, err
	}

	if err := transport.Open(); err != nil {
		return nil, err
	}

	client := impala.NewImpalaServiceClientFactory(transport, protocolFactory)

	return &Connection{client, nil, transport, options}, nil
}

func CloseCnn(itf interface{}) (err error) {
	if cnn, ok := itf.(*Connection); ok {
		return cnn.Close()
	}

	return errors.New("connection conversion failed")
}

func KeepAlive(itf interface{}) (err error) {
	if cnn, ok := itf.(*Connection); ok {
		return cnn.Ping()
	}

	return errors.New("connection convert failed")
}

func (c *Connection) isOpen() bool {
	return c.client != nil
}

func (c *Connection) Close() error {
	if c.isOpen() {
		if c.handle != nil {
			_, err := c.client.Cancel(c.handle)
			if err != nil {
				return err
			}
			c.handle = nil
		}

		c.transport.Close()
		c.client = nil
	}
	return nil
}

func (c *Connection) Query(query string) (RowSet, error) {
	bquery := beeswax.Query{}

	bquery.Query = query
	bquery.Configuration = []string{}

	handle, err := c.client.Query(&bquery)

	if err != nil {
		return nil, err
	}

	return newRowSet(c.client, handle, c.options), nil
}

// Add ping func for connection
func (c *Connection) Ping() error {
	if c.isOpen() {
		return c.client.PingImpalaService()
	}
	return nil
}
