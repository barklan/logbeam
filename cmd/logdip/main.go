package main

import (
	"fmt"
	"log"

	"github.com/barklan/logdip/pkg/config"
	"github.com/barklan/logdip/pkg/ingestion"
	"github.com/barklan/logdip/pkg/logging"
	"github.com/barklan/logdip/pkg/system"
	_ "go.uber.org/automaxprocs"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.Println("starting logdip")
	go system.HandleSignals()

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	lg := logging.NewAuto()

	g := new(errgroup.Group)
	ingestCtrl := ingestion.NewCtrl(lg, cfg)
	g.Go(func() error {
		return ingestCtrl.Serve()
	})

	if err := g.Wait(); err == nil {
		fmt.Println("main exited")
	} else {
		log.Panic(err)
	}
}
