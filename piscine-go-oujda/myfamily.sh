#!/bin/bash
curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r ' .[] | select(.id == '$HERO_ID') | .connections.relatives | gsub("\\n";"\\n")'
