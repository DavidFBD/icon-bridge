package watcher_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/icon-project/icon-bridge/cmd/endpoint/chainAPI/chain"
	"github.com/icon-project/icon-bridge/cmd/endpoint/watcher"
	"github.com/icon-project/icon-bridge/common/log"
)

func TestWatchInit(t *testing.T) {
	l := log.New()
	log.SetGlobalLogger(l)
	cfg, err := loadConfig("/home/manish/go/src/work/icon-bridge/cmd/endpoint/example-config.json")
	if err != nil {
		log.Error(err)
		return
	}
	cfgPerMap := map[chain.ChainType]*chain.ChainConfig{}
	for _, ch := range cfg.Chains {
		cfgPerMap[ch.Name] = ch
	}
	_, err = watcher.New(l, cfgPerMap, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	Chains []*chain.ChainConfig `json:"chains"`
}

func loadConfig(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = json.NewDecoder(f).Decode(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
