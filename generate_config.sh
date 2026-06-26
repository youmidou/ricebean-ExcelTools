#!/bin/bash

#go_output_dir="./../../config/cfg_go"
#json_output_dir="./../../config/cfg_json"
go_output_dir="./config/cfg_go"
json_output_dir="./config/cfg_json"
# 这样传参就完全合法且能被 Go 正确接收了！
./ExcelToGo --go_output_dir="$go_output_dir" --json_output_dir="$json_output_dir"

echo "生成服务器配置文件..."