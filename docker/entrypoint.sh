#!/bin/sh -l -xe

openapi2confluence -p $1 -id $2 -s $3 -a $4 -t $5 -lid $6 -mid $7 -env $8 >> $GITHUB_STEP_SUMMARY