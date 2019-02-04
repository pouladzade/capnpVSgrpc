package client

import (
	"fmt"
	"github.com/cryptobank/client"
)

func main() {

	var cli client.Cli
	cli.LoadFlags()
	err := cli.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}

}
