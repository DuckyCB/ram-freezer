#!/usr/bin/env bash


trap 'handle_error $? "$BASH_COMMAND" ${LINENO}' ERR

function handle_error() {
  log_error "Error en script"

  # Opcionalmente finalizar el script
  # exit $exit_code
}

function _get_caller_info() {
  # Necesitamos subir dos niveles en la pila de llamadas:
  # 1. Esta función (_get_caller_info)
  # 2. La función log que la llama
  # 3. La función log_* que llama a log
  # 4. El punto real donde se llama a log_*
  local frame=3
  local caller_info

  caller_info=$(caller $frame 2>/dev/null || caller 2 2>/dev/null || caller 1 2>/dev/null || caller 0)

  echo "$caller_info"
}

function log() {
  local caller_info
  caller_info=$(_get_caller_info)
  local lineno
  lineno=$(echo "$caller_info" | awk '{print $1}')
  local source
  source=$(echo "$caller_info" | awk '{print $3}')
  #  local source
  #  source="${BASH_SOURCE[1]}"
  #  local lineno
  #  lineno="${BASH_LINENO[0]}"

  local level=$1
  local message=$2
  local timestamp
  timestamp=$(date -u +"%Y-%m-%dT%H:%M:%S.000Z")

  #  if [ -z "$source" ]; then
  #    source="bash"
  #  fi

  shift 2
  local fields=""
  while [[ $# -gt 0 ]]; do
    fields="$fields,\"$1\":\"$2\""
    shift 2
  done

  echo "$level: $message"

  today=$(date +%Y-%m-%d)
  echo "{\"timestamp\":\"$timestamp\",\"level\":\"$level\",\"message\":\"$message\",\"file\":\"$source\",\"line\":$lineno${fields}}" >> /opt/ram-freezer/bin/logs/"$today".log
}

function log_fatal() {
  log "fatal" "$@"
  exit 1
}

function log_error() {
  log "error" "$@"
}

function log_warn() {
  log "warn" "$@"
}

function log_info() {
  log "info" "$@"
}
