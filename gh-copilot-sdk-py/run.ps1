cls

function Resolve-Python {
  # Prefer project venv
  $venvPy = Join-Path $PSScriptRoot ".venv\Scripts\python.exe"
  if (Test-Path -LiteralPath $venvPy) { return $venvPy }

  # Explicit known installs (skip WindowsApps Store stub)
  $candidates = @(
    "$env:LOCALAPPDATA\Programs\Python\Python312\python.exe",
    "$env:LOCALAPPDATA\Programs\Python\Python313\python.exe",
    "$env:LOCALAPPDATA\Programs\Python\Python311\python.exe",
    "$env:ProgramFiles\Python312\python.exe",
    "$env:ProgramFiles\Python313\python.exe",
    "C:\Python312\python.exe",
    "C:\Python311\python.exe"
  )
  foreach ($c in $candidates) {
    if (Test-Path -LiteralPath $c) { return $c }
  }

  # py launcher
  $py = Get-Command py -ErrorAction SilentlyContinue
  if ($py -and $py.Source -notmatch 'WindowsApps') {
    return $py.Source
  }

  # PATH python, but reject Microsoft Store alias
  foreach ($name in @('python', 'python3')) {
    $cmd = Get-Command $name -ErrorAction SilentlyContinue
    if ($cmd -and $cmd.Source -and ($cmd.Source -notmatch 'WindowsApps')) {
      return $cmd.Source
    }
  }
  return $null
}

$python = Resolve-Python
if (-not $python) {
  Write-Error @"
No real Python interpreter found.

The Microsoft Store 'python' stub is on PATH but is not a full install.
Install Python 3.12+ from https://www.python.org/downloads/ (check 'Add python.exe to PATH'),
or: winget install -e --id Python.Python.3.12

Then from this folder:
  python -m venv .venv
  .\.venv\Scripts\pip install -r requirements.txt
  .\run.ps1
"@
  exit 1
}

Write-Host "Using Python: $python"

# Create venv on first run if missing
$venvPy = Join-Path $PSScriptRoot ".venv\Scripts\python.exe"
if (-not (Test-Path -LiteralPath $venvPy)) {
  Write-Host "Creating .venv ..."
  & $python -m venv (Join-Path $PSScriptRoot ".venv")
  if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
  $venvPy = Join-Path $PSScriptRoot ".venv\Scripts\python.exe"
}

Write-Host "Ensuring dependencies ..."
& $venvPy -m pip install -q -r (Join-Path $PSScriptRoot "requirements.txt")
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

& $venvPy (Join-Path $PSScriptRoot "PyGhCopilot.Cli\main.py")