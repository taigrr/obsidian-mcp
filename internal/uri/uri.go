// Package uri provides Obsidian URI generation.
package uri

import (
	"net/url"
	"strings"
)

// GenerateObsidianURI generates an Obsidian URI for a given note path.
// Uses the absolute path format: obsidian:///absolute/path/to/note
func GenerateObsidianURI(vaultPath, notePath string) string {
	// Remove leading slash from notePath if present
	cleanPath := notePath
	cleanPath = strings.TrimPrefix(cleanPath, "/")

	// Construct absolute path
	absolutePath := vaultPath + "/" + cleanPath

	// Remove .md extension if present (Obsidian handles this automatically)
	absolutePath = strings.TrimSuffix(absolutePath, ".md")

	// URI encode the path, but keep slashes as slashes
	parts := strings.Split(absolutePath, "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}
	encodedPath := strings.Join(parts, "/")

	// Remove leading slash since we add obsidian:/// prefix
	encodedPath = strings.TrimPrefix(encodedPath, "/")

	return "obsidian:///" + encodedPath
}
