// md_to_pdf.go — Convert FIXES.md → FIXES.html (styled) → FIXES.pdf (headless Chrome)
//
// Usage:
//   go run md_to_pdf.go
//
// Requires: Google Chrome or Chromium installed.

package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const (
	mdFile   = "FIXES.md"
	htmlFile = "FIXES.html"
	pdfFile  = "FIXES.pdf"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// ── 1. Read Markdown ────────────────────────────────────────────────
	md, err := os.ReadFile(mdFile)
	if err != nil {
		slog.Error("read markdown", "err", err)
		os.Exit(1)
	}
	slog.Info("read markdown", "file", mdFile, "bytes", len(md))

	// ── 2. Convert Markdown → HTML body ─────────────────────────────────
	body := markdownToHTML(string(md))

	// ── 3. Wrap in full HTML with CSS ───────────────────────────────────
	html := wrapHTML(body)

	if err := os.WriteFile(htmlFile, []byte(html), 0644); err != nil {
		slog.Error("write html", "err", err)
		os.Exit(1)
	}
	slog.Info("wrote HTML", "file", htmlFile, "bytes", len(html))

	// ── 4. Find Chrome ──────────────────────────────────────────────────
	chromePath := findChrome()
	if chromePath == "" {
		slog.Error("Chrome/Chromium not found. Please install or set CHROME_PATH env var.")
		os.Exit(1)
	}
	slog.Info("found Chrome", "path", chromePath)

	// ── 5. Convert HTML → PDF via headless Chrome ───────────────────────
	absHTML, _ := abs(htmlFile)
	absPDF, _ := abs(pdfFile)

	cmd := exec.Command(chromePath,
		"--headless",
		"--disable-gpu",
		"--no-sandbox",
		"--print-to-pdf="+absPDF,
		"--print-to-pdf-no-header",
		"--no-pdf-header-footer",
		"file://"+absHTML,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		slog.Error("chrome PDF generation failed", "err", err)
		os.Exit(1)
	}

	info, _ := os.Stat(pdfFile)
	slog.Info("wrote PDF", "file", pdfFile, "bytes", info.Size())
	fmt.Printf("\n✅ Generated %s (%d KB)\n", pdfFile, info.Size()/1024)
}

// ─── Markdown → HTML converter (stdlib only, handles our subset) ────────────

func markdownToHTML(md string) string {
	lines := strings.Split(md, "\n")
	var out strings.Builder
	inTable := false
	inUL := false
	headerRowDone := false

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		// Close lists if not a list item
		if inUL && !strings.HasPrefix(trimmed, "- ") && trimmed != "" {
			out.WriteString("</ul>\n")
			inUL = false
		}

		// Blank line
		if trimmed == "" {
			if inTable {
				out.WriteString("</tbody></table>\n")
				inTable = false
				headerRowDone = false
			}
			continue
		}

		// Horizontal rule
		if trimmed == "---" && !inTable {
			out.WriteString("<hr>\n")
			continue
		}

		// Headings
		if strings.HasPrefix(trimmed, "# ") {
			level := 0
			for _, c := range trimmed {
				if c == '#' {
					level++
				} else {
					break
				}
			}
			text := strings.TrimSpace(trimmed[level:])
			text = inlineFormat(text)
			out.WriteString(fmt.Sprintf("<h%d>%s</h%d>\n", level, text, level))
			continue
		}

		// Table row
		if strings.HasPrefix(trimmed, "|") && strings.HasSuffix(trimmed, "|") {
			// Check if separator row
			isSep := true
			cells := splitTableRow(trimmed)
			for _, c := range cells {
				c = strings.TrimSpace(c)
				if c != "" && !isSepCell(c) {
					isSep = false
					break
				}
			}

			if !inTable {
				out.WriteString("<table>\n<thead>\n")
				inTable = true
				headerRowDone = false
			}

			if isSep {
				// Separator: close thead, open tbody
				out.WriteString("</thead>\n<tbody>\n")
				headerRowDone = true
				continue
			}

			tag := "td"
			if !headerRowDone {
				tag = "th"
			}

			out.WriteString("<tr>")
			for _, cell := range cells {
				cell = strings.TrimSpace(cell)
				cell = inlineFormat(cell)
				out.WriteString(fmt.Sprintf("<%s>%s</%s>", tag, cell, tag))
			}
			out.WriteString("</tr>\n")
			continue
		}

		// Unordered list
		if strings.HasPrefix(trimmed, "- ") {
			if !inUL {
				out.WriteString("<ul>\n")
				inUL = true
			}
			text := inlineFormat(trimmed[2:])
			out.WriteString(fmt.Sprintf("<li>%s</li>\n", text))
			continue
		}

		// Paragraph
		out.WriteString(fmt.Sprintf("<p>%s</p>\n", inlineFormat(trimmed)))
	}

	if inTable {
		out.WriteString("</tbody></table>\n")
	}
	if inUL {
		out.WriteString("</ul>\n")
	}

	return out.String()
}

