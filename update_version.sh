#!/bin/bash

# There are two places where we need to change the version number
# - in the README
# - in the jenrik source code

# First we get the current version from the readme
current_version=$(grep version README.md -A 2 | grep "^v" | cut -d ' ' '-f2-')

# The new_version is the current one + 0.1
new_version=$(echo ${current_version} + 0.1 | bc)

# Now we save it

# In README
sed -i "s/${current_version}/${new_version}/" README.md

# In jenrik
sed -i "s/${current_version}/${new_version}/" jenrik
