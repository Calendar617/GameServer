import os
import os.path
import subprocess
import shutil
import json
if os.path.exists('temp'):
    shutil.rmtree('temp')

os.mkdir('temp')
os.system("protoc-3.17.3-win64\\bin\\protoc.exe src\*.proto --go_out=./temp")

import CopyCodeToProject
#os.system("pause")
