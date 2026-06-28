package termchrome

import "github.com/0xbenc/termtheme"

// KVRow renders an aligned "label   value" row: the label is muted and padded to
// gutter cells so values line up in a column, the value is foreground-styled. One
// gutter, one grammar — replacing the per-screen 7/8/9/13/14 drift.
func KVRow(theme termtheme.Theme, label, value string, gutter int) string {
	return theme.Style(termtheme.RoleMuted, termtheme.PadRight(label, gutter)) +
		theme.Style(termtheme.RoleForeground, value)
}
