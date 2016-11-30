package command

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

//CmdBackup backs up consul kv's to a local file
func CmdBackup(c *cli.Context) error {
	srcHost := c.String("f")
	fileName := c.String("n")
	path := c.String("p")

	//check flags
	if srcHost == "" {
		srcHost = localConsulURL
	}
	if fileName == "" {
		fileName = defaultFileName
	}

	//grab src kv client
	srcKV, err := getKVClient(srcHost)
	if err != nil {
		return err
	}

	fullKVs, _, err := srcKV.List(fmt.Sprintf("/%v", path), nil)
	if err != nil {
		return err
	}

	return saveKVsToFile(fullKVs, fileName)
}
