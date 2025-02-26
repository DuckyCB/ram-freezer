#!/usr/bin/env bash

# Echo commands to stdout.
set -x
# Exit on first error.
set -e
# Treat undefined environment variables as errors.
set -u

# Start project-manager at boot
cp project-manager.service /lib/systemd/system/project-manager.service

systemctl daemon-reload
systemctl enable project-manager.service