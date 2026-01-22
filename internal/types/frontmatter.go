package types

type (
	// UpdateFrontmatterParams contains parameters for updating frontmatter.
	UpdateFrontmatterParams struct {
		Path        string         `json:"path"`
		Frontmatter map[string]any `json:"frontmatter"`
		Merge       bool           `json:"merge,omitempty"`
	}

	// FrontmatterValidationResult contains the result of frontmatter validation.
	FrontmatterValidationResult struct {
		IsValid  bool     `json:"isValid"`
		Errors   []string `json:"errors"`
		Warnings []string `json:"warnings"`
	}
)
