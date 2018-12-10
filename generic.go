package main

import (
	"os"
)

type Generic struct {
}

func (g *Generic) Name() string {
	a := os.Getenv("XDG_CUPRRENT_DESKTOP")
	if a != "" {
		return a
	}
	b := os.Getenv("XDG_SESSION_DESKTOP")
	if b != "" {
		return b
	}
	return "Missing"
}

func (g *Generic) ExecutablesExists() bool {
	return true
}

func (g *Generic) Running() bool {
	return true
}

func (g *Generic) SetWallpaper(imageFilename string) error {
	return run("feh --bg-scale " + imageFilename)
}
