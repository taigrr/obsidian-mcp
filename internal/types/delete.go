package types

type (
	// DeleteNoteParams contains parameters for deleting a note.
	DeleteNoteParams struct {
		Path        string `json:"path"`
		ConfirmPath string `json:"confirmPath"`
	}

	// DeleteResult contains the result of a delete operation.
	DeleteResult struct {
		Success bool   `json:"success"`
		Path    string `json:"path"`
		Message string `json:"message"`
	}
)
