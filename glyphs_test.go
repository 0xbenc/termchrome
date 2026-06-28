package termchrome

import "testing"

// TestASCIIGlyphsAreSevenBit guards the fallback promise: every rune in the
// ASCII glyph set is <= 0x7e so it renders on a legacy codepage terminal.
func TestASCIIGlyphsAreSevenBit(t *testing.T) {
	g := ASCIIGlyphs()
	all := append([]string{g.BarFull, g.BarEmpty}, g.Spinner...)
	for _, s := range all {
		for _, r := range s {
			if r > 0x7e {
				t.Fatalf("ASCII glyph %q contains non-7-bit rune %U", s, r)
			}
		}
	}
}

func TestResolveGlyphsFromLocale(t *testing.T) {
	cases := []struct {
		name      string
		env       []string
		wantASCII bool
	}{
		{"utf8 lang", []string{"LANG=en_US.UTF-8"}, false},
		{"utf8 lc_all wins", []string{"LC_ALL=C.UTF-8", "LANG=C"}, false},
		{"posix", []string{"LANG=C"}, true},
		{"lc_ctype non-utf8 wins over lang", []string{"LC_CTYPE=POSIX", "LANG=en_US.UTF-8"}, true},
		{"empty", []string{}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ResolveGlyphs(tc.env)
			if got.ASCII != tc.wantASCII {
				t.Fatalf("ResolveGlyphs(%v).ASCII = %v, want %v", tc.env, got.ASCII, tc.wantASCII)
			}
		})
	}
}

// TestResolveGlyphsUTF8FrameIsBraille gates the visible spinner change: on a
// UTF-8 locale ResolveGlyphs yields the braille frame (not ASCII), and that
// frame is exactly one cell wide so it can never tear a spinner line.
func TestResolveGlyphsUTF8FrameIsBraille(t *testing.T) {
	g := ResolveGlyphs([]string{"LANG=en_US.UTF-8"})
	if g.ASCII {
		t.Fatalf("UTF-8 locale resolved to ASCII glyphs")
	}
	f := g.Frame(0)
	if f != "⠋" {
		t.Fatalf("UTF-8 Frame(0) = %q, want braille ⠋", f)
	}
	if r := []rune(f); len(r) != 1 || r[0] < 0x2800 || r[0] > 0x28ff {
		t.Fatalf("UTF-8 Frame(0) %q is not a single braille rune", f)
	}
}

func TestFrameCycles(t *testing.T) {
	g := ASCIIGlyphs()
	if g.Frame(0) != "|" || g.Frame(1) != "/" || g.Frame(4) != "|" {
		t.Fatalf("Frame cycling wrong: %q %q %q", g.Frame(0), g.Frame(1), g.Frame(4))
	}
	if (GlyphSet{}).Frame(3) != "" {
		t.Fatal("empty glyph set Frame should be empty string")
	}
}
