package main

import (
	"crypto/tls"
	"fmt"
	"github.com/hoisie/web"
)

func main() {
	c, err := LoadConfig("./conf.yaml")
	if err != nil {
		fmt.Println("Error loading config file:", err)
	}

	//setup server and handlers
	server := web.NewServer()
	te, err := NewThemeEngine()
	if err != nil {
		fmt.Println("Error creating theme engine:", err)
		return
	}
	controller := Controller{themeEngine: te, articlesPerPage: c.ArticlesPerPage}
	controller.Init(c, server)
	te.Run()
	switch c.Protocol {
	case "https":
		//setup for https
		config := tls.Config{
			Time: nil,
		}

		config.Certificates = make([]tls.Certificate, 1)
		config.Certificates[0], err = tls.LoadX509KeyPair(c.Certfile, c.Keyfile)

		if err != nil {
			fmt.Println("error, could not load ssl cert and/or key")
			return
		}
		server.RunTLS(c.Address+":"+c.Port, &config)
	case "http":
		server.Run(c.Address + ":" + c.Port)
	case "fcgi":
		server.RunFcgi(c.Address + ":" + c.Port)

	}
}
