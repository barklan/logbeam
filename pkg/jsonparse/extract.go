package jsonparse

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
)

type Record struct {
	Raw    []byte
	KeyMap map[string]string
	valid  bool
}

func Extract(records chan<- Record, ndjson []byte, keys ...string) error {
	scanner := bufio.NewScanner(bytes.NewReader(ndjson))
	var p fastjson.Parser
	for scanner.Scan() {
		bytes := scanner.Bytes()
		raw := make([]byte, len(bytes))
		copy(raw, bytes)

		v, err := p.ParseBytes(bytes)
		if err != nil {
			return fmt.Errorf("failed to parse record: %w", err)
		}
		valid := true
		km := make(map[string]string)
		for _, key := range keys {
			val := v.GetStringBytes(key)
			if val == nil {
				valid = false
				continue
			}
			km[key] = string(val)
		}

		records <- Record{
			Raw:    raw,
			KeyMap: km,
			valid:  valid,
		}
	}
	close(records)
	return nil
}
