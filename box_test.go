package termchrome

import (
	"strings"
	"testing"

	"github.com/0xbenc/termtheme"
)

// TestEdgeCanonicalStyling pins the resolved divergence: the fill dashes are
// border-styled in color mode, while NO_COLOR output is a clean exact string.
func TestEdgeCanonicalStyling(t *testing.T) {
	const width = 20
	wantPlain := "╭ TITLE ───────────╮" // ╭ + " TITLE " (7) + 11 dashes + ╮ == 20 cells

	if got := Edge(plainTheme(), "╭", "╮", "TITLE", width, nil); got != wantPlain {
		t.Fatalf("NO_COLOR edge = %q, want %q", got, wantPlain)
	}

	color := colorTheme()
	c := Edge(color, "╭", "╮", "TITLE", width, nil)
	if !strings.Contains(c, "\x1b[") {
		t.Fatalf("color edge has no styling: %q", c)
	}
	if got := termtheme.Strip(c); got != wantPlain {
		t.Fatalf("color edge strips to %q, want %q", got, wantPlain)
	}
	// The trailing dashes must be border-styled (the canonical choice), not left
	// default-colored as the old overlay fork did.
	borderDash := color.Style(termtheme.RoleBorder, "─")
	openSGR := borderDash[:strings.Index(borderDash, "m")+1]
	if !strings.Contains(c, openSGR+"─") {
		if !strings.Contains(termtheme.Strip(c[strings.LastIndex(c, "TITLE"):]), "─") {
			t.Fatalf("fill dashes are not border-styled: %q", c)
		}
	}
}

func TestEdgeWidths(t *testing.T) {
	plain := plainTheme()
	for _, w := range []int{8, 20, 48, 100} {
		for _, label := range []string{"", "MAP", "a-very-long-label-that-overflows"} {
			edge := Edge(plain, "╭", "╮", label, w, termtheme.Truncate)
			if got := termtheme.VisibleWidth(edge); got != w {
				t.Errorf("Edge width %d != %d (label=%q): %q", got, w, label, edge)
			}
		}
		if got := termtheme.VisibleWidth(Divider(plain, w)); got != w {
			t.Errorf("Divider width %d != %d", got, w)
		}
		if got := termtheme.VisibleWidth(Line(plain, "body", w, termtheme.Truncate)); got != w {
			t.Errorf("Line width %d != %d", got, w)
		}
	}
}

// TestTopBottom pins the rounded corner glyphs and width for the convenience
// wrappers.
func TestTopBottom(t *testing.T) {
	plain := plainTheme()
	top := Top(plain, "T", 10, termtheme.Truncate)
	if !strings.HasPrefix(top, "╭") || !strings.HasSuffix(top, "╮") {
		t.Errorf("Top corners wrong: %q", top)
	}
	bottom := Bottom(plain, 10)
	if !strings.HasPrefix(bottom, "╰") || !strings.HasSuffix(bottom, "╯") {
		t.Errorf("Bottom corners wrong: %q", bottom)
	}
	for _, line := range []string{top, bottom} {
		if w := termtheme.VisibleWidth(line); w != 10 {
			t.Errorf("width %d != 10: %q", w, line)
		}
	}
}
