package main

import (
	"context"
	"log"
	"time"

	"github.com/cocobao/rpcxs/client"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func main() {
	d := client.NewEtcdV3Discovery(
		"/rpcx_test",
		"Arith",
		[]string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"},
		nil)
	xclient := client.NewXClient(
		"Arith",
		client.Failover,
		client.RoundRobin,
		d,
		client.DefaultOption)
	defer xclient.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Printf("failed to call: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("%d * %d = %d\n", args.A, args.B, reply.C)

		time.Sleep(5 * time.Second)
	}
}
