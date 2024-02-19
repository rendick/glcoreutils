#!/bin/bash

cd ..

PS3='Please select an option: '
types=("Configuration (!)" "Building")

BUILD_DIR=$(pwd)

BOLD="\e[1m"
CLEAR="\e[0m"

configuration () {
    if [[ $1 == "Configuration (!)" ]] ; then
			echo "Soon..."
    elif [[ $1 == "Building" ]] ; then
		echo "Soon..."
    else
        echo 'Invalid installation type!'
    fi
}

echo

select type in "${types[@]}" ; do
    if [[ -n $type ]] ; then
        break
    else
        echo 'Please select a valid option!'
    fi
done

configuration "$type"
