#!/bin/bash

# Fetch the JSON data and filter the name, power, and gender for subject ID 170
curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | [.name, .powerstats.power, .appearance.gender] | @tsv'

