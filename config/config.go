package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Address struct {
	Net  string `yaml:",omitempty"`
	Addr string
}

func (a Address) String() string {
	return fmt.Sprintf("%s://%s", a.Net, a.Addr)
}

type Port struct {
	Listen  Address `yaml:",flow"`
	Connect Address `yaml:",flow"`
}

type Config struct {
	Ports []Port
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer f.Close()
	raw, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}
	cfg := &Config{}
	err = yaml.Unmarshal(raw, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}
	for i := range cfg.Ports {
		p := &cfg.Ports[i]
		if p.Listen.Net == "" {
			p.Listen.Net = "tcp"
		}
		if p.Connect.Net == "" {
			p.Connect.Net = "tcp"
		}
	}
	return cfg, nil
}