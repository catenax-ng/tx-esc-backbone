#!/bin/bash
# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0

# Start the chain using ignite chain serve command.


ubc_module="catenax13er304pz9kz6dd8zs2e9uvqlmr5jtw67rvslmp"
ubc_initiator=$(esc-backboned keys show -a ubc_initiator)
ubc_trader=$(esc-backboned keys show -a ubc_trader)
ubc_operator=$(esc-backboned keys show -a ubc_operator)

echo "\n config chain id\n\n"
esc-backboned config chain-id escbackbone

echo "\n balances before init\n\n"
esc-backboned query bank balances $ubc_module
esc-backboned query bank balances $ubc_initiator # ubc_initiator account


echo "\n ubc init - errors\n\n"
esc-backboned tx ubcmm init "a" "1" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from $ubc_initiator --yes
esc-backboned tx ubcmm init "6000000000" "a" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from $ubc_initiator --yes
esc-backboned tx ubcmm init "6000000000" "1" "a" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from $ubc_initiator --yes
esc-backboned tx ubcmm init "6000000000" "1" "10" "a" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from $ubc_initiator --yes
esc-backboned tx ubcmm init "6000000000" "1" "10" "100000000" "a" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from $ubc_initiator --yes
esc-backboned tx ubcmm init "6000000000" "1" "10" "100000000" "100000000" "a" "0.000000000666666667" "0.2" "15832600001" --from $ubc_initiator --yes
esc-backboned tx ubcmm init "6000000000" "1" "10" "100000000" "100000000" "0.0000000001" "a" "0.2" "15832600001" --from $ubc_initiator --yes
esc-backboned tx ubcmm init "6000000000" "1" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "a" "15832600001" --from $ubc_initiator --yes
esc-backboned tx ubcmm init "6000000000" "1" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "a" --from $ubc_initiator --yes

echo "\n ubc init - happy\n\n"
esc-backboned tx ubcmm init "6000000000" "1" "10" "100000000" "100000000" "0.0000000001" "0.000000000666666667" "0.2" "15832600001" --from $ubc_initiator --yes

sleep 3

echo "\n ubc query curve\n\n"
esc-backboned query ubcmm show-curve

echo "\n balances after init\n\n"
esc-backboned query bank balances $ubc_module
esc-backboned query bank balances $ubc_initiator


echo "\n# ubc buy - value in tokens\n\n"
esc-backboned query bank balances $ubc_trader
esc-backboned tx ubcmm buy "10000ucax" --from $ubc_trader --yes
sleep 3
esc-backboned query bank balances $ubc_trader


echo "\n# ubc sell - value in tokens\n\n"
esc-backboned query bank balances $ubc_trader
esc-backboned tx ubcmm sell "10000ucax" --from=$ubc_trader --yes
sleep 3
esc-backboned query bank balances $ubc_trader

echo "\n# ubc send tokens from one address to another \n\n"
esc-backboned query bank balances $ubc_trader --denom=ucax
esc-backboned query bank balances $ubc_initiator --denom=ucax
esc-backboned tx bank send $ubc_trader $ubc_initiator 10ucax --yes
sleep 3
esc-backboned query bank balances $ubc_trader --denom=ucax
esc-backboned query bank balances $ubc_initiator --denom=ucax

echo "\n# ubc send vouchers from one address to another (should fail) \n\n"
esc-backboned query bank balances $ubc_trader --denom=uvoucher
esc-backboned query bank balances $ubc_initiator --denom=uvoucher
esc-backboned tx bank send $ubc_trader $ubc_initiator 10uvoucher --yes
sleep 3
esc-backboned query bank balances $ubc_trader --denom=uvoucher
esc-backboned query bank balances $ubc_initiator --denom=uvoucher

echo "\n# ubc undergird \n\n"
esc-backboned query bank balances $ubc_operator
esc-backboned tx ubcmm undergird "10000000000000uvoucher" --from=$ubc_operator --yes
sleep 3
esc-backboned query bank balances $ubc_operator
esc-backboned query ubcmm show-curve


echo "\n# ubc shiftup - value in tokens\n\n"
esc-backboned query bank balances $ubc_operator
echo "\n# buy large amount of tokens (500e5 CAX) to shift current supply significanly beyond P2.\n\n"
esc-backboned tx ubcmm buy "50000000000000ucax" --from=$ubc_operator --yes
sleep 3
esc-backboned query bank balances $ubc_operator
esc-backboned tx ubcmm shiftup "10000000000000uvoucher" "1" --from=$ubc_operator --yes
sleep 3
esc-backboned query bank balances $ubc_operator
esc-backboned query ubcmm show-curve
