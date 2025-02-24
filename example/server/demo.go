package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/zeromicro/goctl-android/example/internal/config"
	"github.com/zeromicro/goctl-android/example/internal/handler"
	"github.com/zeromicro/goctl-android/example/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/demo-api.yaml", "the config file")

func listIps() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ips := []string{}
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ips = append(ips, v.IP.String())
			case *net.IPAddr:
				ips = append(ips, v.IP.String())
			}
		}
	}
	return ips, nil
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	if ips, err := listIps(); err != nil {
		fmt.Printf("list ips failed: %v\n", err)
	} else {
		fmt.Println("list ips:")
		for i, ip := range ips {
			fmt.Printf("[%d] %s\n", i, ip)
		}
	}
	server.Start()
}
