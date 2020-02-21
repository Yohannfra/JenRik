#!/bin/bash

if (( $EUID != 0 )); then
    echo "Please run as root"
    exit
fi

sudo cp ./JenRik /usr/bin/jenerik
echo "Done !"
