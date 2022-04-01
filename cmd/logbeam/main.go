package main

import (
	"log"

	"github.com/barklan/logbeam/internal/config"
	"github.com/barklan/logbeam/pkg/logging"
	"github.com/barklan/logbeam/pkg/system"
	_ "go.uber.org/automaxprocs"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.Println("starting logbeam")
	go system.HandleSignals()

	_, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v\n", err)
	}

	lg, err := logging.NewAuto()
	if err != nil {
		log.Fatalf("failed to init logger: %v\n", err)
	}

	g := new(errgroup.Group)
	// ingestCtrl := ingestion.NewCtrl(lg, cfg)
	// g.Go(func() error {
	// 	if err := ingestCtrl.Serve(); err != nil {
	// 		return fmt.Errorf("failed to serve ingestion service: %w", err)
	// 	}

	// 	return nil
	// })

	if err := g.Wait(); err == nil {
		lg.Info("main exited")
	} else {
		log.Panic(err)
	}
}
