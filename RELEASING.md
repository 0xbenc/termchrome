# Releasing

`termchrome` is the shared chrome-widget module (box/footer/kvrow + glyphs/countdown)
in a three-module shared-UI stack, consumed by
[passage](https://github.com/0xbenc/passage) and
[ssherpa](https://github.com/0xbenc/ssherpa):

```
termtheme (leaf; no bubbletea)
   ├─► termnav     (navigation / list windowing)
   ├─► termchrome  (this repo — depends on termtheme ONLY; no bubbletea/os)
   └─► consumed by ─► passage, ssherpa
```

Tag **bottom-up**: `termtheme` → `{termnav, termchrome}` → `{passage, ssherpa}`.
termchrome depends on termtheme only — if a change also needs a new termtheme, tag
termtheme **first**, pin it here (`go get github.com/0xbenc/termtheme@vX.Y.Z`), then
tag termchrome.

## Release order (when a change touches termchrome)

1. **Dev loop:** in each consumer add `replace github.com/0xbenc/termchrome => ../termchrome`,
   apply call-site changes, and get both green (`go build ./... && go test ./...`).
   ssherpa consumes termchrome through its `internal/chrome` shim (box/footer/kvrow) and
   directly for glyphs; passage through `renderWorkflowShell` + the `internal/termstyle` shim.
2. **Tag termchrome** (after both consumers are green locally):

   ```sh
   git push origin main
   git tag -a vX.Y.Z -m "termchrome vX.Y.Z — ..."
   git push origin vX.Y.Z
   ```

3. **Pin in each consumer** (one commit each):

   ```sh
   go get github.com/0xbenc/termchrome@vX.Y.Z
   go mod edit -dropreplace=github.com/0xbenc/termchrome
   go mod tidy && go test ./...
   git commit -am "<app>: pin termchrome vX.Y.Z (drop local replace)"
   ```

**No `replace` in any released `go.mod`** (`go mod verify` enforces it). **Pin lockstep:**
passage and ssherpa end on identical termtheme/termnav/termchrome versions (hotfix
exception: one app may bump ahead urgently; restore lockstep next release).

## Scope discipline

Keep only genuinely shared, app-agnostic primitives here. Per-app shell *composition*
(`renderWorkflowShell` / `workflowShell` step rail) and the overflow `Truncator` policy
(trusted chrome = `Sanitize`; raw transcript = `Strip`) stay in each app and are injected
— they must not migrate into termchrome.

## Versioning

Semantic versioning. A change to an exported signature or to rendered geometry/glyph
output is at least a **minor** bump; breaking one is a **major** bump.
