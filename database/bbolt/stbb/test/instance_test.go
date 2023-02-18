package test

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/kshamiev/urok/database/bbolt/stbb"
)

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
