# Uso de RAM Scraper
## 1️⃣ Ejecutar la extracción de RAM
Desde la PC Windows y ejecutar:
```powershell
run_scraper.ps1
```
o
```bash
run_scraper.bat'
```
Esto generará el archivo en `data/`.

## 2️⃣ Validar el archivo desde la Raspberry
Ejecutar:
```bash
bash run_validation.sh
```
Esto verificará: ✔️ Que el tamaño sea similar a la RAM del sistema
✔️ Que los primeros bytes sean válidos

# Posibles Errores
Error	Causa	Solución
Can not open SCM. Are you administrator?	Falta ejecutar como admin	Abrir CMD como administrador
No such file or directory	No existe RAM_winpmem.raw	Revisar settings.json
File size mismatch	Tamaño incorrecto	Verificar si el sistema tiene suficiente RAM
