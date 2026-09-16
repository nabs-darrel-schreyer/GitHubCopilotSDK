cls
# Ensure module checksums exist (needed after a fresh clone / scaffold)
if (-not (Test-Path ".\go.sum")) {
  go mod tidy
}

# Go SDK does not bundle the CLI; point at a local install when unset.
if (-not $env:COPILOT_CLI_PATH) {
  $candidates = @(
    "$env:LOCALAPPDATA\GitHubCopilotCLI\copilot.exe",
    "$env:LOCALAPPDATA\GitHub CLI\copilot\copilot.exe",
    (Get-Command copilot -ErrorAction SilentlyContinue | Select-Object -ExpandProperty Source)
  )
  foreach ($c in $candidates) {
    if ($c -and (Test-Path -LiteralPath $c)) {
      $env:COPILOT_CLI_PATH = $c
      Write-Host "Using COPILOT_CLI_PATH=$env:COPILOT_CLI_PATH"
      break
    }
  }
  if (-not $env:COPILOT_CLI_PATH) {
    Write-Warning "COPILOT_CLI_PATH not set and no Copilot CLI found. Install GitHub Copilot CLI or set COPILOT_CLI_PATH."
  }
}

go run ./GoGhCopilot.Cli