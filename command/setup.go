package command

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	consul "github.com/hashicorp/consul/api"
)

const (
	localConsulURL  = "127.0.0.1:8500"
	defaultFileName = "data.json"
)

func getKVClient(ipaddress string) (*consul.KV, error) {
	config := consul.DefaultConfig()
	config.Address = ipaddress

	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return client.KV(), nil
}

func readJSON(filename string, prefix string) (consul.KVPairs, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	inpairs := map[string]string{}
	err = json.Unmarshal(data, &inpairs)
	if err != nil {
		return nil, err
	}

	var kvs consul.KVPairs
	for key, value := range inpairs {
		if prefix == "" || strings.HasPrefix(key, prefix) {
			kvs = append(kvs, &consul.KVPair{Key: key, Value: []byte(value)})
		}
	}
	return kvs, nil
}

func saveKVsToFile(fullKVs consul.KVPairs, fileName string) error {
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
