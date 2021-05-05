package dns

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/Fndroid/network-helper/common"
)

func SetDNS(servers []string) error {
	nts, err := common.NetworkTypes()
	if err != nil {
		return err
	}
	for _, nt := range nts {
		args := append([]string{"-setdnsservers", nt.String()}, servers...)
		cmd := exec.Command(common.COMMAND, args...)
		cmd.Run()
	}
	return nil
}

func ShowDNS() (string, error) {
	nts, err := common.NetworkTypes()
	if err != nil {
		return "", err
	}
	result := []string{}
	for _, nt := range nts {
		args := []string{"-getdnsservers", nt.String()}
		cmd := exec.Command(common.COMMAND, args...)
		out, err := cmd.CombinedOutput()
		if err != nil || string(out) == "" {
			continue
		}
		servers := common.SplitFilter(string(out), "\n", func(s string) bool {
			return net.ParseIP(s) != nil
		})
		result = append(result, fmt.Sprintf("%s=%s; ", nt.String(), strings.Join(servers, ",")))
	}
	return strings.Join(result, "\n"), nil
}
