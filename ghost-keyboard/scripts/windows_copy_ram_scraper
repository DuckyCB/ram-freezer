$origen = "D:\ram.exe"
$destino = "$env:USERPROFILE\Downloads\ram-scraper"

if (!(Test-Path -Path $destino)) {
    New-Item -ItemType Directory -Path $destino -Force
}

Copy-Item -Path $origen -Destination $destino -Force
