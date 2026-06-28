// Package termchrome owns the shared, opinionated TUI chrome widgets that
// sibling terminal apps (passage, ssherpa, …) render through: rounded box
// geometry (Edge/Top/Bottom/Divider/Line), the canonical key-hint Footer, the
// aligned KVRow, plus the locale-aware GlyphSet (spinner + progress Bar) and the
// countdown UrgencyRole.
//
// It renders STRINGS ONLY over a termtheme.Theme. It depends on termtheme alone
// — no Bubble Tea, no os/net — so a box, footer, or countdown on a non-list
// screen never drags a navigation/runtime dependency in. List windowing lives in
// termnav; the per-app overflow policy (Strip vs Sanitize) stays in each app and
// is injected via the Truncator seam, never baked in here.
package termchrome
