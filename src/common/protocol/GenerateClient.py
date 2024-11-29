import os
import os.path
import subprocess
import shutil
import json
if os.path.exists('temp'):
    shutil.rmtree('temp')

os.mkdir('temp')
os.system("protoc-3.17.3-win64\\bin\\protoc.exe src\*.proto --go_out=./temp")

##os.mkdir('temp\\lua')
##for luafile in os.listdir('./src'):
##    if luafile.endswith('.proto'):
##        os.system("protoc-3.17.3-win64\\bin\\protoc.exe src\\" + luafile + " -o " + './temp/lua/' + luafile.replace(".proto", ".bytes"))

import CopyCodeToClient
#os.system("pause")
