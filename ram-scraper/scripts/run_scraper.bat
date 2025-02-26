@echo off
cd /d "%~dp0.."

:: Definir rutas
set "config_file=config\settings.json"

:: Obtener timestamp de inicio
for /f %%a in ('wmic os get localdatetime ^| findstr /r "[0-9]"') do set datetime=%%a
set start_time=%time%
set timestamp=%datetime:~0,4%-%datetime:~4,2%-%datetime:~6,2% %datetime:~8,2%:%datetime:~10,2%:%datetime:~12,2%:%datetime:~15,3%


:: Extraer valores de configuración desde settings.json
for /f "delims=: tokens=2" %%a in ('findstr /C:"\"exe_name\"" "%config_file%"') do set "exe_name=%%a"
set "exe_name=%exe_name:~2,-2%"

for /f "delims=: tokens=2" %%a in ('findstr /C:"\"output_file\"" "%config_file%"') do set "output_file=%%a"
set "output_file=%output_file:~2,-2%"

for /f "delims=: tokens=2" %%a in ('findstr /C:"\"output_folder\"" "%config_file%"') do set "output_folder=%%a"
set "output_folder=%output_folder:~2,-2%"

for /f "delims=: tokens=2" %%a in ('findstr /C:"\"log_folder\"" "%config_file%"') do set "log_folder=%%a"
set "log_folder=%log_folder:~2,-2%"

for /f "delims=: tokens=2" %%a in ('findstr /C:"\"log_file\"" "%config_file%"') do set "log_file=%%a"
set "log_file=%log_file:~2,-1%"

:: Construir la ruta final del archivo de logs
set "log_path=%log_folder%bat/%log_file%"

:: Construir la ruta final del archivo de salida
set "output_path=%output_folder%bat/%output_file%"

:: Crear directorios si no existen
if not exist "%output_folder%bat" mkdir "%output_folder%bat"
if not exist "%log_folder%bat" mkdir "%log_folder%bat"


:: Iniciar log
echo %timestamp% - INFO: Inicio de ejecucion >> %log_path%

:: Verificar permisos de Administrador
net session >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo Solicitando permisos de Administrador...
    echo %timestamp% - INFO: Solicitando permisos de Administrador... >> %log_path%
    echo Set UAC = CreateObject^("Shell.Application"^) > "%temp%\admin.vbs"
    echo UAC.ShellExecute "%~s0", "", "", "runas", 1 >> "%temp%\admin.vbs"
    "%temp%\admin.vbs"
    exit /b
)

:: Registrar inicio de ejecución
echo %timestamp% - INFO: Ejecutando %exe_name% con salida en %output_path% >> %log_path%

:: Ejecutar el programa
bin\%exe_name% %output_path%

:: Obtener timestamp de finalización
set end_time=%time%

:: Calcular duración
for /f "tokens=1-4 delims=:." %%a in ("%start_time%") do set /a "start_ms=(((%%a*60+%%b)*60+%%c)*100)+%%d"
for /f "tokens=1-4 delims=:." %%a in ("%end_time%") do set /a "end_ms=(((%%a*60+%%b)*60+%%c)*100)+%%d"
set /a duration=(end_ms-start_ms)/100

:: Registrar fin de ejecución
echo %timestamp% - INFO: Ejecucion finalizada. Duracion: %duration% segundos. >> %log_path%
