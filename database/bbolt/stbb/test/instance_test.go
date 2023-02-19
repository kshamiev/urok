package test

import (
	"os"
	"testing"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/kshamiev/urok/database/bbolt/stbb"
)

// go test -run=TestStats
func TestStats(t *testing.T) {
	inst := newInstance(t)
	go inst.Stats()
	time.Sleep(time.Hour)
}

func TestNewInstance(t *testing.T) {
	newInstance(t)
}

func newInstance(t *testing.T) *stbb.Instance {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	cfg := &stbb.Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		t.Fatal(err)
	}
	inst, err := stbb.NewInstance(cfg)
	if err != nil {
		t.Fatal(err)
	}
	return inst
}
