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
    $logLine = "$timestamp - $message"
    Add-Content -Path $logPath -Value $logLine -Encoding UTF8
}

# Obtener memoria total
$totalram = [math]::Round(((Get-WmiObject -Class Win32_ComputerSystem).TotalPhysicalMemory / 1GB), 2)

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

# Marcar el estado
Write-Log "INFO: Marcando el estado de la ejecución..."
$stateFile = ".\$($config.state_file)"

# rm state file if exists
if (Test-Path $stateFile) {
    Remove-Item $stateFile -Force
    Write-Log "INFO: Estado de la ejecución eliminado."
}


$state = [PSCustomObject]@{
    status        = "pending"
    start_time    = $startTime.ToString("yyyy-MM-dd HH:mm:ss:fff")
    end_time      = $null
    duration      = $null
    error_message = $null
    total_ram     = $totalram
}

$state.status = "running"

# Obtener el directorio donde se está ejecutando el script
$currentDir = Get-Location

# Construir la ruta completa para el archivo de estado
$stateFile = Join-Path -Path $currentDir -ChildPath "data\state.json"

# Verificar si el directorio existe, si no, crearlo
$dir = [System.IO.Path]::GetDirectoryName($stateFile)
if (-not (Test-Path -Path $dir)) {
    New-Item -ItemType Directory -Force -Path $dir
}

# Convertir el estado a JSON
$stateJson = $state | ConvertTo-Json

# Crear una codificación UTF-8 sin BOM
$utf8NoBom = New-Object System.Text.UTF8Encoding($false)

# Usar StreamWriter para escribir el archivo en UTF-8 sin BOM
$writer = New-Object System.IO.StreamWriter($stateFile, $false, $utf8NoBom)
$writer.Write($stateJson)
$writer.Close()
Write-Log "INFO: Estado de la ejecución marcado como 'running'."


# Ejecutar el programa
Write-Host "Ejecutando $($config.exe_name)..."
Write-Log "INFO: Ejecutando $($config.exe_name) con salida en $outputPath"
$output = & $exePath $outputPath 2>&1

Write-Log "INFO: Salida del programa"
$output | ForEach-Object { Write-Log "  $_" }

Write-Log "INFO: Codigo de salida del programa: $exitcode"

# Obtener timestamp de finalización
$endTime = Get-Date
$duration = ($endTime - $startTime).TotalSeconds

# Verificar si existe el archivo
if (Test-Path $outputPath) {
    Write-Log "INFO: Archivo de salida creado: $outputPath"
    Write-Log "INFO: Ejecucion exitosa."
    $state.status = "completed"
    $state.end_time = $endTime.ToString("yyyy-MM-dd HH:mm:ss:fff")
    $state.duration = $duration
}
else {
    Write-Log "ERROR: Error en la ejecucion. Codigo de salida: $LASTEXITCODE"
    $state.status = "error"
    $state.error_message = "Error en la ejecucion del programa."
    $state.end_time = $endTime.ToString("yyyy-MM-dd HH:mm:ss:fff")
    $state.duration = $duration
}

$state | ConvertTo-Json | Out-File -FilePath $stateFile -Force
Write-Log "INFO: Estado de la ejecución actualizado."

# Registrar fin de ejecución
Write-Log "INFO: Ejecucion finalizada. Duracion: $duration segundos."
