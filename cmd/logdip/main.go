package main

import (
	"fmt"
	"log"

	"github.com/barklan/logdip/pkg/ingestion"
	"github.com/barklan/logdip/pkg/logdip/config"
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
		log.Fatalf("failed to read config: %v\n", err)
	}

	lg, err := logging.NewAuto()
	if err != nil {
		log.Fatalf("failed to init logger: %v\n", err)
	}

	g := new(errgroup.Group)
	ingestCtrl := ingestion.NewCtrl(lg, cfg)
	g.Go(func() error {
		if err := ingestCtrl.Serve(); err != nil {
			return fmt.Errorf("failed to serve ingestion service: %w", err)
		}

		return nil
	})

	if err := g.Wait(); err == nil {
		lg.Info("main exited")
	} else {
		log.Panic(err)
	}
}
