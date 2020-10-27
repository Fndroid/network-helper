package sysproxy

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Fndroid/network-helper/common"
	C "github.com/Fndroid/network-helper/constant"
)

func SetBypass(domains []string) error {
	nts, err := common.NetworkTypes()
	if err != nil {
		return err
	}
	for _, nt := range nts {
		args := append([]string{"-setproxybypassdomains", nt.String()}, domains...)
		cmd := exec.Command(common.COMMAND, args...)
		cmd.Run()
	}
	return nil
}

func SetProxy(pt C.ProxyType, server string, port int) error {
	nts, err := common.NetworkTypes()
	if err != nil {
		return err
	}
	for _, nt := range nts {
		args := []string{pt.SetCommand(), nt.String(), server, strconv.Itoa(port)}
		cmd := exec.Command(common.COMMAND, args...)
		cmd.Run()
	}
	return nil
}

func StopProxy(pt C.ProxyType) error {
	nts, err := common.NetworkTypes()
	if err != nil {
		return err
	}
	for _, nt := range nts {
		args := []string{pt.StopCommand(), nt.String(), "off"}
		cmd := exec.Command(common.COMMAND, args...)
		cmd.Run()
	}
	return nil
}

func ShowProxy() (string, error) {
	nts, err := common.NetworkTypes()
	if err != nil {
		return "", err
	}
	result := []string{}
	for _, nt := range nts {
		str := ""
		for _, pt := range []C.ProxyType{C.HTTP, C.HTTPS, C.SOCKS} {
			args := []string{pt.ShowCommand(), nt.String()}
			cmd := exec.Command(common.COMMAND, args...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				continue
			}
			o := common.Format(string(out))
			if o.Enabled {
				if str == "" {
					str = fmt.Sprintf("%s: ", nt.String())
				}
				str += fmt.Sprintf("%s=%s:%d; ", pt.String(), o.Server, o.Port)
			}
		}
		result = append(result, str)
	}
	return strings.Join(result, "\n"), nil
}


