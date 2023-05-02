#!/bin/bash
# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0

# Start the chain using ignite chain serve command.

# config chain id
esc-backboned config chain-id escbackbone

# balances before init
esc-backboned query bank balances catenax13er304pz9kz6dd8zs2e9uvqlmr5jtw67rvslmp # module account
esc-backboned query bank balances catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw # ubc_initiator account


# ubc init - errors
esc-backboned tx ubc init "a" "1" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes
esc-backboned tx ubc init "6000000000" "a" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes
esc-backboned tx ubc init "6000000000" "1" "a" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes
esc-backboned tx ubc init "6000000000" "1" "10" "a" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes
esc-backboned tx ubc init "6000000000" "1" "10" "100000000" "a" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes
esc-backboned tx ubc init "6000000000" "1" "10" "100000000" "100000000" "a" "0.000000000666666667" "0.2" "15832600001" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes
esc-backboned tx ubc init "6000000000" "1" "10" "100000000" "100000000" "0.0000000001" "a" "0.2" "15832600001" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes
esc-backboned tx ubc init "6000000000" "1" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "a" "15832600001" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes
esc-backboned tx ubc init "6000000000" "1" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "a" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes

# ubc init - happy
esc-backboned tx ubc init "6000000000" "1" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw --yes

sleep 3

# ubc query object
esc-backboned query ubc show-ubcobject

# balances after init
esc-backboned query bank balances catenax13er304pz9kz6dd8zs2e9uvqlmr5jtw67rvslmp # module account
esc-backboned query bank balances catenax1tu9hgj7pj7cupuv7e5qn0xqpvdapl5kvs4g6tw # ubc_initiator account

