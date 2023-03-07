<!--
 Copyright (c) 2022-2023 - for information on the respective copyright owner
 see the NOTICE file and/or the repository at
 https://github.com/catenax-ng/product-esc-backbone-code

 SPDX-License-Identifier: Apache-2.0
-->
<!--
order: 4
-->

# Events

The resourcesync module emits the following events:

## Handlers

### MsgCreateResource

| Type                                                 | Attribute Key | Attribute Value          |
|------------------------------------------------------|---------------|--------------------------|
| catenax.escbackbone.resourcesync.EventCreateResource | creator       | {senderAddress}          |
| catenax.escbackbone.resourcesync.EventCreateResource | resource      | {created types.Resource} |

### MsgDeleteResource

| Type                                                 | Attribute Key | Attribute Value          |
|------------------------------------------------------|---------------|--------------------------|
| catenax.escbackbone.resourcesync.EventDeleteResource | creator       | {senderAddress}          |
| catenax.escbackbone.resourcesync.EventDeleteResource | resource      | {deleted types.Resource} |

### MsgUpdateResource

| Type                                                 | Attribute Key | Attribute Value          |
|------------------------------------------------------|---------------|--------------------------|
| catenax.escbackbone.resourcesync.EventUpdateResource | creator       | {senderAddress}          |
| catenax.escbackbone.resourcesync.EventUpdateResource | resource      | {updated types.Resource} |

## Keeper events

In addition to handlers events, the resourcesync keeper will produce events when the following methods are called

### none-yet