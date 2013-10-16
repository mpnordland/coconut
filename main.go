package main

import (
	"crypto/tls"
	"github.com/hoisie/web"
    "fmt"
)

func main() {
	c, err := LoadConfig("./conf.yaml")
    handleFatalError(err)

	sm := NewSessionManager(c)
	ph := NewPageHandler(sm)

    //add handlers
	web.Get("/", ph.front)
	web.Get("/publish", ph.publish)
	web.Post("/publish", ph.publishPost)
	web.Get("/login", ph.login)
	web.Post("/login", ph.loginPost)
	web.Get("/(.*)", ph.article)

    //setup the key for secret cookies
	s, err := makeKey(64)
	handleFatalError(err)
	web.Config.CookieSecret = string(s)

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

        web.RunTLS(c.Address+":"+c.Port, &config)

    } else {
        web.Run(c.Address+":"+c.Port)
    }
}
