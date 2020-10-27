package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
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
		host, port, err := net.SplitHostPort(*httpProxy)
		if err == nil {
			portInt, err := strconv.Atoi(port)
			if err == nil {
				S.SetProxy(C.HTTP, host, portInt)
			} else {
				fmt.Println("port is not a integer")
			}
		} else {
			fmt.Println("http proxy format error")
		}
	}

	if *httpsProxy != "" {
		host, port, err := net.SplitHostPort(*httpsProxy)
		if err == nil {
			portInt, err := strconv.Atoi(port)
			if err == nil {
				S.SetProxy(C.HTTPS, host, portInt)
			} else {
				fmt.Println("port is not a integer")
			}
		} else {
			fmt.Println("https proxy format error")
		}
	}

	if *socksProxy != "" {
		host, port, err := net.SplitHostPort(*socksProxy)
		if err == nil {
			portInt, err := strconv.Atoi(port)
			if err == nil {
				S.SetProxy(C.SOCKS, host, portInt)
			} else {
				fmt.Println("port is not a integer")
			}
		} else {
			fmt.Println("socks proxy format error")
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
