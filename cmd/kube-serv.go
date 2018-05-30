package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oleg-kalashnikov/kube-serv/pkg/config"
	"github.com/oleg-kalashnikov/kube-serv/pkg/service"
	"github.com/oleg-kalashnikov/kube-serv/pkg/system"
)

func main() {
	c := new(config.Config)
	if err := c.Load(config.SERVICENAME); err != nil {
		log.Fatal(err)
	}

	r, l, err := service.Setup(c)
	if err != nil {
		log.Fatal(err)
	}

	go http.ListenAndServe(fmt.Sprintf("%s:%d", c.LocalHost, c.LocalPort), r)

	signals := system.NewSignals()
	if err := signals.Wait(l, new(system.Handling)); err != nil {
		l.Fatal(err)
	}
}
