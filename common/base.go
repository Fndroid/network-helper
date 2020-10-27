package common

import (
	"errors"
	"os/exec"
	"strings"

	N "github.com/Fndroid/network-helper/network"
)

const COMMAND = "networksetup"

func NetworkTypes() ([]N.NetworkType, error) {
	cmd := exec.Command(COMMAND, "-listnetworkserviceorder")
	out, err := cmd.CombinedOutput()
	nts := []N.NetworkType{}
	if err != nil {
		return nts, err
	}
	for _, nt := range N.ParseFromText(string(out)) {
		err := testNetwork(nt)
		if err == nil {
			nts = append(nts, nt)
		}
	}
	return nts, nil
}

func testNetwork(nt N.NetworkType) error {
	grep := exec.Command("grep", nt.Device)
	netstat := exec.Command("netstat", "-nr")
	pipe, err := netstat.StdoutPipe()
	if err != nil {
		return err
	}
	defer pipe.Close()
	grep.Stdin = pipe
	netstat.Start()
	out, err := grep.Output()
	if err != nil {
		return err
	}
	if strings.Contains(string(out), "default") {
		return nil
	}
	return errors.New("testNetwork failed")
}
