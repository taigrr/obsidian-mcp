package types

type (
	// PatchNoteParams contains parameters for patching a note.
	PatchNoteParams struct {
		Path       string `json:"path"`
		OldString  string `json:"oldString"`
		NewString  string `json:"newString"`
		ReplaceAll bool   `json:"replaceAll,omitempty"`
	}

	// PatchNoteResult contains the result of a patch operation.
	PatchNoteResult struct {
		Success    bool   `json:"success"`
		Path       string `json:"path"`
		Message    string `json:"message"`
		MatchCount int    `json:"matchCount,omitempty"`
	}
)
