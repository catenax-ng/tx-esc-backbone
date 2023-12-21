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