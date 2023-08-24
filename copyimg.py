from io import BytesIO
import win32clipboard
from PIL import Image
import sys
import keyboard
import time

def send_to_clipboard(clip_type, data):
    win32clipboard.OpenClipboard()
    win32clipboard.EmptyClipboard()
    win32clipboard.SetClipboardData(clip_type, data)
    win32clipboard.CloseClipboard()

# filepath = 'Ico2.png'
# image = Image.open(filepath)

filepath = sys.argv[1]
#convert to png
jimage = Image.open(filepath)
jimage.save(filepath.replace('.jpg', '.png'))

image = Image.open(filepath.replace('.jpg', '.png'))

output = BytesIO()
image.convert("RGB").save(output, "BMP")
data = output.getvalue()[14:]
output.close()

send_to_clipboard(win32clipboard.CF_DIB, data)

keyboard.press_and_release('ctrl+v')
time.sleep(0.1)
keyboard.press_and_release('backspace')