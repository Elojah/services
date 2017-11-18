package services

import (
	"encoding/json"
	"io/ioutil"
)

// Launcher represents a service launcher.
type Launcher interface {
	Up(Configs) error
	Down(Configs) error

	read(Configs) error
}

// Launchers is a slice of Launcher used to represent all application services.
type Launchers []Launcher

// Up all launchers in slice order.
func (ls Launchers) Up(filename string) error {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	configs := make(Configs)
	if err := json.Unmarshal(raw, &configs); err != nil {
		return err
	}
	for _, l := range ls {
		if err := l.read(configs); err != nil {
			return err
		}
		if err := l.Up(configs); err != nil {
			return err
		}
	}
	return nil
}
