package cli

// Defaults

const VERSION = "1.0/1.1"
const INTERPRETED = true
const LOGGING = false
const DEBUG = true

type Cli struct {
	version         string
	flagInterpreted bool
	flagLog         bool
	flagDebug       bool
	flagHelp        bool
}

// Gonna work on this
