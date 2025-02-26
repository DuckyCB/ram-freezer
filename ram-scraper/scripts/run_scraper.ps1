# Verificar si el script se ejecuta como Administrador
$currentUser = [Security.Principal.WindowsIdentity]::GetCurrent()
$isAdmin = (New-Object Security.Principal.WindowsPrincipal $currentUser).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)

# Moverse a la raiz del proyecto (ram-scraper)
$rootPath = Split-Path -Parent $PSScriptRoot
Set-Location -Path $rootPath

# Cargar configuración
$type = "ps1"
$configPath = ".\config\settings.json"
$config = Get-Content $configPath | ConvertFrom-Json
$exePath = ".\bin\$($config.exe_name)"
$outputFolder = ".\$($config.output_folder)$type\"
$outputPath = "$outputFolder$($config.output_file)"
$logFolder = ".\$($config.log_folder)$type\"
$logPath = "$logFolder$($config.log_file)"

# Obtener timestamp de inicio
$startTime = Get-Date
$timestamp = $startTime.ToString("yyyy-MM-dd HH:mm:ss:fff")

# Función para escribir logs
function Write-Log {
    param ([string]$message)
    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    "$timestamp - $message" | Out-File -Append -FilePath $logPath
}

# Creando directorios si no existen
if (-Not (Test-Path $logFolder)) {
    New-Item -ItemType Directory -Path $logFolder
}
if (-Not (Test-Path $outputFolder)) {
    New-Item -ItemType Directory -Path $outputFolder
}

# Iniciar log
Write-Log "INFO: Inicio de ejecucion"

if (-Not $isAdmin) {
    Write-Host "Solicitando permisos de Administrador..."
    Write-Log "INFO: Solicitando permisos de Administrador..."
    Start-Process powershell -ArgumentList "-NoProfile -ExecutionPolicy Bypass -File `"$PSCommandPath`"" -Verb RunAs
    exit
}

# Ejecutar el programa
Write-Host "Ejecutando $($config.exe_name)..."
Write-Log "INFO: Ejecutando $($config.exe_name) con salida en $outputPath"
Start-Process -FilePath $exePath -ArgumentList $outputPath -NoNewWindow -Wait

# Obtener timestamp de finalización
$endTime = Get-Date
$duration = ($endTime - $startTime).TotalSeconds

# Registrar fin de ejecución
Write-Log "INFO: Ejecucion finalizada. Duracion: $duration segundos."
