<!--
order: 3
-->

# Messages

## MsgCreateResource

Creates a resource entry to publish its creation via an event.

The message will fail under the following conditions:

- There is already a resource with the same originator origResId combination.