package main

// Sway windowmanager detector
type Sway struct {
}

func (s *Sway) Name() string {
	return "sway"
}

func (s *Sway) ExecutablesExists() bool {
	return which("sway") != "" && which("swaymsg") != ""
}

func (s *Sway) Running() bool {
	return hasE("SWAYSOCK") && (containsE("GDMSESSION", "sway") || containsE("XDG_SESSION_DESKTOP", "sway") || containsE("XDG_CURRENT_DESKTOP", "sway"))
}

func (s *Sway) SetWallpaper(imageFilename string) error {
	// TODO: Check if imageFilename exists (and expand $HOME and ~)
	return run("swaymsg 'output \"*\" background " + imageFilename + " fill'")
}
