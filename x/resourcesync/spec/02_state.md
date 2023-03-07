<!--
 Copyright (c) 2022-2023 - for information on the respective copyright owner
 see the NOTICE file and/or the repository at
 https://github.com/catenax-ng/product-esc-backbone-code

 SPDX-License-Identifier: Apache-2.0
-->
<!--
order: 2
-->

# State

The `x/resourcesync` module keeps state of a list of resources.

- ResourceMap: ` byte('ResourceMap/value/') -> byte(ResourceMap.originator) | byte('/') | byte(ResourceMap.origResId) | byte('/')-> ProtocolBuffer(ResourceMap)`
