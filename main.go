package main

import (
	"crypto/tls"
	"github.com/hoisie/web"
    "fmt"
)

func main() {
	c, err := LoadConfig("./conf.yaml")
    if err != nil {
        fmt.Println("Error loading config file:", err)
    }
    //setup the key for secret cookies
	s, err := makeKey(64)
    if err != nil {
        fmt.Println("Error setting up secure cookies:", err)
        return
    }
	web.Config.CookieSecret = string(s)

    //setup server and handlers
    server := web.NewServer()
	sm := NewSessionManager(c)
    te, err := NewThemeEngine()
    if err != nil {
        fmt.Println("Error creating theme engine:", err)
        return
    }
	controller := Controller{te, sm}

    controller.Init(c, server)


    if c.UseHTTPS {
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

    } else {
        server.Run(c.Address+":"+c.Port)
    }
}
