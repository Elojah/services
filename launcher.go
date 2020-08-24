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
type Launchers struct {
	launchers []Launcher
	configs   Configs
}

// Add a launcher in global up/down.
func (ls *Launchers) Add(l Launcher) {
	ls.launchers = append(ls.launchers, l)
}

// Up all launchers in slice order.
func (ls *Launchers) Up(filename string) error {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	configs := make(Configs)
	if err := json.Unmarshal(raw, &configs); err != nil {
		return err
	}

	ls.configs = configs
	for _, l := range ls.launchers {
		if err := l.read(configs); err != nil {
			return err
		}

		if err := l.Up(configs); err != nil {
			return err
		}
	}

	return nil
}

// Down all launchers in slice order.
func (ls *Launchers) Down() error {
	var e error

	for _, l := range ls.launchers {
		// return last error only
		if err := l.Down(ls.configs); err != nil {
			e = err
		}
	}

	return e
}
