package command

import (
	"fmt"

	consul "github.com/hashicorp/consul/api"
	cli "gopkg.in/urfave/cli.v1"
)

const setArgs = 3

//CmdSet sets a kv
func CmdSet(c *cli.Context) error {
	if len(c.Args()) != setArgs {
		return fmt.Errorf("Incorrect number of arguments. Found %v, expected %v.\nUsage: consul-kv set <consulAddr> <key> <value>", len(c.Args()), setArgs)
	}

	destHost := c.Args().Get(0)
	key := c.Args().Get(1)
	value := c.Args().Get(2)

	//grab dest kv client
	destKV, err := getKVClient(destHost)
	if err != nil {
		return err
	}

	if _, err = destKV.Put(&consul.KVPair{Key: key, Value: []byte(value)}, nil); err != nil {
		return fmt.Errorf("Error setting key: %v", err)
	}

	return nil
}
