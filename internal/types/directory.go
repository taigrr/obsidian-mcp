package types

type (
	// DirectoryListing contains the files and directories in a directory.
	DirectoryListing struct {
		Files       []string `json:"files"`
		Directories []string `json:"directories"`
	}

	// PathFilterConfig contains configuration for the path filter.
	PathFilterConfig struct {
		IgnoredPatterns   []string `json:"ignoredPatterns"`
		AllowedExtensions []string `json:"allowedExtensions"`
	}
)
