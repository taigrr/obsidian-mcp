package uri

import "testing"

func TestGenerateObsidianURI(t *testing.T) {
	tests := []struct {
		name      string
		vaultPath string
		notePath  string
		want      string
	}{
		{
			name:      "simple path",
			vaultPath: "/Users/test/vault",
			notePath:  "notes/test.md",
			want:      "obsidian:///Users/test/vault/notes/test",
		},
		{
			name:      "leading slash in note path",
			vaultPath: "/Users/test/vault",
			notePath:  "/notes/test.md",
			want:      "obsidian:///Users/test/vault/notes/test",
		},
		{
			name:      "path with spaces",
			vaultPath: "/Users/test/my vault",
			notePath:  "my notes/test file.md",
			want:      "obsidian:///Users/test/my%20vault/my%20notes/test%20file",
		},
		{
			name:      "path with special chars",
			vaultPath: "/Users/test/vault",
			notePath:  "notes/test (copy).md",
			want:      "obsidian:///Users/test/vault/notes/test%20%28copy%29",
		},
		{
			name:      "no md extension",
			vaultPath: "/Users/test/vault",
			notePath:  "notes/test.txt",
			want:      "obsidian:///Users/test/vault/notes/test.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateObsidianURI(tt.vaultPath, tt.notePath)
			if got != tt.want {
				t.Errorf("GenerateObsidianURI() = %q, want %q", got, tt.want)
			}
		})
	}
}
