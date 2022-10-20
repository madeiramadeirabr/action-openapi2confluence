#!/bin/sh -l

openapi2confluence -p $1 -id $2 -s $3 -a $4 -t $5 -lid $6 -mid $7 -env $8 >> $GITHUB_OUTPUT