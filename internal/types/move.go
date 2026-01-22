package types

type (
	// MoveNoteParams contains parameters for moving a note.
	MoveNoteParams struct {
		OldPath   string `json:"oldPath"`
		NewPath   string `json:"newPath"`
		Overwrite bool   `json:"overwrite,omitempty"`
	}

	// MoveResult contains the result of a move operation.
	MoveResult struct {
		Success bool   `json:"success"`
		OldPath string `json:"oldPath"`
		NewPath string `json:"newPath"`
		Message string `json:"message"`
	}
)
