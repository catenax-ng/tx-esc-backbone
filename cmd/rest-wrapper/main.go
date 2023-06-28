package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}
	client, err := NewChainClient(context.Background(), *config)
	if err != nil {
		log.Fatal(err)
	}
	client.Poll(context.Background())
	newRouter(client).handleRequests()
}

func readConfig() (*Config, error) {
	config := &Config{}
	configFile, err := os.Open("wrapper-config.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	bz, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bz, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
