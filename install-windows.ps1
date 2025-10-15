# PowerShell installation script for terminal-trainer CLI
param(
    [string]$InstallPath = "$env:LOCALAPPDATA\terminal-trainer"
)

Write-Host "🚀 Installing Terminal Trainer CLI..." -ForegroundColor Green

# Detect architecture
$arch = if ([Environment]::Is64BitOperatingSystem) { "x86_64" } else { "i386" }
$os = "Windows"

try {
    # Get latest release
    Write-Host "🔍 Getting latest release..." -ForegroundColor Yellow
    $response = Invoke-RestMethod -Uri "https://api.github.com/repos/hungovercoders/terminal-trainer/releases/latest"
    $version = $response.tag_name
    
    Write-Host "📦 Latest version: $version" -ForegroundColor Blue
    
    # Download URL
    $fileName = "terminal-trainer_${os}_${arch}.zip"
    $downloadUrl = "https://github.com/hungovercoders/terminal-trainer/releases/download/${version}/${fileName}"
    
    Write-Host "⬇️  Downloading terminal-trainer ${version}..." -ForegroundColor Yellow
    
    # Create temp directory
    $tempDir = [System.IO.Path]::GetTempPath() + [System.Guid]::NewGuid()
    New-Item -ItemType Directory -Path $tempDir | Out-Null
    
    $zipPath = Join-Path $tempDir "terminal-trainer.zip"
    
    # Download
    Invoke-WebRequest -Uri $downloadUrl -OutFile $zipPath -ErrorAction Stop
    
    # Extract
    Write-Host "📦 Extracting..." -ForegroundColor Yellow
    Expand-Archive -Path $zipPath -DestinationPath $tempDir -Force
    
    # Install
    Write-Host "📍 Installing to $InstallPath..." -ForegroundColor Yellow
    if (!(Test-Path $InstallPath)) {
        New-Item -ItemType Directory -Path $InstallPath -Force | Out-Null
    }
    
    $exePath = Join-Path $tempDir "terminal-trainer.exe"
    $targetPath = Join-Path $InstallPath "terminal-trainer.exe"
    
    Copy-Item $exePath $targetPath -Force
    
    # Add to PATH
    Write-Host "🔗 Adding to PATH..." -ForegroundColor Yellow
    $currentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
    if ($currentPath -notlike "*$InstallPath*") {
        $newPath = "$currentPath;$InstallPath"
        [Environment]::SetEnvironmentVariable("PATH", $newPath, "User")
        Write-Host "   ✅ Added $InstallPath to user PATH" -ForegroundColor Green
    }
    
    # Cleanup
    Remove-Item $tempDir -Recurse -Force
    
    Write-Host ""
    Write-Host "🎉 Installation complete!" -ForegroundColor Green
    Write-Host ""
    Write-Host "You can now use the CLI with:" -ForegroundColor White
    Write-Host "  terminal-trainer --help" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "💡 Restart your terminal to use the command" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "🚀 Try it out:" -ForegroundColor White
    Write-Host "  terminal-trainer quiz --topic linux" -ForegroundColor Cyan
    Write-Host "  terminal-trainer cheat --topic git" -ForegroundColor Cyan
    
} catch {
    Write-Host "❌ Installation failed: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "Please download manually from: https://github.com/hungovercoders/terminal-trainer/releases" -ForegroundColor Yellow
    exit 1
}