package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cocobao/rpcxs/server"
	"github.com/cocobao/rpcxs/serverplugin"
	"github.com/rcrowley/go-metrics"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	s := server.NewServer()

	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@127.0.0.1:8972",
		EtcdServers:    []string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"},
		BasePath:       "/rpcx_test",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}

	err := r.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Plugins.Add(r)

	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", "127.0.0.1:8972")

	defer s.Close()

	if len(r.Services) != 1 {
		fmt.Println("failed to register services in etcd")
		return
	}

	if err := r.Stop(); err != nil {
		fmt.Println(err)
		return
	}
}
