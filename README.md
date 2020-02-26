# rpcxs

根据rpcx v5.0.0裁剪，只支持etcd服务发现，tcp以及http协议
由于rpcx支持组件比较多，下载安装比较复杂，有些插件不使用代理还下载不了，所以这里做了一些裁剪，只留下比较基本的功能，从而简化下载安装。

### Server
```
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
```

