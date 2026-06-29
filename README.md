# termchrome

> **Part of [termsystem](https://github.com/0xbenc/termsystem)** — the shared terminal-UI ecosystem (`termtheme` · `termnav` · `termchrome` · `termintro` powering `passage` · `ssherpa` · `dangit`). The ecosystem map, dependency graph, and the agent guide ([AGENTS.md](https://github.com/0xbenc/termsystem/blob/main/AGENTS.md)) live there.

Shared, opinionated TUI **chrome widgets** for sibling terminal apps
([passage](https://github.com/0xbenc/passage),
[ssherpa](https://github.com/0xbenc/ssherpa), …): rounded box geometry, a
canonical key-hint footer, aligned key/value rows, a locale-aware glyph set
(spinner + progress bar), and a countdown urgency ramp.

```
go get github.com/0xbenc/termchrome
```

Requires Go 1.26+. It renders **strings only** over a
[`termtheme.Theme`](https://github.com/0xbenc/termtheme) and depends on
`termtheme` **alone** — no Bubble Tea, no `os`/`net` — so drawing a box, footer,
or countdown on a non-list screen never drags a navigation or runtime dependency
in. List windowing lives in [`termnav`](https://github.com/0xbenc/termnav).

## What's here

- **Box geometry** — `Edge`/`Top`/`Bottom`/`Divider`/`Line` draw a rounded
  bordered shell. Fill dashes are always border-styled (the canonical choice).
  The per-app overflow policy (Strip vs Sanitize) is injected via the
  **`Truncator`** seam, never baked in — so the trusted-chrome path and a
  raw-transcript path can each keep their own policy while sharing the geometry.
- **`Footer` / `KeyHint` / `FooterSep`** — render key hints in one grammar
  (`"key label / key label"`) with progressive `+N` overflow instead of a silent
  `~` truncation.
- **`KVRow`** — an aligned `label   value` row with one gutter, one grammar.
- **`GlyphSet`** — `UnicodeGlyphs`/`ASCIIGlyphs`/`ResolveGlyphs(env)` pick a
  spinner + progress-bar cell set by locale (braille on UTF-8, ASCII otherwise),
  decoupled from color. `Frame(n)` cycles the spinner; `Bar(remaining,total,w)`
  draws a progress bar.
- **`UrgencyRole(remaining,total)`** — ramps a countdown success → warning →
  danger as it drains (danger in the last sixth, warning in the last half).

## Design

`termchrome` is the next extraction after `termtheme` (the must-agree theme
interchange core) and `termnav` (the navigation/list-windowing engine): the
opinionated *widgets* that compose over a theme. Apps keep their own shell
*composition* (which rows, which footer content, the wizard step rail) and their
own overflow `Truncator`; only the genuinely shared primitives live here.
