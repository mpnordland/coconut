package main

import (
	"io/ioutil"
	"launchpad.net/goyaml"
)

type Config struct {
	Address         string
	Port            string
	Pages           map[string]string
	ArticlesPerPage int
	Protocol        string
	Certfile        string
	Keyfile         string
}

func LoadConfig(filename string) (*Config, error) {
	cont, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := Config{}
	err = goyaml.Unmarshal(cont, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
