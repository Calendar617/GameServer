#!/bin/bash

if [ -f FilesForCSharp.txt ]; then
	allFilesTxt=`cat FilesForCSharp.txt`
else
	allFilesTxt=`find . | grep [.]sdp\$`
fi
files=""
for item in "${allFilesTxt}"
do
	item=`echo $item | tr -d '\r'`
	files="${files} ${item}"
done
mkdir -p temp
./SdpX/SdpX.exe  --template SdpX/templates/csharp.gotmpl --type_config SdpX/templates/csharp.tc --outdir ./temp --suffix cs --dir . --files ${files}

mkdir -p temp/lua
# for item in `ls ../.. | grep [.]sdp\$`
# do
 ../Tools/sdp2lua.exe --dir=./temp/lua/ *.sdp
# done

python ./CopyCodeToProject.py