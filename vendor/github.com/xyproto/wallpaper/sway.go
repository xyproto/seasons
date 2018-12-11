package wallpaper

// Sway windowmanager detector
type Sway struct {
}

func (s *Sway) Name() string {
	return "Sway"
}

func (s *Sway) ExecutablesExists() bool {
	return which("sway") != "" && which("swaymsg") != ""
}

func (s *Sway) Running() bool {
	return hasE("SWAYSOCK") && (containsE("GDMSESSION", "sway") || containsE("XDG_SESSION_DESKTOP", "sway") || containsE("XDG_CURRENT_DESKTOP", "sway"))
}

// SetWallpaper sets the desktop wallpaper, given an image filename.
// The image must exist and be readable.
func (s *Sway) SetWallpaper(imageFilename string) error {
	return run("swaymsg 'output \"*\" background " + imageFilename + " fill'")
}
