#!/bin/sh

git config --global --add safe.directory /workspaces

go mod download && go mod verify
