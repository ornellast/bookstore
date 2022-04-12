#!/usr/bin/env bash

source .env


while read -r line
do
if [[ ${#line} -gt 3 ]]; then
  vr=$(echo ${line} | cut -d= -f 1)
  vl=$(echo ${line} | cut -d= -f 2)
  export $vr="${vl}"
fi
  # ct=ct+1
done < '.env'

go run .