#!/usr/bin/env bash


trap 'handle_error $? "$BASH_COMMAND" ${LINENO}' ERR

function handle_error() {
  local command="$2"
  local error_message
  error_message=$(eval "$command" 2>&1 >/dev/null)

  log_error "$error_message"

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

function _get_log_path() {
    local log_type="$1"
    local log_path=""

    case "$log_type" in
        "install")
            if [ -z "${RF_VERSION+x}" ] || [ -z "${RF_VERSION}" ]; then
                RF_VERSION=$(git rev-parse --short HEAD 2>/dev/null)
                if [ -z "${RF_VERSION}" ]; then
                    RF_VERSION="unknown_version"
                fi
            fi
            log_path="/opt/ram-freezer/bin/install/${RF_VERSION}.log"
            ;;
        "general")
            local base_path_file="/opt/ram-freezer/.out"
            if [ -f "${base_path_file}" ]; then
                local base_path
                base_path=$(cat "${base_path_file}")
                log_path="${base_path}/ram-freezer.log"
            else
                echo "Warn: ${base_path_file} no encontrado." >&2
                log_path="/opt/ram-freezer/bin/ram-freezer.log"
            fi
            ;;
        *)
            echo "Error: Tipo de log '${log_type}' no reconocido. Use 'install' o 'general'." >&2
            log_path="/opt/ram-freezer/bin/ram-freezer.log"
            ;;
    esac

    echo "${log_path}"
}

function log() {
  local caller_info
  caller_info=$(_get_caller_info)
  local lineno
  lineno=$(echo "$caller_info" | awk '{print $1}')
  local source
  source=$(echo "$caller_info" | awk '{print $3}')

  local path
  path=$(_get_log_path "$1")
  local level=$2
  local message=$3
  local timestamp
  timestamp=$(date -u +"%Y-%m-%dT%H:%M:%S.%6NZ")

  #  if [ -z "$source" ]; then
  #    source="bash"
  #  fi

  shift 3
  local fields=""
  while [[ $# -gt 0 ]]; do
    fields="$fields,\"$1\":\"$2\""
    shift 2
  done

  if [[ ! -f "$path" ]]; then
      touch "$path"
  fi

  echo "$level: $message"

  echo "{\"timestamp\":\"$timestamp\",\"level\":\"$level\",\"message\":\"$message\",\"file\":\"$source\",\"line\":$lineno${fields}}" >> "$path"
}

function log_fatal() {
  log "general" "FATAL" "$@"
  exit 1
}

function log_error() {
  log "general" "ERROR" "$@"
}

function log_warn() {
  log "general" "WARN" "$@"
}

function log_info() {
  log "general" "INFO" "$@"
}

function log_debug() {
  log "general" "DEBUG" "$@"
}

function log_install_fatal() {
  log "install" "FATAL" "$@"
  exit 1
}

function log_install_error() {
  log "install" "ERROR" "$@"
}

function log_install_warn() {
  log "install" "WARN" "$@"
}

function log_install_info() {
  log "install" "INFO" "$@"
}

function log_install_debug() {
  log "install" "DEBUG" "$@"
}
