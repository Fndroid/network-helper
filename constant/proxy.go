package constant

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

const (
	HTTP ProxyType = iota
	HTTPS
	SOCKS
)

type Proxy struct {
	Host string
	Port int
	Username string
	Password string
}

func ParseProxyURL(url string) (p *Proxy, err error) {
	var username, password  string
	proxy := url
	ps := strings.Split(url, "@")
	if len(ps) == 2{
		up := strings.Split(ps[0], ":")
		if len(up) == 2 {
			username = up[0]
			password = up[1]
		}
		proxy = ps[1]
	}
	host, port, err := splitHostPort(proxy)
	if err != nil {
		return nil, errors.New("proxy url format error")
	}
	return &Proxy{
		Host: host,
		Port: port,
		Username: username,
		Password: password,
	}, nil
}

func splitHostPort(url string) (host string, port int, err error) {
	host, portStr, err := net.SplitHostPort(url)
	if err != nil {
		return "", 0, err
		
	}
	port, err = strconv.Atoi(portStr)
	if err != nil {
		return "", 0, err	
	}
	return host, port, nil
}

type ProxyType int

func (pt ProxyType) String() string {
	switch pt {
	case HTTPS:
		return "https"
	case SOCKS:
		return "socks"
	default:
		return "http"
	}
}

func (pt ProxyType) SetCommand() string {
	switch pt {
	case HTTPS:
		return "-setsecurewebproxy"
	case SOCKS:
		return "-setsocksfirewallproxy"
	default:
		return "-setwebproxy"
	}
}

func (pt ProxyType) StopCommand() string {
	switch pt {
	case HTTPS:
		return "-setsecurewebproxystate"
	case SOCKS:
		return "-setsocksfirewallproxystate"
	default:
		return "-setwebproxystate"
	}
}

func (pt ProxyType) ShowCommand() string {
	switch pt {
	case HTTPS:
		return "-getsecurewebproxy"
	case SOCKS:
		return "-getsocksfirewallproxy"
	default:
		return "-getwebproxy"
	}
}
