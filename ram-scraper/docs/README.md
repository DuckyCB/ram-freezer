# RAM Scraper 🚀
**RAM Scraper** es un sistema para extraer y validar la memoria RAM de una computadora con Windows desde una Raspberry Pi con Linux.  

## 📁 Estructura del Proyecto
```plaintext
ram-scraper/
│── bin/                            # Ejecutables
│   ├── winpmem_mini_x64_rc2.exe    # Extrae la RAM en Windows
│── config/                         # Configuración
│   ├── settings.json               # Archivo con rutas y parámetros
│── data/                           # Archivos generados
│   ├── bat/                        # Archivos generados por .bat
│   │   ├── RAM_winpmem.raw         
│   ├── ps1/                        # Archivos generados por .ps1
│   │   ├── RAM_winpmem.raw
│── logs/                           # Registros de ejecución
│   ├── bat/                        
│   │   ├── run_scraper.log
│   ├── ps1/
│   │   ├── run_scraper.log
│── raspberry/                      # Código que corre en la Raspberry Pi
│   ├── validate_output.go          # Valida el archivo de RAM
│   ├── run_validation.sh           # Ejecuta el validador
│── scripts/                        # Scripts que corren en Windows
│   ├── run_scraper.ps1             # Ejecuta el scraper en PowerShell
│   ├── run_scraper.bat             # Ejecuta el scraper en CMD
│── README.md                       # Documentación del proyecto

## 🚀 Flujo de Trabajo
   1. La Raspberry ejecuta el script run_scraper.ps1 o run_scraper.bat en Windows.
   2. Windows ejecuta winpmem_mini_x64_rc2.exe y guarda la RAM en data/.
   3. La Raspberry copia RAM_winpmem.raw a su sistema.
   4. Ejecuta run_validation.sh, que:
      - Verifica el tamaño del archivo.
   5. Muestra el resultado (SUCCESS o ERROR).

## 🛠️ Uso
En la Raspberry Pi:
```bash
bash run_validation.sh
```

En Windows:
```powershell
.\run_scraper.ps1
```
o
```bash
run_scraper.bat
```

Para más detalles, revisa los archivos INSTALL.md y USAGE.md.

