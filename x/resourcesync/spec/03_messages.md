<!--
 Copyright (c) 2022-2023 - for information on the respective copyright owner
 see the NOTICE file and/or the repository at
 https://github.com/catenax-ng/product-esc-backbone-code

 SPDX-License-Identifier: Apache-2.0
-->
<!--
order: 3
-->

# Messages

## MsgCreateResource

Creates a resource entry to publish its creation via an event.

The message will fail under the following conditions:

- There is already a resource with the same `originator` `origResId` combination.

## MsgDeleteResource

Deletes a resource defined by the `originator` and the `origResId` entry to publish its deletion via an event.

The message will fail under the following conditions:

- There is no resource with the same `originator` `origResId` combination.

## MsgUpdateResource

Changes a resource entry to publish its update via an event.

The message will fail under the following conditions:

- There is no resource with the same `originator` `origResId` combination.