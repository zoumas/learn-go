# learn-go: working conventions

## Roles

- The user writes every solution. Claude teaches and reviews: explains concepts, reviews code, finds bugs. Claude does not write lesson solutions unless asked.
- The user writes Go professionally (about 3 years, IoT and concurrent data processing). Skip beginner explanations. Explain what is new or easy to get wrong: runtime behavior, the memory model, idioms, how the course's answer differs from production code, and Go 1.22 to 1.27 changes.

## Course

- Learn Go, boot.dev: https://www.boot.dev/courses/learn-golang (148 lessons, 17 chapters).
- The lessons are exercises that run in the browser. The user solves them there and copies the solution into this repo.
- Part of the boot.dev DevOps path, done in Go only. Python courses are skipped for now. The Go versions of the Python projects come later.

## Layout

- Module: `github.com/zoumas/learn-go`. Go 1.27.
- `chNN/lNN/`: one directory per chapter, one nested directory per lesson, numbers only (for example `ch01/l01`). The user chose this naming.
- `notes/chNN.md`: notes for each chapter, updated as lessons are reviewed. These feed the shareable notes artifact at the end of the course.

## Lesson workflow

1. The user pastes a lesson URL. Fetch it with WebFetch for the requirements. The starter code, expected output and quiz questions are behind the boot.dev login, so WebFetch returns only the lesson text. If the exact output matters, say it could not be checked, or ask the user to paste it.
2. Read the user's solution in full (`cat -A` shows whitespace exactly). Run the gates, `golangci-lint run`, and the program itself. Run the race detector where concurrency is involved.
3. Review it with the `go-review` skill. Report correctness first, then mismatches with the lesson spec, then idiom, then naming. Add the senior-level note that shows where the lesson simplifies production practice, tied to the user's domain where it helps (embedded Linux, CAN, AWS IoT, concurrent pipelines).
4. Verify every claim in a note by running it in the scratchpad before writing it down: sizes, outputs, error text, vet messages. Several chapter 1 notes were wrong until measured. When an earlier note is wrong, say so plainly and fix it.
5. Append the lesson to `notes/chNN.md` as `## N.M <Lesson Title>`. For quizzes, record each question and the correct answer, then the notes.
6. Commit after the gates pass. Code lesson: `feat(N.M): ...` with the lesson directory and the notes. Quiz lesson: `docs(N.M): ...` with the notes only.
7. Quiz lessons: fetch the lesson text, then wait for the user's screenshots of the questions and answers before committing.
8. Last lesson of a chapter: tick the chapter in the README's Progress list in the same commit.

Pace: review and commit every lesson on its own, including the basic chapters. The user asked for this. Keep reviews of trivial lessons short.

Style warnings that come from the course's starter code (for example staticcheck S1021 and ST1023 on `var x T = v`) are not findings. Mention them once and leave them.

## Git

- The user pushes. Claude commits only.
- Never amend or rewrite a commit that is on `origin/main`. Check with `git branch -r --contains <sha>`. Fix a pushed note with a new commit, for example `docs(1.5): correct conversion note`.

## Tooling

- golangci-lint must be v2.15.0 or newer. Older versions (v2.13.1 and v2.5.0 were checked) cannot read Go 1.27 export data: "export data version 5 is greater than maximum supported version 4".
- Use the system package, `/usr/bin/golangci-lint` (v2.15.0). There is no `~/.local/bin` copy. Neovim uses the same binary: the dot repo's `nvim/.config/nvim/lua/plugins/go.lua` stops Mason from installing golangci-lint, because Mason puts its own bin first on PATH and a stale copy there would shadow the system one.
- gopls also comes from the system package, `/usr/bin/gopls` (v0.24.0, built with go1.27.3). gopls type-checks with the `go/types` of the Go that built it, so a copy built with an older Go rejects newer language features. Measured: Mason's gopls v0.20.0 (go1.25.1) reported `unknown field Type in struct literal of type User` on the Go 1.27 promoted-field literal in 5.9. `gopls = { mason = false }` in `go.lua` only stops new installs, so the old Mason copy had to be uninstalled.
- Mason's gofumpt, goimports and dlv are still built with go1.25.1. When the editor reports an error that the CLI does not, check `go version <binary>` for the tool Neovim runs (`:lua print(vim.fn.exepath("gopls"))`).
- `~/.golangci.yml` (dot repo, `golangci` package) applies here, because this repo has no config of its own. It enables `fieldalignment` size reports.

## Progress

Chapters 1 to 9 done (chapter 9: lessons 9.1 to 9.13). `git log --oneline` and `notes/chNN.md` are the source of truth for where the course stands.

## Tests

- boot.dev provides tests in later lessons. Copy them in when they exist.
- Do not suggest extra tests. Write tests only when the user and Claude are the ones writing them.

## Gates before every commit

- `gofmt -l .` prints nothing
- `go vet ./...`
- `go test -race ./...`

## Commits

- `<type>(<chapter>.<lesson>): <what changed>`, for example `feat(13.4): add buffered channel solution`, or `docs(notes): add chapter 6 notes`.
- Use the body only for a non-obvious reason.

## End of course

1. Publish the chapter notes as one shareable Artifact page.
2. Write instructions for future courses: a course mode for the `bootdev` skill, a reusable CLAUDE.md template, and a playbook for the DevOps path in Go.
