package main

import (
	"fmt"
	"os"
	"os/exec"
)

type util struct {
	Cluster       bool
	ClusterStatus string
}

var test string

func init() {
	//infoLog.Println("init util running---------")
	test = "initial test"
}

func (u *util) GetCluster() {
	u.ClusterStatus = test
}

func (u *util) RunCommand(commandStr string) (string, error) {
	cmdstr := commandStr
	//out, _ := exec.Command("sh", "-c", cmdstr).Output()
	out, res := exec.Command("sh", "-c", cmdstr).CombinedOutput()
	//out, res := exec.Command("sh", "-c", cmdstr).Output()
	if res != nil {
		fmt.Println("show command error res ", res, string(out))
		return "", res
	}
	//fmt.Println("show command res %s", res)
	strout := string(out)
	return strout, nil
}

//background process
func (u *util) RunCommandB(commandStr string) error {
	//  cmd := exec.Command("./script.sh")
	cmd := exec.Command("sh", "-c", commandStr)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		infoLog.Println(err)
	}
	infoLog.Println("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	return nil
}
func (u *util) RunScriptsPath(cmdStr string, chDir string) (string, error) {
	//os.Chdir("/opt/script")
	//commandStr := "bash /opt/script/run.sh"
	commandStr := cmdStr
	cmd := exec.Command("sh", "-c", commandStr)
	//cmd.Dir = "/opt/script"
	cmd.Dir = chDir

	out, err := cmd.Output()
	//      err := cmd.Start()
	if err != nil {
		infoLog.Println(err)
	}
	infoLog.Println(out)
	//infoLog.Println("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	return string(out), nil
}
