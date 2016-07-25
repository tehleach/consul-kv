package command

import (
	"encoding/json"
	"io/ioutil"

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

func readJSON(filename string) (consul.KVPairs, error) {
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
		kvs = append(kvs, &consul.KVPair{Key: key, Value: []byte(value)})
	}
	return kvs, nil
}
