#!/bin/bash

echo "enter jce dir"
cd ../testjce

echo "begin pull github"
git pull

echo "begin build jce"
jce2go -o /tmp/api -mod github.com/erpc-go/jce2go  ./*

echo "enter tesjce2go/"
cd /home/lighthouse/testjce2go

echo "cp to api project"
cp -r /tmp/api/* .

echo "begin go mod tidy"
go mod tidy

echo "begin push api project"
git pull
git status
git add *
git add ./
git commit -m "feat(*): update"
git push

echo "update succ"
