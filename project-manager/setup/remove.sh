#!/bin/bash

printf "Eliminando project-manager..."

# Service
systemctl stop project-manager.service
systemctl disable project-manager.service
rm /etc/systemd/system/project-manager.service
rm /usr/lib/systemd/system/project-manager.service
systemctl daemon-reload

printf "Project manager eliminado"