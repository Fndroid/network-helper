package main

import (
	"flag"
	"fmt"
	"strings"

	C "github.com/Fndroid/network-helper/constant"
	D "github.com/Fndroid/network-helper/dns"
	S "github.com/Fndroid/network-helper/sysproxy"
)

var bypass = flag.String("bypass", "", "bypass join by \",\"")
var httpProxy = flag.String("http", "", "http proxy server and port")
var httpsProxy = flag.String("https", "", "https proxy server and port")
var socksProxy = flag.String("socks", "", "socks proxy server and port")
var stop = flag.Bool("stop", false, "disable all proxies")
var show = flag.Bool("show", false, "show all proxies")
var dns = flag.String("dns", "", "dns servers join by \",\" to set, \"reset\" to clear, \"query\" to show current servers")

func main() {
	flag.Parse()

	if *bypass != "" {
		dms := strings.Split(*bypass, ",")
		S.SetBypass(dms)
	}

	if *httpProxy != "" {
		proxy, err := C.ParseProxyURL(*httpProxy)
		if err == nil {
			S.SetProxy(C.HTTP, *proxy)
		} else {
			fmt.Println(err)
		}
	}

	if *httpsProxy != "" {
		proxy, err := C.ParseProxyURL(*httpsProxy)
		if err == nil {
			S.SetProxy(C.HTTPS, *proxy)
		} else {
			fmt.Println(err)
		}
	}

	if *socksProxy != "" {
		proxy, err := C.ParseProxyURL(*socksProxy)
		if err == nil {
			S.SetProxy(C.SOCKS, *proxy)
		} else {
			fmt.Println(err)
		}
	}

	if *stop {
		S.StopProxy(C.HTTP)
		S.StopProxy(C.SOCKS)
		S.StopProxy(C.HTTPS)
	}

	if *show {
		out, err := S.ShowProxy()
		if err == nil {
			fmt.Println(out)
		}
	}

	if *dns != "" {
		var servers []string
		if *dns == "reset" {
			servers = []string{"Empty"}
		} else if *dns == "query" {
			out, err := D.ShowDNS()
			if err == nil {
				fmt.Println(out)
			}
		} else {
			servers = strings.Split(*dns, ",")
		}
		D.SetDNS(servers)
	}
}
