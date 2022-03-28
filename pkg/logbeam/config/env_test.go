package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestRead(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		envVars map[string]string
		want    *Config
		wantErr bool
	}{
		{
			"test env var",
			map[string]string{
				"LOGBEAM_USER":            "bobik",
				"LOGBEAM_PASSWORD":        "secret",
				"LOGBEAM_RETENTION_HOURS": "6",
			}, // pragma: allowlist secret
			&Config{
				Username:       "bobik",
				Password:       "secret",
				RetentionHours: 6,
			}, // pragma: allowlist secret
			false,
		},
		{
			"default env vars",
			map[string]string{},
			&Config{
				Username:       "logbeam",
				Password:       "logbeam",
				RetentionHours: 48,
			},
			false,
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
			got, err := Read()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
			for k := range tt.envVars {
				if err := os.Unsetenv(k); err != nil {
					t.Fatalf("failed to unset env var: %v", err)
				}
			}
		})
	}
}
