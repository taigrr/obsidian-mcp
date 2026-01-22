// Package types defines all data structures used across the MCP server.
package types

type (
	// ParsedNote represents a parsed markdown note with frontmatter.
	ParsedNote struct {
		Frontmatter     map[string]any `json:"frontmatter"`
		Content         string         `json:"content"`
		OriginalContent string         `json:"originalContent"`
	}

	// NoteWriteParams contains parameters for writing a note.
	NoteWriteParams struct {
		Path        string         `json:"path"`
		Content     string         `json:"content"`
		Frontmatter map[string]any `json:"frontmatter,omitempty"`
		Mode        string         `json:"mode,omitempty"` // "overwrite", "append", "prepend"
	}

	// NoteInfo contains metadata about a note.
	NoteInfo struct {
		Path           string `json:"path"`
		Size           int64  `json:"size"`
		Modified       int64  `json:"modified"` // timestamp in milliseconds
		HasFrontmatter bool   `json:"hasFrontmatter"`
		ObsidianURI    string `json:"obsidianUri,omitempty"`
	}
)
