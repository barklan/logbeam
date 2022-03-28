package config_test

import (
	"os"
	"testing"

	"github.com/barklan/logbeam/pkg/logbeam/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestRead(t *testing.T) { //nolint:funlen
	t.Parallel()
	tests := []struct {
		name    string
		envVars map[string]string
		want    *config.Config
		wantErr bool
	}{
		{
			"test env var",
			map[string]string{
				"LOGBEAM_USER":            "bobik",  // pragma: allowlist secret
				"LOGBEAM_PASSWORD":        "secret", // pragma: allowlist secret
				"LOGBEAM_RETENTION_HOURS": "6",
			},
			&config.Config{
				Username:       "bobik",  // pragma: allowlist secret
				Password:       "secret", // pragma: allowlist secret
				RetentionHours: 6,
			},
			false,
		},
		{
			"default env vars",
			map[string]string{},
			&config.Config{
				Username:       "logbeam", // pragma: allowlist secret
				Password:       "logbeam", // pragma: allowlist secret
				RetentionHours: 48,
			},
			false,
		},
		{
			"negative retention hours should cause error",
			map[string]string{
				"LOGBEAM_RETENTION_HOURS": "-6",
			},
			&config.Config{
				Username:       "logbeam", // pragma: allowlist secret
				Password:       "logbeam", // pragma: allowlist secret
				RetentionHours: 48,
			},
			true,
		},
	}
	for _, tt := range tests { // nolint:paralleltest
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envVars {
				if err := os.Setenv(k, v); err != nil {
					t.Fatalf("failed to set env var: %v", err)
				}
			}
			got, err := config.Read()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, got, tt.want)
			}
			for k := range tt.envVars {
				if err := os.Unsetenv(k); err != nil {
					t.Fatalf("failed to unset env var: %v", err)
				}
			}
		})
	}
}
