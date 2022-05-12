package jsonparse

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtract(t *testing.T) {
	records := make(chan Record)
	var wg sync.WaitGroup
	count := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for record := range records {
			require.True(t, record.valid)
			count++
		}
	}()

	content, err := os.ReadFile(filepath.Join("testdata", "ndjson"))
	require.NoError(t, err)

	err = Extract(records, content, "hostname", "tag", "createTime")
	require.NoError(t, err)
	wg.Wait()
	require.Equal(t, 83, count)
}

func BenchmarkExtract(b *testing.B) {
	content, err := os.ReadFile(filepath.Join("testdata", "ndjson"))
	require.NoError(b, err)
	for i := 0; i < b.N; i++ {
		count := 0
		records := make(chan Record)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for record := range records {
				require.True(b, record.valid)
				count++
			}
		}()
		err = Extract(records, content, "hostname", "tag", "createTime")
		require.NoError(b, err)
		wg.Wait()
		require.Equal(b, 83, count)
	}
}
