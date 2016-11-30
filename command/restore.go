package command

import (
	"errors"
	"fmt"
	"strings"

	consul "github.com/hashicorp/consul/api"
	"gopkg.in/urfave/cli.v1"
)

//CmdRestore restores values to consul kv
func CmdRestore(c *cli.Context) error {
	//get flags
	destHost := c.String("t")
	srcHost := c.String("f")
	srcFile := c.String("n")
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
		kvs, err = readJSON(srcFile, path)
		if err != nil {
			return err
		}
	} else { //srcHost != ""
		srcKV, err := getKVClient(srcHost)
		if err != nil {
			return err
		}
		kvs, _, err = srcKV.List(fmt.Sprintf("/%v", path), nil)
		if err != nil {
			return err
		}
	}

	fmt.Println("Keys to restore:")
	for _, val := range kvs {
		fmt.Println(val.Key)
	}
	var response string
	fmt.Printf("Restore above keys to %v (y/n)? ", destHost)
	for {
		fmt.Scanln(&response)
		if response == "y" || response == "n" {
			if response == "n" {
				fmt.Println("Will not restore. Exiting...")
				return nil
			}
			break
		}
		fmt.Println("Please enter y or n.")
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
