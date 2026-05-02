package main

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/taigrr/obsidian-mcp/internal/filesystem"
	"github.com/taigrr/obsidian-mcp/internal/frontmatter"
	"github.com/taigrr/obsidian-mcp/internal/pathfilter"
)

func TestExtractTags(t *testing.T) {
	frontmatter := map[string]any{
		"tags": []any{"Project", "go/mcp", "Project"},
	}
	content := "# Title\n\nInline tags: #Daily #go/mcp #go_mcp"

	got := extractTags(frontmatter, content)
	want := []string{"daily", "go/mcp", "go_mcp", "project"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("extractTags() = %v, want %v", got, want)
	}
}

func TestExtractLinks(t *testing.T) {
	content := "See [[Note One]], [[note one|alias]], [[Second Note#Heading]], and [[Third]]."

	got := extractLinks(content)
	want := []string{"note one", "second note", "third"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("extractLinks() = %v, want %v", got, want)
	}
}

func TestAddRelation(t *testing.T) {
	tests := []struct {
		name     string
		existing string
		newRel   string
		want     string
	}{
		{name: "empty", existing: "", newRel: "backlink", want: "backlink"},
		{name: "append", existing: "shared-tags", newRel: "backlink", want: "shared-tags,backlink"},
		{name: "dedupe", existing: "shared-tags,backlink", newRel: "backlink", want: "shared-tags,backlink"},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := addRelation(testCase.existing, testCase.newRel)
			if got != testCase.want {
				t.Fatalf("addRelation(%q, %q) = %q, want %q", testCase.existing, testCase.newRel, got, testCase.want)
			}
		})
	}
}

func TestHandleListDefaultsToRoot(t *testing.T) {
	vaultPath := t.TempDir()
	pf := pathfilter.New(nil)
	fh := frontmatter.New()
	fileSystem = filesystem.New(vaultPath, pf, fh)

	testFile := filepath.Join(vaultPath, "note.md")
	testDir := filepath.Join(vaultPath, "projects")
	if err := os.Mkdir(testDir, 0o755); err != nil {
		t.Fatalf("Mkdir() error = %v", err)
	}
	if err := os.WriteFile(testFile, []byte("# test\n"), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	_, got, err := handleList(context.Background(), nil, ListInput{})
	if err != nil {
		t.Fatalf("handleList() error = %v", err)
	}

	wantFiles := []string{"note.md"}
	wantDirs := []string{"projects"}
	if !reflect.DeepEqual(got.Files, wantFiles) {
		t.Fatalf("handleList().Files = %v, want %v", got.Files, wantFiles)
	}
	if !reflect.DeepEqual(got.Directories, wantDirs) {
		t.Fatalf("handleList().Directories = %v, want %v", got.Directories, wantDirs)
	}
}
