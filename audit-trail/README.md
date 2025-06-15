
# Audit Trail

## Log on bash

- import log.sh
- Log with corresponding level:
  - Debug: `log_debug "Message"`
  - Info: `log_info "Message"`
  - Warn: `log_warn "Message"`
  - Error: `log_error "Message"`
  - Fatal: `log_fatal "Message"`

```shell
source /opt/ram-freezer/audit-trail/log.sh

[...]

log_info "Los logs funcionan!"
```

> Esto guarda el log completo en `/opt/ram-freezer/bin/YYYY-MM-DD_X/ram-freezer.log`

### Install logs

Logs for the installation process

- import log.sh
- Log with corresponding level:
  - Debug: `log_install_debug "Message"`
  - Info: `log_install_info "Message"`
  - Warn: `log_install_warn "Message"`
  - Error: `log_install_error "Message"`
  - Fatal: `log_install_fatal "Message"`

```shell
source /opt/ram-freezer/audit-trail/log.sh

[...]

log_install_info "Los logs de instalación funcionan!"
```

> Esto guarda el log completo en `/opt/ram-freezer/bin/install/VERSION.log`

## Log on Go

[//]: # (TODO: Esto no funciona bien)

```go
package main

import (
	"ram-freezer/audit-trail/pkg/logger"
)

var Log *logger.SimpleLogger

func main() {
	Log, err := logger.NewRFLogger()
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}
	defer logs.Close()

	logs.Info("Los logs funcionan!")
}
```