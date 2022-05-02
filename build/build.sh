#!/bin/bash


DOCKER_BUILDKIT=1 docker build --progress=plain --build-arg  BRANCH=main --secret id=ssh_key,src=./ssh_key --target  node -t esc-backbone .
