package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	consul "github.com/hashicorp/consul/api"
)

//CmdRestore restores values to consul kv
func CmdRestore(c *cli.Context) error {
	//get flags
	destHost := c.String("t")
	srcHost := c.String("f")
	srcFile := c.String("s")
	path := c.String("p")

	//check flags
	if destHost == "" {
		destHost = localConsulURL
	}
	if srcHost == "" && srcFile == "" {
		return errors.New("Missing required flag -from or -srcfile, need to specify one")
	}

	//grab dest kv client
	destKV, err := getKVClient(destHost)
	if err != nil {
		return err
	}

	//grab src kvs
	var kvs consul.KVPairs
	if srcFile != "" {
		kvs, err = readJSON(srcFile)
		if err != nil {
			return err
		}
	} else { //srcHost != ""
		srcKV, err := getKVClient(srcHost)
		if err != nil {
			return err
		}
		kvs, _, err = srcKV.List("/", nil)
		if err != nil {
			return err
		}
	}

	//set dest kvs
	for _, val := range kvs {
		if strings.HasPrefix(val.Key, path) {
			_, err := destKV.Put(val, nil)
			if err != nil {
				fmt.Printf("Error on put: %v", err)
			}
		}
	}

	return nil
}
