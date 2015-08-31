package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

// 处理配置，返回一个map
func DealConf(confFile string) map[string]interface{} {
	c := make(map[string]interface{})

	// read file
	content, err := ioutil.ReadFile(confFile)
	if err != nil {
		fmt.Println("Read file failed:", err)
		return nil
	}

	// unmarshal
	err = yaml.Unmarshal(content, c)
	if err != nil {
		fmt.Println("Unmarshal conf failed:", err)
		return nil
	}

	return c
}
