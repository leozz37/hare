# Hare Sockets

Hare is a small Go module: a user-friendly TCP sockets library (root package `hare`) plus a CLI tool (`cmd/cli/cli.go`) for sending/listening on TCP ports. Standard commands live in the `README.md` and `Makefile`.

## Cursor Cloud specific instructions

- Go toolchain (1.22) is preinstalled; the update script only runs `go mod download`. There are no services to leave running — this is a library + one-shot CLI.
- Lint/test/build/run against the root package, which is what CI runs (`.github/workflows/go.yml` runs `go test ... .`):
  - Lint/vet: `go vet .` (use `.`, not `./...` — see gotcha below)
  - Test: `go test .` (or the coverage form in the README)
  - Build CLI: `make build` (produces `./hare-cli`)
  - Run CLI: `./hare-cli -l -p <port>` to listen, `./hare-cli -s -p <port> -d '<msg>'` to send.
- Gotcha — do not use `./...` for vet/test/build. The `examples/` package intentionally has two `main` functions (`examples/send.go` and `examples/listen.go`), and `cmd/cli/cli.go` has unused `fmt.Errorf` results, so `go vet ./...` / `go test ./...` fail on pre-existing issues unrelated to your change. The `Makefile` `test` target uses `./...` and will hit this; prefer `go test .`.
- Gotcha — the test suite binds fixed TCP ports 3000–3005. Do NOT leave a `hare-cli` listener (or anything else) bound to those ports while running tests, or `Listen()` returns an error and the tests nil-panic. Use a different port (e.g. 4000) for manual CLI demos.
- Gotcha — `TestListenStop` is pre-existing flaky (its `Stop()` only flips a flag while the accept loop is blocked on `Accept()`), so a single `go test .` run may intermittently fail on that test. Re-run to confirm; it is not caused by your changes.
