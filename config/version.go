package config

// Version is set at build time via ldflags.
// Example: go build -ldflags "-X 'swoptape/config.Version=0.1.0'"

var Version = "dev"
var Branch = "unknown"
var Commit = "unknown"
