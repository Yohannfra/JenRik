#!/bin/bash

if (( $EUID != 0 )); then
    echo "Please run as root"
    exit 1
fi

sudo cp ./jenrik /usr/bin/jenrik
echo "Done !"
