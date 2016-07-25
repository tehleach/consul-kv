package command

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
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

	kvs := map[string]string{}
	for _, element := range fullKVs {
		kvs[element.Key] = string(element.Value)
	}

	data, err := json.MarshalIndent(kvs, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if _, err := file.Write([]byte(data)[:]); err != nil {
		return err
	}

	return nil
}
