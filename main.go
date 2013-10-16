package main

import (
	"crypto/tls"
	"github.com/hoisie/web"
)

func main() {
	c, err := LoadConfig("./conf.yaml")
    handleFatalError(err)
	sm := NewSessionManager(c)
	ph := NewPageHandler(sm)
	config := tls.Config{
		Time: nil,
	}

	s, err := makeKey(64)
	handleFatalError(err)

	web.Config.CookieSecret = string(s)
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.LoadX509KeyPair(c.Certfile, c.Keyfile)
    handleFatalError(err)

	web.Get("/", ph.front)
	web.Get("/publish", ph.publish)
	web.Post("/publish", ph.publishPost)
	web.Get("/login", ph.login)
	web.Post("/login", ph.loginPost)
	web.Get("/(.*)", ph.article)
	web.RunTLS(c.Address+":"+c.Port, &config)
}
