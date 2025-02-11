
KEYCODE_NONE = 0
KEYCODE_A = 0x04
KEYCODE_B = 0x05
KEYCODE_C = 0x06
KEYCODE_D = 0x07
KEYCODE_E = 0x08
KEYCODE_F = 0x09
KEYCODE_G = 0x0a
KEYCODE_H = 0x0b
KEYCODE_I = 0x0c
KEYCODE_J = 0x0d
KEYCODE_K = 0x0e
KEYCODE_L = 0x0f
KEYCODE_M = 0x10
KEYCODE_N = 0x11
KEYCODE_O = 0x12
KEYCODE_P = 0x13
KEYCODE_Q = 0x14
KEYCODE_R = 0x15
KEYCODE_S = 0x16
KEYCODE_T = 0x17
KEYCODE_U = 0x18
KEYCODE_V = 0x19
KEYCODE_W = 0x1a
KEYCODE_X = 0x1b
KEYCODE_Y = 0x1c
KEYCODE_Z = 0x1d
KEYCODE_NUMBER_1 = 0x1e
KEYCODE_NUMBER_2 = 0x1f
KEYCODE_NUMBER_3 = 0x20
KEYCODE_NUMBER_4 = 0x21
KEYCODE_NUMBER_5 = 0x22
KEYCODE_NUMBER_6 = 0x23
KEYCODE_NUMBER_7 = 0x24
KEYCODE_NUMBER_8 = 0x25
KEYCODE_NUMBER_9 = 0x26
KEYCODE_NUMBER_0 = 0x27
KEYCODE_ENTER = 0x28
KEYCODE_ESCAPE = 0x29
KEYCODE_BACKSPACE_DELETE = 0x2a
KEYCODE_TAB = 0x2b
KEYCODE_SPACEBAR = 0x2c
KEYCODE_MINUS = 0x2d
KEYCODE_EQUAL_SIGN = 0x2e
KEYCODE_LEFT_BRACKET = 0x2f
KEYCODE_RIGHT_BRACKET = 0x30
KEYCODE_BACKSLASH = 0x31
KEYCODE_HASH = 0x32
KEYCODE_SEMICOLON = 0x33
KEYCODE_SINGLE_QUOTE = 0x34
KEYCODE_ACCENT_GRAVE = 0x35
KEYCODE_COMMA = 0x36
KEYCODE_PERIOD = 0x37
KEYCODE_FORWARD_SLASH = 0x38
KEYCODE_CAPS_LOCK = 0x39
KEYCODE_F1 = 0x3a
KEYCODE_F2 = 0x3b
KEYCODE_F3 = 0x3c
KEYCODE_F4 = 0x3d
KEYCODE_F5 = 0x3e
KEYCODE_F6 = 0x3f
KEYCODE_F7 = 0x40
KEYCODE_F8 = 0x41
KEYCODE_F9 = 0x42
KEYCODE_F10 = 0x43
KEYCODE_F11 = 0x44
KEYCODE_F12 = 0x45
KEYCODE_PRINT_SCREEN = 0x46
KEYCODE_SCROLL_LOCK = 0x47
KEYCODE_PAUSE_BREAK = 0x48
KEYCODE_INSERT = 0x49
KEYCODE_HOME = 0x4a
KEYCODE_PAGE_UP = 0x4b
KEYCODE_DELETE = 0x4c
KEYCODE_END = 0x4d
KEYCODE_PAGE_DOWN = 0x4e
KEYCODE_RIGHT_ARROW = 0x4f
KEYCODE_LEFT_ARROW = 0x50
KEYCODE_DOWN_ARROW = 0x51
KEYCODE_UP_ARROW = 0x52
KEYCODE_CLEAR = 0x53
KEYCODE_NUM_LOCK = 0x53
KEYCODE_NUMPAD_DIVIDE = 0x54
KEYCODE_NUMPAD_MULTIPLY = 0x55
KEYCODE_NUMPAD_MINUS = 0x56
KEYCODE_NUMPAD_PLUS = 0x57
KEYCODE_NUMPAD_ENTER = 0x58
KEYCODE_NUMPAD_1 = 0x59
KEYCODE_NUMPAD_2 = 0x5a
KEYCODE_NUMPAD_3 = 0x5b
KEYCODE_NUMPAD_4 = 0x5c
KEYCODE_NUMPAD_5 = 0x5d
KEYCODE_NUMPAD_6 = 0x5e
KEYCODE_NUMPAD_7 = 0x5f
KEYCODE_NUMPAD_8 = 0x60
KEYCODE_NUMPAD_9 = 0x61
KEYCODE_NUMPAD_0 = 0x62
KEYCODE_NUMPAD_DOT = 0x63
KEYCODE_102ND = 0x64  # Right of left Shift on non-US keyboards
KEYCODE_CONTEXT_MENU = 0x65
KEYCODE_F13 = 0x68
KEYCODE_F14 = 0x69
KEYCODE_F15 = 0x6a
KEYCODE_F16 = 0x6b
KEYCODE_F17 = 0x6c
KEYCODE_F18 = 0x6d
KEYCODE_F19 = 0x6e
KEYCODE_F20 = 0x6f
KEYCODE_F21 = 0x70
KEYCODE_F22 = 0x71
KEYCODE_F23 = 0x72
KEYCODE_EXECUTE = 0x74
KEYCODE_HELP = 0x75
KEYCODE_SELECT = 0x77
KEYCODE_INTL_RO = 0x87
KEYCODE_INTL_YEN = 0x89
KEYCODE_HANGEUL = 0x90
KEYCODE_HANJA = 0x91
KEYCODE_LEFT_CTRL = 0xe0
KEYCODE_LEFT_SHIFT = 0xe1
KEYCODE_LEFT_ALT = 0xe2
KEYCODE_LEFT_META = 0xe3
KEYCODE_RIGHT_CTRL = 0xe4
KEYCODE_RIGHT_SHIFT = 0xe5
KEYCODE_RIGHT_ALT = 0xe6
KEYCODE_RIGHT_META = 0xe7
KEYCODE_MEDIA_PLAY_PAUSE = 0xe8
KEYCODE_REFRESH = 0xfa

