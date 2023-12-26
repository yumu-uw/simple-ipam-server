#!/bin/zsh
oapi-codegen --config=cfg/server.cfg.yml ../simple-ipam-opai/api.yml
oapi-codegen --config=cfg/types.cfg.yml ../simple-ipam-opai/api.yml