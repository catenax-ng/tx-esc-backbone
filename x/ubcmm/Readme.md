<!--
 Copyright (c) 2022-2023 Contributors to the Eclipse Foundation

 See the NOTICE file(s) distributed with this work for additional
 information regarding copyright ownership.

 This program and the accompanying materials are made available under the
 terms of the Apache License, Version 2.0 which is available at
 https://www.apache.org/licenses/LICENSE-2.0.

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 License for the specific language governing permissions and limitations
 under the License.

 SPDX-License-Identifier: Apache-2.0
-->
# `x/ubcmm`

* [Technical aspects](technical aspects)

## Technical aspects

### Handling Precision

#### Library used

We use the
[Dec](https://github.com/cosmos/cosmos-sdk/blob/8f6a94cd1f9f1c6bf1ad83a751da86270db92e02/types/math.go#L40)
type defined in types package of cosmos-sdk for all our numerical computations.
This implements arbitrary precision fixed point arithmetics.

The library 

- Allows maximum of 18 decimal places.
- Allows maximum of 315 bits to store decimal values (256 for the whole number part and 59 for the fractional part).
- Does all the computations with a precision of 18 decimal places.

See [here](https://github.com/cosmos/cosmos-sdk/blob/main/math/dec.go#L15) for more details on this type.

#### Considerations in ubcmm module

In ubcmm module, we have three difference precision to be considered:

##### 1. Precision for system token (value is 6)

This is the accuracy with which we specify the fractional part of token value.
This is same as the smallest of base denomination of the token. Currently the
base denomination is "ucax" and hence this value is 6.

All token values such as amount of tokens to buy/sell/undergird/upshift and
BPool should use this precision.

##### 2. Precision for voucher (value is 2)

This is the accuracy with which we specify the fractional part of voucher
value. This is same as the smallest of base denomination of the voucher.
Currently the base denomination is "cvoucher" and hence this value is 2.

All voucher values such as amount of vouchers consumed during a buy, earned
during a sell, held by the ubcmm module should use this precision.

##### 3. Precision for curve computation (value is 18)

This is the precision used for all intermediate calculations. For this, we use
the maximum precision offered by the library, which is 18.

To ensure the correctness of numerical operations, 

1. For Init, Undergird and Shift-up operations, we validate the curve conditions after the
   operation, to check if there has been an errors due to precision in computation.

2. In case of buy, sell, the computations and rounding off scheme are designed to avoid
   precision errors. Tests for these function check if the intended design works.

3. For checking if BPool value equals integral over the curve, we use the precision for system token.


Note: In addition to the three precision numbers described above, there will
additional values introduced when implementing the orderbook feature.