_MODIFIER_KEYCODES = [
    'AltLeft',
    'AltRight',
    'ControlLeft',
    'ControlRight',
    'MetaLeft',
    'MetaRight',
    'ShiftLeft',
    'ShiftRight',
]

_MAPPING = {
    'AltLeft': hid.KEYCODE_LEFT_ALT,
    'AltRight': hid.KEYCODE_RIGHT_ALT,
    'ArrowDown': hid.KEYCODE_DOWN_ARROW,
    'ArrowLeft': hid.KEYCODE_LEFT_ARROW,
    'ArrowRight': hid.KEYCODE_RIGHT_ARROW,
    'ArrowUp': hid.KEYCODE_UP_ARROW,
    'Backquote': hid.KEYCODE_ACCENT_GRAVE,
    'Backslash': hid.KEYCODE_BACKSLASH,
    'Backspace': hid.KEYCODE_BACKSPACE_DELETE,
    'BracketLeft': hid.KEYCODE_LEFT_BRACKET,
    'BracketRight': hid.KEYCODE_RIGHT_BRACKET,
    'CapsLock': hid.KEYCODE_CAPS_LOCK,
    'Comma': hid.KEYCODE_COMMA,
    'ContextMenu': hid.KEYCODE_CONTEXT_MENU,
    'ControlLeft': hid.KEYCODE_LEFT_CTRL,
    'ControlRight': hid.KEYCODE_RIGHT_CTRL,
    'Delete': hid.KEYCODE_DELETE,
    'Digit0': hid.KEYCODE_NUMBER_0,
    'Digit1': hid.KEYCODE_NUMBER_1,
    'Digit2': hid.KEYCODE_NUMBER_2,
    'Digit3': hid.KEYCODE_NUMBER_3,
    'Digit4': hid.KEYCODE_NUMBER_4,
    'Digit5': hid.KEYCODE_NUMBER_5,
    'Digit6': hid.KEYCODE_NUMBER_6,
    'Digit7': hid.KEYCODE_NUMBER_7,
    'Digit8': hid.KEYCODE_NUMBER_8,
    'Digit9': hid.KEYCODE_NUMBER_9,
    'End': hid.KEYCODE_END,
    'Enter': hid.KEYCODE_ENTER,
    'Equal': hid.KEYCODE_EQUAL_SIGN,
    'Escape': hid.KEYCODE_ESCAPE,
    'F1': hid.KEYCODE_F1,
    'F2': hid.KEYCODE_F2,
    'F3': hid.KEYCODE_F3,
    'F4': hid.KEYCODE_F4,
    'F5': hid.KEYCODE_F5,
    'F6': hid.KEYCODE_F6,
    'F7': hid.KEYCODE_F7,
    'F8': hid.KEYCODE_F8,
    'F9': hid.KEYCODE_F9,
    'F10': hid.KEYCODE_F10,
    'F11': hid.KEYCODE_F11,
    'F12': hid.KEYCODE_F12,
    'F13': hid.KEYCODE_F13,
    'F14': hid.KEYCODE_F14,
    'F15': hid.KEYCODE_F15,
    'F16': hid.KEYCODE_F16,
    'F17': hid.KEYCODE_F17,
    'F18': hid.KEYCODE_F18,
    'F19': hid.KEYCODE_F19,
    'F20': hid.KEYCODE_F20,
    'F21': hid.KEYCODE_F21,
    'F22': hid.KEYCODE_F22,
    'F23': hid.KEYCODE_F23,
    'Home': hid.KEYCODE_HOME,
    'Insert': hid.KEYCODE_INSERT,
    'IntlBackslash': hid.KEYCODE_102ND,
    'IntlRo': hid.KEYCODE_INTL_RO,
    'IntlYen': hid.KEYCODE_INTL_YEN,
    'KeyA': hid.KEYCODE_A,
    'KeyB': hid.KEYCODE_B,
    'KeyC': hid.KEYCODE_C,
    'KeyD': hid.KEYCODE_D,
    'KeyE': hid.KEYCODE_E,
    'KeyF': hid.KEYCODE_F,
    'KeyG': hid.KEYCODE_G,
    'KeyH': hid.KEYCODE_H,
    'KeyI': hid.KEYCODE_I,
    'KeyJ': hid.KEYCODE_J,
    'KeyK': hid.KEYCODE_K,
    'KeyL': hid.KEYCODE_L,
    'KeyM': hid.KEYCODE_M,
    'KeyN': hid.KEYCODE_N,
    'KeyO': hid.KEYCODE_O,
    'KeyP': hid.KEYCODE_P,
    'KeyQ': hid.KEYCODE_Q,
    'KeyR': hid.KEYCODE_R,
    'KeyS': hid.KEYCODE_S,
    'KeyT': hid.KEYCODE_T,
    'KeyU': hid.KEYCODE_U,
    'KeyV': hid.KEYCODE_V,
    'KeyW': hid.KEYCODE_W,
    'KeyX': hid.KEYCODE_X,
    'KeyY': hid.KEYCODE_Y,
    'KeyZ': hid.KEYCODE_Z,
    'MetaLeft': hid.KEYCODE_LEFT_META,
    'MetaRight': hid.KEYCODE_RIGHT_META,
    'Minus': hid.KEYCODE_MINUS,
    'Numpad0': hid.KEYCODE_NUMPAD_0,
    'Numpad1': hid.KEYCODE_NUMPAD_1,
    'Numpad2': hid.KEYCODE_NUMPAD_2,
    'Numpad3': hid.KEYCODE_NUMPAD_3,
    'Numpad4': hid.KEYCODE_NUMPAD_4,
    'Numpad5': hid.KEYCODE_NUMPAD_5,
    'Numpad6': hid.KEYCODE_NUMPAD_6,
    'Numpad7': hid.KEYCODE_NUMPAD_7,
    'Numpad8': hid.KEYCODE_NUMPAD_8,
    'Numpad9': hid.KEYCODE_NUMPAD_9,
    'NumpadMultiply': hid.KEYCODE_NUMPAD_MULTIPLY,
    'NumpadAdd': hid.KEYCODE_NUMPAD_PLUS,
    'NumpadSubtract': hid.KEYCODE_NUMPAD_MINUS,
    'NumpadDecimal': hid.KEYCODE_NUMPAD_DOT,
    'NumpadDivide': hid.KEYCODE_NUMPAD_DIVIDE,
    'NumpadEnter': hid.KEYCODE_NUMPAD_ENTER,
    'NumLock': hid.KEYCODE_NUM_LOCK,
    'OSLeft': hid.KEYCODE_LEFT_META,
    'OSRight': hid.KEYCODE_RIGHT_META,
    'PageUp': hid.KEYCODE_PAGE_UP,
    'PageDown': hid.KEYCODE_PAGE_DOWN,
    'Pause': hid.KEYCODE_PAUSE_BREAK,
    'Period': hid.KEYCODE_PERIOD,
    'PrintScreen': hid.KEYCODE_PRINT_SCREEN,
    'Quote': hid.KEYCODE_SINGLE_QUOTE,
    'ScrollLock': hid.KEYCODE_SCROLL_LOCK,
    'Select': hid.KEYCODE_SELECT,
    'Semicolon': hid.KEYCODE_SEMICOLON,
    'ShiftLeft': hid.KEYCODE_LEFT_SHIFT,
    'ShiftRight': hid.KEYCODE_RIGHT_SHIFT,
    'Space': hid.KEYCODE_SPACEBAR,
    'Slash': hid.KEYCODE_FORWARD_SLASH,
    'Tab': hid.KEYCODE_TAB,
}

