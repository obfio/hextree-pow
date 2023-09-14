package hextree

// Config is the config returned from hextree.io to solve POW
type Config struct {
	Prefix     string `json:"prefix"`
	Difficulty int    `json:"difficulty"`
}
