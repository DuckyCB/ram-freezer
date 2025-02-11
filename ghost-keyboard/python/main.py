

def send_key(hid_device, keycode):
    """Envía un keycode HID a /dev/hidg0 simulando una tecla presionada y soltada."""
    hid_device.write(b'\x00' + keycode + b'\x00\x00\x00\x00\x00\x00')  # Presiona tecla
    time.sleep(0.05)  # Pequeño delay
    hid_device.write(b'\x00\x00\x00\x00\x00\x00\x00\x00')  # Suelta tecla

if __name__ == '__main__':
    text = """hola como estas?
    bien.
    """
    with open("/dev/hidg0", "wb") as hid:
        for char in text:
            if char in hid_map:
                send_key(hid, hid_map[char])
            else:
                print(f"Caracter no mapeado: {char}")