func splitTableRow(line string) []string {
	// Remove leading/trailing |
	line = strings.TrimPrefix(line, "|")
	line = strings.TrimSuffix(line, "|")
	return strings.Split(line, "|")
}

func isSepCell(s string) bool {
	s = strings.TrimSpace(s)
	for _, c := range s {
		if c != '-' && c != ':' {
			return false
		}
	}
	return len(s) > 0
}

var (
	reBold      = regexp.MustCompile(`\*\*(.+?)\*\*`)
	reCode      = regexp.MustCompile("`([^`]+)`")
	reBoldFirst = regexp.MustCompile(`\*\*`)
	reItalic    = regexp.MustCompile(`\*(.+?)\*`)
	reLineBreak = regexp.MustCompile(`<br>`)
)

func inlineFormat(s string) string {
	// Escape bare HTML angle brackets (but keep <br>)
	s = strings.ReplaceAll(s, "<br>", "\x00BR\x00")
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\x00BR\x00", "<br>")

	// Bold **text**
	s = reBold.ReplaceAllString(s, "<strong>$1</strong>")

	// Inline code `text`
	s = reCode.ReplaceAllString(s, "<code>$1</code>")

	// Italic *text* (after bold is consumed)
	s = reItalic.ReplaceAllString(s, "<em>$1</em>")

	return s
}

func wrapHTML(body string) string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>FIXES REPORT — to-review.csv (PRU-15 2022 Sarawak)</title>
<style>
  @page {
    size: A4 landscape;
    margin: 12mm;
  }
  * { box-sizing: border-box; }
  body {
    font-family: -apple-system, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
    font-size: 9pt;
    line-height: 1.45;
    color: #1a1a1a;
    max-width: 100%;
    margin: 0 auto;
    padding: 10px;
  }
  h1 {
    font-size: 16pt;
    border-bottom: 2px solid #2c3e50;
    padding-bottom: 6px;
    margin-top: 0;
    color: #2c3e50;
  }
  h2 {
    font-size: 13pt;
    color: #2c3e50;
    border-bottom: 1px solid #bdc3c7;
    padding-bottom: 4px;
    margin-top: 18px;
  }
  h3 {
    font-size: 11pt;
    color: #34495e;
    margin-top: 14px;
  }
  hr {
    border: none;
    border-top: 1px solid #ddd;
    margin: 14px 0;
  }
  table {
    border-collapse: collapse;
    width: 100%;
    margin: 8px 0;
    font-size: 8pt;
    page-break-inside: auto;
  }
  tr {
    page-break-inside: avoid;
    page-break-after: auto;
  }
  th, td {
    border: 1px solid #bdc3c7;
    padding: 3px 6px;
    text-align: left;
    vertical-align: top;
    word-wrap: break-word;
    max-width: 420px;
  }
  th {
    background-color: #2c3e50;
    color: white;
    font-weight: 600;
    font-size: 8pt;
  }
  tbody tr:nth-child(even) {
    background-color: #f7f9fa;
  }
  tbody tr:hover {
    background-color: #eaf2f8;
  }
  code {
    background: #f0f0f0;
    padding: 1px 4px;
    border-radius: 3px;
    font-family: "SF Mono", "Consolas", "Monaco", monospace;
    font-size: 7.5pt;
    color: #c0392b;
    word-break: break-all;
  }
  strong { color: #2c3e50; }
  ul {
    margin: 4px 0;
    padding-left: 20px;
  }
  li { margin: 2px 0; }
  p { margin: 6px 0; }
  em { color: #7f8c8d; }
</style>
</head>
<body>
` + body + `
</body>
</html>`
}

func findChrome() string {
	// Check env override
	if p := os.Getenv("CHROME_PATH"); p != "" {
		return p
	}

	// macOS paths
	candidates := []string{
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		"/Applications/Chromium.app/Contents/MacOS/Chromium",
		"/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary",
		"/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
	}
	// Also check PATH
	for _, name := range []string{"google-chrome", "google-chrome-stable", "chromium", "chromium-browser"} {
		if p, err := exec.LookPath(name); err == nil {
			candidates = append([]string{p}, candidates...)
		}
	}

	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return ""
}

func abs(path string) (string, error) {
	if strings.HasPrefix(path, "/") {
		return path, nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return path, err
	}
	return wd + "/" + path, nil
}
