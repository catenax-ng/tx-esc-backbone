#!/bin/bash
# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone
#
# SPDX-License-Identifier: Apache-2.0

#!/bin/bash

# Function to check if Java is installed and version >= 11
check_java_version() {
    if ! command -v java &> /dev/null; then
        echo "Error: Java is not installed or not in the PATH."
        exit 1
    else
        java_version=$(java -version 2>&1 | awk -F '"' '/version/ {print $2}')
        if [[ "$java_version" > "11" ]]; then
            echo "Java version is greater than 11: $java_version"
        else
            echo "Error: Java version is not greater than 11. Installed Java version: $java_version"
            exit 1
        fi
    fi
}

# Function to download the JAR file if not present
download_jar_file() {
    script_dir="license-header"
    file_name="org.eclipse.dash.licenses-1.1.0.jar"
    download_url="https://repo.eclipse.org/content/repositories/dash-licenses/org/eclipse/dash/org.eclipse.dash.licenses/1.1.0/org.eclipse.dash.licenses-1.1.0.jar"

    if [ ! -e "$script_dir/$file_name" ]; then
        echo "$file_name is not found in $script_dir. Downloading it now..."
        cd "$script_dir"
        wget "$download_url"
        if [ $? -eq 0 ]; then
            echo "$file_name downloaded successfully."
        else
            echo "Error downloading $file_name."
        fi
    else
        echo "$file_name is found in $script_dir."
    fi
}

output_file="DEPENDENCIES"
input_file="./go.sum"

output_file_web="DEPENDENCIES_WEB"
input_file_web="./web/package-lock.json"

echo -e "\e[1mSTEP 1: Check for java version\e[0m" && check_java_version
echo -e "\n\e[1mSTEP 2: Check for jar file\e[0m" && download_jar_file
echo -e "\n\e[1mSTEP 3: Execute the command for root directory (output in $output_file file)\e[0m" && java -jar "./$script_dir/$file_name" -summary ./$output_file ./$input_file
echo -e "\n\e[1mSTEP 3: Execute the command for web directory (output in $output_file_web file)\e[0m" && java -jar "./$script_dir/$file_name" -summary ./$output_file_web ./$input_file_web
