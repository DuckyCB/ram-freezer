#!/usr/bin/env bash

set -x
set -e
set -u

cd /opt/ram-freezer/project-manager/setup

# Start project-manager at boot
cp project-manager.service /lib/systemd/system/project-manager.service

systemctl daemon-reload
systemctl enable project-manager.service