import os
import shutil
import json
import io
print("Copying Codes to projects")
DirCode = "../../GoServer/"

for gofile in os.listdir('temp'):
    shutil.move('temp/' + gofile, DirCode + '/protocol/' + gofile)

print("Copying done!")
