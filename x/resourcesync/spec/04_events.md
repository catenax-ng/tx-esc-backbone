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


## Keeper events

In addition to handlers events, the resourcesync keeper will produce events when the following methods are called

### none-yet