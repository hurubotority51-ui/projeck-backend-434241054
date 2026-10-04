$ErrorActionPreference = "Continue"

$repoRoot = (git rev-parse --show-toplevel).Trim()
if (-not $repoRoot) {
    throw "Tidak dapat menemukan root repository Git."
}
Set-Location $repoRoot

$watcher = New-Object System.IO.FileSystemWatcher
$watcher.Path = $repoRoot
$watcher.IncludeSubdirectories = $true
$watcher.NotifyFilter = [System.IO.NotifyFilters]::FileName -bor [System.IO.NotifyFilters]::LastWrite
$watcher.InternalBufferSize = 65536
$watcher.EnableRaisingEvents = $true

$createdSubscription = Register-ObjectEvent -InputObject $watcher -EventName Created -SourceIdentifier "AutoCommitNewFileCreated"
$renamedSubscription = Register-ObjectEvent -InputObject $watcher -EventName Renamed -SourceIdentifier "AutoCommitNewFileRenamed"

Write-Host "Memantau file baru di $repoRoot. Tekan Ctrl+C untuk berhenti."

try {
    while ($true) {
        $event = Wait-Event -Timeout 1
        if ($null -eq $event) {
            continue
        }

        $eventPath = $event.SourceEventArgs.FullPath
        Remove-Event -EventIdentifier $event.EventIdentifier

        if (-not (Test-Path -LiteralPath $eventPath -PathType Leaf)) {
            continue
        }

        $stableChecks = 0
        $previousLength = -1
        $previousWriteTime = [DateTime]::MinValue
        for ($attempt = 0; $attempt -lt 60 -and $stableChecks -lt 3; $attempt++) {
            if (-not (Test-Path -LiteralPath $eventPath -PathType Leaf)) {
                break
            }

            $file = Get-Item -LiteralPath $eventPath -ErrorAction SilentlyContinue
            if ($null -eq $file) {
                break
            }

            if ($file.Length -eq $previousLength -and $file.LastWriteTimeUtc -eq $previousWriteTime) {
                $stableChecks++
            } else {
                $stableChecks = 0
                $previousLength = $file.Length
                $previousWriteTime = $file.LastWriteTimeUtc
            }
            [System.Threading.Thread]::Sleep(500)
        }

        if ($stableChecks -lt 3 -or -not (Test-Path -LiteralPath $eventPath -PathType Leaf)) {
            continue
        }

        $relativePath = $eventPath.Substring($repoRoot.Length).TrimStart('\', '/') -replace '\\', '/'
        if ($relativePath -match '(^|/)(\.env($|\.)|.*\.(pem|key|p12|pfx|jks|keystore|sqlite|db|dump|bak)$|secrets?(/|$)|credentials?(/|$))') {
            Write-Host "Lewati file sensitif: $relativePath"
            continue
        }

        git check-ignore --quiet -- $relativePath
        if ($LASTEXITCODE -eq 0) {
            continue
        }

        $untracked = git ls-files --others --exclude-standard -- ":(literal)$relativePath"
        if ($LASTEXITCODE -ne 0 -or -not $untracked) {
            continue
        }

        git diff --cached --quiet
        if ($LASTEXITCODE -ne 0) {
            Write-Host "Lewati $relativePath karena ada perubahan staged lain."
            continue
        }

        git add -- ":(literal)$relativePath"
        if ($LASTEXITCODE -ne 0) {
            Write-Host "Gagal men-stage $relativePath."
            continue
        }

        git commit -m "Add new file: $relativePath" -- ":(literal)$relativePath"
        if ($LASTEXITCODE -eq 0) {
            Write-Host "File baru di-commit: $relativePath"
        } else {
            Write-Host "Gagal meng-commit $relativePath."
        }
    }
} finally {
    Unregister-Event -SourceIdentifier $createdSubscription.Name -ErrorAction SilentlyContinue
    Unregister-Event -SourceIdentifier $renamedSubscription.Name -ErrorAction SilentlyContinue
    $watcher.Dispose()
}
