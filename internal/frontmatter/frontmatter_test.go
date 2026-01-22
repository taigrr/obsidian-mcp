package frontmatter

import (
	"strings"
	"testing"
)

func TestHandler_ParseWithFrontmatter(t *testing.T) {
	handler := New()

	content := `---
title: Test Note
tags: [test, example]
created: 2023-01-01
---

# Test Note

This is a test note with frontmatter.`

	result := handler.Parse(content)

	if result.Frontmatter["title"] != "Test Note" {
		t.Errorf("Frontmatter[title] = %v, want %q", result.Frontmatter["title"], "Test Note")
	}

	tags, ok := result.Frontmatter["tags"].([]any)
	if !ok {
		t.Errorf("Frontmatter[tags] is not []any: %T", result.Frontmatter["tags"])
	} else if len(tags) != 2 || tags[0] != "test" || tags[1] != "example" {
		t.Errorf("Frontmatter[tags] = %v, want [test, example]", tags)
	}

	expectedContent := "# Test Note\n\nThis is a test note with frontmatter."
	if strings.TrimSpace(result.Content) != expectedContent {
		t.Errorf("Content = %q, want %q", strings.TrimSpace(result.Content), expectedContent)
	}
}

func TestHandler_ParseWithoutFrontmatter(t *testing.T) {
	handler := New()

	content := `# Test Note

This is a test note without frontmatter.`

	result := handler.Parse(content)

	if len(result.Frontmatter) != 0 {
		t.Errorf("Frontmatter = %v, want empty map", result.Frontmatter)
	}

	if result.Content != content {
		t.Errorf("Content = %q, want %q", result.Content, content)
	}
}

func TestHandler_StringifyWithFrontmatter(t *testing.T) {
	handler := New()

	fm := map[string]any{
		"title": "Test Note",
		"tags":  []string{"test", "example"},
	}
	content := "# Test Note\n\nContent here."

	result, err := handler.Stringify(fm, content)
	if err != nil {
		t.Fatalf("Stringify() error = %v", err)
	}

	if !strings.Contains(result, "---") {
		t.Error("Result should contain frontmatter delimiters")
	}
	if !strings.Contains(result, "title: Test Note") {
		t.Error("Result should contain title")
	}
	if !strings.Contains(result, "tags:") {
		t.Error("Result should contain tags")
	}
	if !strings.Contains(result, "# Test Note") {
		t.Error("Result should contain content")
	}
}

func TestHandler_StringifyWithoutFrontmatter(t *testing.T) {
	handler := New()

	content := "# Test Note\n\nContent here."

	result, err := handler.Stringify(map[string]any{}, content)
	if err != nil {
		t.Fatalf("Stringify() error = %v", err)
	}

	if result != content {
		t.Errorf("Result = %q, want %q", result, content)
	}
}

func TestHandler_ValidateValidFrontmatter(t *testing.T) {
	handler := New()

	fm := map[string]any{
		"title":   "Valid Title",
		"tags":    []string{"tag1", "tag2"},
		"count":   42,
		"enabled": true,
	}

	result := handler.Validate(fm)

	if !result.IsValid {
		t.Errorf("IsValid = false, want true. Errors: %v", result.Errors)
	}
	if len(result.Errors) != 0 {
		t.Errorf("Errors = %v, want empty", result.Errors)
	}
}

func TestHandler_ValidateInvalidFrontmatterWithFunction(t *testing.T) {
	handler := New()

	fm := map[string]any{
		"title":       "Invalid",
		"badFunction": func() string { return "not allowed" },
	}

	result := handler.Validate(fm)

	if result.IsValid {
		t.Error("IsValid = true, want false")
	}
	if len(result.Errors) == 0 {
		t.Error("Errors should not be empty")
	}

	// Check that error mentions functions
	hasError := false
	for _, err := range result.Errors {
		if strings.Contains(err, "Functions are not allowed") || strings.Contains(err, "Invalid YAML") {
			hasError = true
			break
		}
	}
	if !hasError {
		t.Errorf("Errors should mention functions or invalid YAML, got: %v", result.Errors)
	}
}

func TestHandler_UpdateFrontmatter(t *testing.T) {
	handler := New()

	content := `---
title: Old Title
tags: [old]
---

# Content

Some content here.`

	updates := map[string]any{
		"title":    "New Title",
		"modified": "2023-12-01",
	}

	result, err := handler.UpdateFrontmatter(content, updates)
	if err != nil {
		t.Fatalf("UpdateFrontmatter() error = %v", err)
	}

	if !strings.Contains(result, "title: New Title") {
		t.Error("Result should contain new title")
	}
	if !strings.Contains(result, "modified:") {
		t.Error("Result should contain modified field")
	}
	if !strings.Contains(result, "tags:") {
		t.Error("Result should preserve tags")
	}
	if !strings.Contains(result, "# Content") {
		t.Error("Result should contain content")
	}
}

func TestHandler_ExtractFrontmatter(t *testing.T) {
	handler := New()

	content := `---
title: Test
value: 42
---

Content here.`

	fm := handler.ExtractFrontmatter(content)

	if fm["title"] != "Test" {
		t.Errorf("Frontmatter[title] = %v, want %q", fm["title"], "Test")
	}
	if fm["value"] != 42 {
		t.Errorf("Frontmatter[value] = %v, want %v", fm["value"], 42)
	}
}
