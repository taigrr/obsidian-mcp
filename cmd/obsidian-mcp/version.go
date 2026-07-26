package main

import "runtime/debug"

// version is overridable via -ldflags "-X main.version=..." for release builds.
var version = getVersion()

func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}

	if v := info.Main.Version; v != "" && v != "(devel)" {
		return v
	}

	var revision, modified string
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			revision = setting.Value
		case "vcs.modified":
			modified = setting.Value
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
