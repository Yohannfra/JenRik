#!/bin/bash

if (( $EUID != 0 )); then
    echo "Please run as root"
    exit
fi

sudo cp ./jenrik /usr/bin/jenerik
echo "Done !"
