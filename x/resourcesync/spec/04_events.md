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