<!--
order: 2
-->

# State

The `x/resourcesync` module keeps state of a list of resources.

- ResourceMap: ` byte('ResourceMap/value/') -> byte(ResourceMap.originator) | byte('/') | byte(ResourceMap.origResId) | byte('/')-> ProtocolBuffer(ResourceMap)`
