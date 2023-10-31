package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go_simple_crud/test/test_reliability/RootMsg"
	"io/ioutil"
	"net/http"
	"strings"
)

var testUniqueID int

func getTaskID() int {
	testUniqueID++
	return testUniqueID
}

func executeCmd(cmd string, ctype int) (interface{}, error) {
	var retData interface{}
	if ctype == RootMsg.STAT_ONLINE || ctype == RootMsg.CYCLE {
		taskID := getTaskID()
		//cmd = strings.TrimSpace(cmdStr) + fmt.Sprintf("--id=%d", taskID)  // need debug
		fmt.Println(taskID)
	}

	resp, err := http.Post("http://192.168.200.10:8888/", "application/x-www-form-urlencoded", strings.NewReader(fmt.Sprintf("input=%s", cmd)))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	//respbody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return nil, err
	//}
	//fmt.Println("Response:", string(respbody))

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}

		if len(respBody) > 0 {
			fmt.Println(string(respBody))

			//用于将 JSON 数据解析为 Go 语言中的结构体或其他数据类型
			err = json.Unmarshal(respBody, &retData)
			if err != nil {
				return nil, err
			}
			// fmt.Printf("%#v\n", retData) // for debug reading
			//fmt.Println(retData)
			//logger.Debug("ret data: ", retData)
		} else {
			//fmt.Println(retData)
			//logger.Debug("ret null: ", retData)
			fmt.Println("[E] respBody empty")
		}
	}

	return retData, nil
}

func main() {
	cmdRaw := "topo"
	cmdStr := base64.StdEncoding.EncodeToString([]byte(cmdRaw))
	fmt.Println(cmdStr)
	executeCmd(cmdStr, 0x49)
}
