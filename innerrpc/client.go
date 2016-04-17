package innerrpc

import (
	"net"
	"net/rpc"
	"time"
	"log"
)

type (
	Client struct {
		conn *rpc.Client
	}
)

func NewClient(dsn string, timeout time.Duration) (*Client, error) {
	for i := 0; i < 1000; i++ {
		log.Println("-- dialing DATA PROVIDER")
		connection, err := net.DialTimeout("tcp", dsn, timeout)
		if err != nil {
			log.Println("-- Redial DATA PROVIDER in 3 seconds")
			time.Sleep(3* time.Second)
		} else {
			log.Println("-- dial finished, DATA PROVIDER is connected")
			return &Client{conn: rpc.NewClient(connection)}, nil
		}
	}
	log.Fatal("-- Internal Network Failure, DATA PROVIDER is unresponsive ")
	return nil, err
}

func (c *Client) Get(key *GetArgs) (*GetReply, error) {
	var item *GetReply
	err := c.conn.Call("RPC.Get", key, &item)
	return item, err
}

func (c *Client) Put(item *PostArgs) (bool, error) {
	var added bool
	err := c.conn.Call("RPC.Put", item, &added)
	return added, err
}

func (c *Client) ClusterPut(item *ClusterPostArgs) (bool, error) {
	var added bool
	err := c.conn.Call("RPC.ClusterPut", item, &added)
	return added, err
}

func (c *Client) Delete(key *GetArgs) (bool, error) {
	var deleted bool
	err := c.conn.Call("RPC.Delete", key, &deleted)
	return deleted, err
}

func (c *Client) Stats() (*Requests, error) {
	requests := &Requests{}
	err := c.conn.Call("RPC.Stats", true, requests)
	return requests, err
}
