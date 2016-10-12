package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/consul/api"
)

var prefixFile = "./dir"
var prefixConsul = "jetstack-vault"
var kv *api.KV

func getKeyValue(path string, info os.FileInfo, err error) error {
	log.Infof("path: %s", path)

	if info.IsDir() {
		return nil
	}

	file, e := ioutil.ReadFile(path)
	if e != nil {
		log.Warnf("File error: %v", e)
	}

	var kvPair api.KVPair
	json.Unmarshal(file, &kvPair)
	kvPair.Key = filepath.Join(prefixConsul, kvPair.Key)
	log.Infof("Results: %+v\n", kvPair)

	_, err = kv.Put(&kvPair, nil)
	if err != nil {
		panic(err)
	}

	return nil

}

func main() {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv = client.KV()

	filepath.Walk("./dir", getKeyValue)
}
