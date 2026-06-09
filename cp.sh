#!/bin/bash
#chmod u+x *.sh
#rm -rf bin
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./examples/ExcelToCsharp.go
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./examples/ExcelToCsharp.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./examples/ExcelToCsharp
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./examples/ExcelToCsharp
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./examples/ExcelToGo
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./examples/ExcelToGo
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./examples/ExcelToGoSys
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./examples/ExcelToGoSys
#输出路径bin
output_dir="/Users/yh/Documents/github/crazy-elimination/Assets/ExcelTo"

# 删除文件夹
#if [ -d "$output_dir" ]; then
#    echo "Deleting folder..."
#    rm -rf "$output_dir"
#else
#    echo "Folder does not exist."
#fi

# 创建文件夹
#echo "Creating folder..."
#mkdir "$output_dir"


#cp ExcelToCsharp "$output_dir/"
#cp ExcelToCsharp.exe "$output_dir/"

