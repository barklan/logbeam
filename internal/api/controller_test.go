package api

import (
	"testing"

	"github.com/barklan/logbeam/internal/config"
	"github.com/barklan/logbeam/pkg/logging"
	"github.com/stretchr/testify/assert"
)

func NewMockController(t *testing.T) *Controller {
	t.Helper()

	lg, err := logging.Dev()
	assert.NoError(t, err)

	conf, err := config.Read()
	assert.NoError(t, err)

	ctrl := NewController(lg, conf)

	return ctrl
}
