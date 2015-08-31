package data

import (
	"fmt"

	"github.com/tjuqxy/build-redis-data/conf"
	"github.com/tjuqxy/build-redis-data/tools"
)

var (
	confFile = "./conf/redis_command.yml"
)

func MakeRedisData() {
	c := conf.DealConf(confFile)

	// Deal conf fail
	if c == nil {
		return
	}

	// deal read command
	readCommand  := c["read-command"].(map[interface{}]interface{})
	for dataType, commandConf := range readCommand {
		fmt.Println(dataType)
		eachCommand := commandConf.(map[interface{}]interface{})
		for cmd, param := range eachCommand {
			paramStr := tools.BuildData(param)
			fmt.Println(cmd, paramStr)
		}
		fmt.Println()
		fmt.Println()
	}

	// deal write command
	writeCommand := c["write-command"].(map[interface{}]interface{})
	for dataType, commandConf := range writeCommand {
		fmt.Println(dataType)
		eachCommand := commandConf.(map[interface{}]interface{})
		for cmd, param := range eachCommand {
			paramStr := tools.BuildData(param)
			fmt.Println(cmd, paramStr)
		}
		fmt.Println()
		fmt.Println()
	}

	return
}
