package termchrome

import (
	"strings"

	"github.com/0xbenc/termtheme"
)

// Truncator shortens a (possibly styled) string to width cells. Callers pass
// their own so the trusted-chrome policy (Sanitize on overflow) and the raw
// transcript-body policy (Strip on overflow, preserve on fit) are each preserved
// while the box geometry is shared.
type Truncator func(value string, width int) string

// Top draws the rounded top border with an optional label.
func Top(theme termtheme.Theme, label string, width int, tr Truncator) string {
	return Edge(theme, "╭", "╮", label, width, tr)
}

// Divider draws a mid-box horizontal rule.
func Divider(theme termtheme.Theme, width int) string {
	return Edge(theme, "├", "┤", "", width, nil)
}

// Bottom draws the rounded bottom border.
func Bottom(theme termtheme.Theme, width int) string {
	return Edge(theme, "╰", "╯", "", width, nil)
}

// Edge draws a top/divider/bottom border row. The fill dashes are always
// border-styled — the canonical choice that resolves the historical divergence
// between a styled picker border and a default-colored overlay border. The label
// is styled by the caller and truncated by tr; an empty label yields a plain
// rule.
func Edge(theme termtheme.Theme, left, right, label string, width int, tr Truncator) string {
	inner := max(0, width-2)
	if inner == 0 {
		return theme.Style(termtheme.RoleBorder, left+right)
	}
	label = strings.TrimSpace(label)
	if label == "" {
		return theme.Style(termtheme.RoleBorder, left+strings.Repeat("─", inner)+right)
	}
	if tr != nil {
		label = " " + tr(label, max(0, inner-2)) + " "
	} else {
		label = " " + label + " "
	}
	remaining := inner - termtheme.VisibleWidth(label)
	if remaining < 0 {
		remaining = 0
	}
	return theme.Style(termtheme.RoleBorder, left) + label + theme.Style(termtheme.RoleBorder, strings.Repeat("─", remaining)+right)
}

// Line wraps content as a box body row ("│ … │"), truncating with tr and padding
// to the inner width.
func Line(theme termtheme.Theme, content string, width int, tr Truncator) string {
	inner := max(0, width-4)
	content = termtheme.PadRight(tr(content, inner), inner)
	return theme.Style(termtheme.RoleBorder, "│ ") + content + theme.Style(termtheme.RoleBorder, " │")
}
