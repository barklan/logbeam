package ingestion

import (
	"github.com/barklan/logbeam/pkg/jsonparse"

	"go.uber.org/zap"
)

type Agent interface {
	ParseNDJSON([]byte) error
}

type SimpleAgent struct {
	log     *zap.Logger
	records chan<- jsonparse.Record
}

func NewSimpleAgent(logger *zap.Logger) *SimpleAgent {
	return &SimpleAgent{
		log: logger,
	}
}

func (a *SimpleAgent) ParseNDJSON(data []byte) error {
	return nil
}
