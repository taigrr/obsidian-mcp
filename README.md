# obsidian-mcp

A Model Context Protocol (MCP) server that enables AI assistants to read, write, and search your Obsidian vault.

## Features

- **Full vault access** — Read, write, edit, delete, rename, and search notes
- **Frontmatter support** — Parse and update YAML frontmatter seamlessly
- **Security first** — Path traversal prevention, blocked system directories, safe defaults
- **Token optimized** — Compact JSON responses for efficient AI interactions
- **Universal compatibility** — Works with any MCP-compatible AI (Claude, ChatGPT, etc.)

## Installation

```bash
go install github.com/taigrr/obsidian-mcp/cmd/obsidian-mcp@latest
```

Or build from source:

```bash
git clone https://github.com/taigrr/obsidian-mcp.git
cd obsidian-mcp
go build -o obsidian-mcp ./cmd/obsidian-mcp
```

## Usage

```bash
obsidian-mcp /path/to/your/vault
```

### MCP Configuration

Add to your MCP client configuration:

```json
{
  "mcpServers": {
    "obsidian": {
      "command": "obsidian-mcp",
      "args": ["/path/to/your/vault"]
    }
  }
}
```

## Tools

| Tool      | Description                                                        |
| --------- | ------------------------------------------------------------------ |
| `read`    | Read a note with frontmatter and content. Supports pagination.     |
| `write`   | Create or overwrite a note with content and optional frontmatter.  |
| `edit`    | Replace text and/or update frontmatter fields in an existing note. |
| `delete`  | Delete a note (requires confirmation).                             |
| `rename`  | Move or rename a note to a new path.                               |
| `search`  | Full-text search with regex support. Returns matches with context. |
| `related` | Find notes related by tags or wiki-links.                          |
| `tags`    | List all unique tags across the vault (frontmatter and inline).    |

## Examples

### Reading a note

```json
{
  "tool": "read",
  "arguments": {
    "path": "notes/my-note.md"
  }
}
```

### Writing a note with frontmatter

```json
{
  "tool": "write",
  "arguments": {
    "path": "notes/new-note.md",
    "content": "# My Note\n\nThis is the content.",
    "frontmatter": {
      "tags": ["idea", "project"],
      "status": "draft"
    }
  }
}
```

### Searching the vault

```json
{
  "tool": "search",
  "arguments": {
    "query": "TODO",
    "contextLines": 2
  }
}
```

### Editing a note

```json
{
  "tool": "edit",
  "arguments": {
    "path": "notes/my-note.md",
    "oldText": "draft",
    "newText": "published",
    "frontmatter": {
      "status": "published"
    }
  }
}
```

## Security

- **Path filtering** — Blocks access to `.obsidian/`, `.git/`, `node_modules/`, and system files
- **File type restrictions** — Only allows `.md`, `.markdown`, and `.txt` files
- **Path traversal prevention** — All paths validated to stay within vault boundaries
- **Confirmation required** — Destructive operations require explicit confirmation

## Development

```bash
go build ./...        # Build
go test ./...         # Test
go vet ./...          # Lint
```

## License

[0BSD](LICENSE) — Tai Groot
