package types

type (
	// SearchParams contains parameters for basic searching notes.
	SearchParams struct {
		Query             string `json:"query"`
		Limit             int    `json:"limit,omitempty"`
		SearchContent     bool   `json:"searchContent,omitempty"`
		SearchFrontmatter bool   `json:"searchFrontmatter,omitempty"`
		CaseSensitive     bool   `json:"caseSensitive,omitempty"`
	}

	// SearchResult contains a single search result with minified field names.
	SearchResult struct {
		Path       string `json:"p"`   // path
		Title      string `json:"t"`   // title
		Excerpt    string `json:"ex"`  // excerpt
		MatchCount int    `json:"mc"`  // matchCount
		LineNumber int    `json:"ln"`  // lineNumber
		URI        string `json:"uri"` // obsidianUri
	}

	// SearchParamsAdvanced contains parameters for advanced searching.
	SearchParamsAdvanced struct {
		Query         string `json:"query"`
		UseRegex      bool   `json:"useRegex,omitempty"`
		CaseSensitive bool   `json:"caseSensitive,omitempty"`
		ContextLines  int    `json:"contextLines,omitempty"`
		Limit         int    `json:"limit,omitempty"`
		Offset        int    `json:"offset,omitempty"`
	}

	// SearchMatchAdvanced represents a single match within a file.
	SearchMatchAdvanced struct {
		Line    int    `json:"line"`
		Context string `json:"context"`
		IsTag   bool   `json:"isTag,omitempty"`
	}

	// SearchResultAdvanced represents search results for a single file.
	SearchResultAdvanced struct {
		Path    string                `json:"path"`
		Matches []SearchMatchAdvanced `json:"matches"`
	}
)
