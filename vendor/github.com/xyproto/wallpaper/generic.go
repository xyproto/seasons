package wallpaper

// This is the fallback if no specific windowmanager has been detected

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

// SetWallpaper sets the desktop wallpaper, given an image filename.
// The image must exist and be readable.
// `feh` is used for setting the desktop background, and must be in the PATH.
func (g *Generic) SetWallpaper(imageFilename string) error {
	return run("feh --bg-scale " + imageFilename)
}
