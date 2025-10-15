@echo off
REM Installation script for terminal-trainer CLI on Windows
echo 🚀 Installing Terminal Trainer CLI...

REM Check if PowerShell is available
powershell -Command "Get-Host" >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ PowerShell is required but not found
    echo Please install PowerShell or download manually from:
    echo https://github.com/hungovercoders/terminal-trainer/releases
    pause
    exit /b 1
)

REM Run PowerShell installation script
echo 📦 Running PowerShell installation script...
powershell -ExecutionPolicy Bypass -Command "Invoke-RestMethod https://raw.githubusercontent.com/hungovercoders/terminal-trainer/main/install-windows.ps1 | Invoke-Expression"

if %errorlevel% equ 0 (
    echo.
    echo 🎉 Installation completed successfully!
    echo 💡 Please restart your command prompt to use 'terminal-trainer'
) else (
    echo.
    echo ❌ Installation failed
    echo Please try manual installation from:
    echo https://github.com/hungovercoders/terminal-trainer/releases
)

pause