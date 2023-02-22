#!/bin/bash
# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone
#
# SPDX-License-Identifier: Apache-2.0

SCRIPT_LOCATION=$(realpath $(dirname -- "${BASH_SOURCE[0]}" ))

cd $SCRIPT_LOCATION/..
npx license-check-and-add@4.0.5 remove -f $SCRIPT_LOCATION/config.json && \
npx license-check-and-add@4.0.5 add -f $SCRIPT_LOCATION/config.json -r "-`date +%Y`"