#!/bin/sh -l 

echo $1
echo $2
echo $3
echo $4
echo $5
echo $6
echo $7
echo $8

openapi2confluence -p $1 -id $2 -s $3 -a $4 -t $5 -lid $6 -mid $7 -env $8 >> $GITHUB_STEP_SUMMARY