class Keystroke:
    keycode: int
    modifier: int = KEYCODE_NONE

def convert(keystroke):
    return Keystroke(keycode=_map_keycode(keystroke),
                         modifier=_map_modifier_keys(keystroke))

def _map_modifier_keys(keystroke):
    modifier_bitmask = 0

    if keystroke.left_ctrl_modifier:
        modifier_bitmask |= hid.MODIFIER_LEFT_CTRL
    if keystroke.right_ctrl_modifier:
        modifier_bitmask |= hid.MODIFIER_RIGHT_CTRL

    if keystroke.left_shift_modifier:
        modifier_bitmask |= hid.MODIFIER_LEFT_SHIFT
    if keystroke.right_shift_modifier:
        modifier_bitmask |= hid.MODIFIER_RIGHT_SHIFT

    if keystroke.left_alt_modifier:
        modifier_bitmask |= hid.MODIFIER_LEFT_ALT
    if keystroke.right_alt_modifier:
        modifier_bitmask |= hid.MODIFIER_RIGHT_ALT

    if keystroke.left_meta_modifier:
        modifier_bitmask |= hid.MODIFIER_LEFT_META
    if keystroke.right_meta_modifier:
        modifier_bitmask |= hid.MODIFIER_RIGHT_META

    return modifier_bitmask

def _map_keycode(keystroke):
    if (keystroke.code in _MODIFIER_KEYCODES and
            _count_modifiers(keystroke) == 1):
        return hid.KEYCODE_NONE

    try:
        return _MAPPING[keystroke.code]
    except KeyError as e:
        raise UnrecognizedKeyCodeError(
            f'Unrecognized key code {keystroke.key} {keystroke.code}') from e


def _count_modifiers(keystroke):
    return (int(keystroke.left_ctrl_modifier) +
            int(keystroke.right_ctrl_modifier) +
            int(keystroke.left_shift_modifier) +
            int(keystroke.right_shift_modifier) +
            int(keystroke.left_alt_modifier) +
            int(keystroke.right_alt_modifier) +
            int(keystroke.left_meta_modifier) +
            int(keystroke.right_meta_modifier))
