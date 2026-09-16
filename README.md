# GitHubCopilotSDK

Multi-language learning repo that runs the **same GitHub Copilot SDK session flow** in **C# (.NET 10)**, **TypeScript**, **Python**, and **Go** — all on SDK **1.0.14**.

Each CLI creates a session (`model: auto`, approve-all permissions), sends a simple prompt, and prints the assistant reply in green:

```text
Calculate the product of 2 and 4.
```

Aimed at engineers comparing Copilot SDK ergonomics across languages for agent or CLI tooling.

## Samples

| Folder | Language | Package | Run |
| --- | --- | --- | --- |
| `gh-copilot-sdk-csharp` | C# / .NET 10 | `GitHub.Copilot.SDK` 1.0.14 | `.\run.ps1` |
| `gh-copilot-sdk-ts` | TypeScript | `@github/copilot-sdk` 1.0.14 | `npm install` then `.\run.ps1` |
| `gh-copilot-sdk-py` | Python 3.12+ | `github-copilot-sdk==1.0.14` | `.\run.ps1` (creates `.venv`) |
| `gh-copilot-sdk-go` | Go 1.24+ | `github.com/github/copilot-sdk/go` v1.0.14 | `.\run.ps1` |

## Prerequisites

- **GitHub Copilot CLI** installed and authenticated (signed-in user is enough for the default samples).
- Language runtimes as noted above.

### Go note

Unlike Node/.NET/Python, the Go SDK does **not** bundle a Copilot runtime. The Go sample resolves `COPILOT_CLI_PATH` (env → `LookPath` → common Windows installs such as `%LOCALAPPDATA%\GitHubCopilotCLI\copilot.exe`).

### Python note

`run.ps1` prefers a project `.venv` and skips the Microsoft Store `python` stub. On first run it installs requirements; the SDK caches its own runtime via `python -m copilot download-runtime` when needed.

## Stack

GitHub Copilot SDK · .NET · TypeScript · Python · Go · CLI

## Main idea

Keep one identical session story across four languages so you can compare APIs, packaging, and runtime wiring without changing the prompt or permission model.