package main

import "runtime/debug"

var version = getVersion()

func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}

	var revision, modified string
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			revision = s.Value
		case "vcs.modified":
			modified = s.Value
		}
	}

	if revision == "" {
		return "dev"
	}

	if len(revision) > 7 {
		revision = revision[:7]
	}

	if modified == "true" {
		return revision + "-dirty"
	}
	return revision
}
