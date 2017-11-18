package services

import (
	"fmt"
)

// Config represents structure configs used for each service.
type Config interface{}

// Configs represents all configs structure retrieved in a config file.
type Configs map[Namespace]Config

// NewConfigs returns a new empty configuration map with namespaces set.
func NewConfigs(ns ...Namespace) *Configs {
	cfgs := Configs{}
	for _, n := range ns {
		cfgs[n] = nil
	}
	return &cfgs
}

func (c *Configs) read(fileconfigs Configs) error {
	for ns := range *c {
		fc, ok := fileconfigs[ns]
		if !ok {
			return fmt.Errorf("missing configuration namespace %s", ns)
		}
		(*c)[ns] = fc
	}
	return nil
}
