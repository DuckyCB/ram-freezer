#!/usr/bin/env python3
import os
import subprocess

def check_module_loaded(module_name):
    """Check if a kernel module is loaded."""
    try:
        output = subprocess.check_output(['lsmod']).decode()
        return module_name in output
    except subprocess.CalledProcessError as e:
        return False

def check_mountpoint(mount_point):
    """Check if a directory is mounted."""
    try:
        with open('/proc/mounts', 'r') as f:
            mounts = f.read()
        return mount_point in mounts
    except FileNotFoundError:
        return False

def check_device_exists(device_path):
    """Check if a device file exists."""
    return os.path.exists(device_path)

def check_permissions(device_path):
    """Check if the device file has the correct permissions."""
    try:
        mode = os.stat(device_path).st_mode
        return bool(mode & 0o222)  # Write permission for owner, group, or others
    except FileNotFoundError:
        return False

def run_checks():
    checks = [
        {"name": " - Module dwc2 loaded", "result": check_module_loaded("dwc2")},
        {"name": " - Module libcomposite loaded", "result": check_module_loaded("libcomposite")},
        {"name": " - ConfigFS mounted", "result": check_mountpoint("/sys/kernel/config")},
        {"name": " - /dev/hidg0 exists", "result": check_device_exists("/dev/hidg0")},
        {"name": " - /dev/hidg0 write permissions", "result": check_permissions("/dev/hidg0")},
    ]

    print("System Check Results:")
    for check in checks:
        status = "OK" if check["result"] else "FAIL"
        print(f"{check['name']}: {status}")

if __name__ == "__main__":
    run_checks()