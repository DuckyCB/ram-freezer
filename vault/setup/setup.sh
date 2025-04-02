#!/usr/bin/env bash

#echo "Verificando módulos necesarios..."
#if ! grep -q "dtoverlay=dwc2" /boot/config.txt; then
#    echo "Añadiendo dtoverlay=dwc2 a config.txt..."
#    echo "dtoverlay=dwc2" | sudo tee -a /boot/config.txt
#fi
#
#if ! grep -q "dwc2" /etc/modules; then
#    echo "Añadiendo dwc2 a modules..."
#    echo "dwc2" | sudo tee -a /etc/modules
#fi
#
#if ! grep -q "libcomposite" /etc/modules; then
#    echo "Añadiendo libcomposite a modules..."
#    echo "libcomposite" | sudo tee -a /etc/modules
#fi

# Virtual storage
echo "Creando imagen de disco de 1GB..."
dd if=/dev/zero of=/piusb.bin bs=1M count=1024

echo "Formateando imagen como FAT32..."
mkdosfs /piusb.bin -F 32 -I

# 5. Crear script para configurar USB gadget
echo "Creando script de configuración USB gadget..."
cat > /usr/local/bin/usb_gadget_setup.sh << 'EOL'
#!/bin/bash

# Configurar USB gadget
cd /sys/kernel/config/usb_gadget/
mkdir -p pi4
cd pi4

# Configurar USB IDs (usando IDs genéricos)
echo 0x1d6b > idVendor  # Linux Foundation
echo 0x0104 > idProduct # Multifunction Composite Gadget
echo 0x0100 > bcdDevice # v1.0.0
echo 0x0200 > bcdUSB    # USB2

# Configurar strings descriptivas
mkdir -p strings/0x409
echo "0123456789" > strings/0x409/serialnumber
echo "Raspberry Pi" > strings/0x409/manufacturer
echo "Pi4 USB Storage" > strings/0x409/product

# Configurar configuraciones
mkdir -p configs/c.1/strings/0x409
echo "Configuración USB" > configs/c.1/strings/0x409/configuration
echo 250 > configs/c.1/MaxPower

# Configurar la función mass_storage
mkdir -p functions/mass_storage.usb0
echo 1 > functions/mass_storage.usb0/stall
echo 0 > functions/mass_storage.usb0/lun.0/cdrom
echo 0 > functions/mass_storage.usb0/lun.0/ro
echo 0 > functions/mass_storage.usb0/lun.0/nofua
echo /piusb.bin > functions/mass_storage.usb0/lun.0/file

# Vincular configuración con la función
ln -s functions/mass_storage.usb0 configs/c.1/

# Activar gadget
ls /sys/class/udc > UDC
EOL

echo "Haciendo ejecutable el script..."
sudo chmod +x /usr/local/bin/usb_gadget_setup.sh

## 6. Crear servicio para iniciar automáticamente
#echo "Creando servicio systemd..."
#cat > /etc/systemd/system/usb-gadget.service << 'EOL'
#[Unit]
#Description=Configuración de USB gadget
#After=local-fs.target
#
#[Service]
#Type=oneshot
#ExecStart=/usr/local/bin/usb_gadget_setup.sh
#RemainAfterExit=yes
#
#[Install]
#WantedBy=multi-user.target
#EOL

## 7. Activar el servicio
#echo "Activando servicio..."
#sudo systemctl enable usb-gadget.service

## 8. Configurar modo OTG en cmdline.txt
#if ! grep -q "modules-load=dwc2" /boot/cmdline.txt; then
#    echo "Configurando cmdline.txt..."
#    sudo sed -i 's/$/ modules-load=dwc2/' /boot/cmdline.txt
#fi

echo "Configuración completada. Reinicia la Raspberry Pi para aplicar los cambios."
echo "Después del reinicio, conecta la Raspberry Pi a un ordenador mediante el puerto USB-C (OTG)"