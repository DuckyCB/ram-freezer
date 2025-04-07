
# Audit Trail

## Log on bash

- import log.sh
- Log with corresponding level:
  - Info: `log_info "Message"`
  - Warn: `log_warn "Message"`
  - Error: `log_error "Message"`
  - Fatal: `log_fatal "Message"`

```shell
source /opt/ram-freezer/audit-trail/log.sh

[...]

log_info "Los logs funcionan!"
```

> Esto guarda el log completo en `/opt/ram-freezer/bin/logs/YYYY-MM-DD.log`

## Log on Go

[//]: # (TODO: Esto no funciona bien)

```go
package main

import (
	"ram-freezer/audit-trail/pkg/logger"
)

func main() {
	logs, err := logger.NewSimpleLogger(logDir)
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}
	defer logs.Close()

	logs.Info("Los logs funcionan!")
}
```