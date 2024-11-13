#!/bin/bash

# Fetch the JSON data and filter the name for subject ID 70
name=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '.[] | select(.id == 70) | .name')

echo "\"$name\""

