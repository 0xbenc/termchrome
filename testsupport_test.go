package termchrome

import "github.com/0xbenc/termtheme"

// colorTheme is an inline test palette covering the roles the chrome widgets
// touch. termtheme ships no builtin palettes (they are app-local), so tests
// construct one directly.
func colorTheme() termtheme.Theme {
	return termtheme.Theme{
		Name: "test",
		Codes: map[termtheme.Role]string{
			termtheme.RoleBorder:     "1;32",
			termtheme.RoleMuted:      "2",
			termtheme.RoleForeground: "39",
			termtheme.RoleSuccess:    "32",
			termtheme.RoleWarning:    "33",
			termtheme.RoleDanger:     "31",
		},
	}
}

func plainTheme() termtheme.Theme {
	t := colorTheme()
	t.NoColor = true
	return t
}
