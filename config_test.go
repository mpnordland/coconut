package main

import (
    "testing"
)
const ConfigFile string = "test_config.yaml"
const ConfigContent string = `#test comment
address: 0.0.0.0
port: 8888
pages: 
    /about: pages/about.md

articlesperpage: 5`

func getConfig(t *testing.T) *Config{
    config, err := LoadConfig(ConfigFile)
    if err != nil || config == nil{
        t.Fatal("Couldn't load config from file", err)
    }
    return config
}

func TestLoadConfigWithExistingFile(t *testing.T) {
    getConfig(t)
}

func TestLoadConfigWithMissingFile(t *testing.T) {
    //This file shouldn't exist
    config, err := LoadConfig("test_missing.yaml")
    if err == nil || config != nil{
        t.Fail()
    }
}

func TestLoadConfigOutput(t *testing.T){
    config := getConfig(t)
    if config.Address != "0.0.0.0" {
        t.Fail()
    }
    if config.Port != "8888"{
        t.Fail()
    }
    if page, ok := config.Pages["/about"]; !ok || page != "pages/about.md" {
        t.Fail()
    }
    if config.ArticlesPerPage != 5 {
        t.Fail()
    } 
    
}

