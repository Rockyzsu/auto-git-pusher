package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var CONFIG_PATH = "config.json"

type gitInfo struct {
	Path string `json:"path"`
	Msg  string `json:"msg"`
}

type GitInfoList []gitInfo

func JsonParse(filename string) (GitInfoList, error) {
	infoList := GitInfoList{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return infoList, err
	}
	err = json.Unmarshal(data, &infoList)
	if err != nil {
		log.Println(err)
		return infoList, err
	}
	return infoList, nil
}

func runCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		tmp_info := string(output)
		if strings.Contains(tmp_info, "nothing to commit") {
			return tmp_info, nil
		}
		fmt.Println("runCommand", err)
		return "", err
	}
	return string(output), nil
}

var COMMANDS = []string{
	"git pull origin master",
	"git status",
	"git add .",
	"git commit -m '%s'",
	"git push origin master",
}

func main() {
	info, err := JsonParse(CONFIG_PATH)
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range info {
		err = os.Chdir(v.Path)
		if err != nil {
			fmt.Println("change directory failed： ", err)
			return
		}
		for _, command := range COMMANDS {
			if strings.Contains(command, "commit") {
				command = fmt.Sprintf(command, v.Msg)
			}
			resutl, err := runCommand(command)
			if err != nil {
				fmt.Printf("执行%s失败\n", command)
				fmt.Println(resutl)

				fmt.Println(err)
				return
			}
			fmt.Println(resutl)
		}

	}
}
