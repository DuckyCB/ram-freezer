#!/bin/bash

base_dir=$(dirname $0)

type=$1

# Verificar si se especificó el tipo de archivo
if [ -z "$type" ]; then
    echo "ERROR: Debe especificar el tipo de archivo a validar."
    echo "Uso: $0 <tipo: bat|ps1>"
    exit 1
fi

# Definir rutas
FILE_PATH="$base_dir/../data/$type/RAM_winpmem.raw"
VALIDATOR_SOURCE="$base_dir/../raspberry/validate_output.go"
VALIDATOR_BINARY="$base_dir/../bin/validate_output"

# Verificar si el ejecutable compilado existe, si no, compilarlo
if [ ! -f "$VALIDATOR_BINARY" ]; then
    echo "Compilando el validador..."
    go build -o "$VALIDATOR_BINARY" "$VALIDATOR_SOURCE"
    if [ $? -ne 0 ]; then
        echo "ERROR: Fallo la compilación del validador."
        exit 1
    fi
fi

# Verificar si el archivo de salida existe antes de validarlo
if [ ! -f "$FILE_PATH" ]; then
    echo "ERROR: El archivo $FILE_PATH no existe."
    exit 1
fi

# Ejecutar la validación
echo "Ejecutando validación del archivo..."
"$VALIDATOR_BINARY" "$FILE_PATH"

# Capturar el código de salida
EXIT_CODE=$?

# Manejar resultados según el código de salida
case $EXIT_CODE in
    0) echo "SUCCESS: Validación completada correctamente." ;;
    1) echo "ERROR: El archivo no fue generado." ;;
    2) echo "WARNING: El archivo es demasiado pequeño para ser válido." ;;
    3) echo "WARNING: No se pudo abrir el archivo, puede estar corrupto." ;;
    4) echo "ERROR: No se pudo obtener la cantidad de RAM del sistema." ;;
    5) echo "WARNING: El archivo es más pequeño de lo esperado." ;;
    6) echo "WARNING: El archivo es más grande de lo esperado." ;;
    *) echo "ERROR: Código de salida desconocido ($EXIT_CODE)." ;;
esac

exit $EXIT_CODE
