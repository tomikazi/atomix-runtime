<!--
SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/value/v1/primitive.proto](#atomix_runtime_value_v1_primitive-proto)
    - [Event](#atomix-runtime-value-v1-Event)
    - [EventsRequest](#atomix-runtime-value-v1-EventsRequest)
    - [EventsResponse](#atomix-runtime-value-v1-EventsResponse)
    - [GetRequest](#atomix-runtime-value-v1-GetRequest)
    - [GetResponse](#atomix-runtime-value-v1-GetResponse)
    - [Object](#atomix-runtime-value-v1-Object)
    - [Precondition](#atomix-runtime-value-v1-Precondition)
    - [SetRequest](#atomix-runtime-value-v1-SetRequest)
    - [SetResponse](#atomix-runtime-value-v1-SetResponse)
  
    - [Event.Type](#atomix-runtime-value-v1-Event-Type)
  
    - [Value](#atomix-runtime-value-v1-Value)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_value_v1_primitive-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/value/v1/primitive.proto



<a name="atomix-runtime-value-v1-Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Event.Type](#atomix-runtime-value-v1-Event-Type) |  |  |
| value | [Object](#atomix-runtime-value-v1-Object) |  |  |






<a name="atomix-runtime-value-v1-EventsRequest"></a>

### EventsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |






<a name="atomix-runtime-value-v1-EventsResponse"></a>

### EventsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| event | [Event](#atomix-runtime-value-v1-Event) |  |  |






<a name="atomix-runtime-value-v1-GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |






<a name="atomix-runtime-value-v1-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| value | [Object](#atomix-runtime-value-v1-Object) |  |  |






<a name="atomix-runtime-value-v1-Object"></a>

### Object



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [atomix.runtime.meta.v1.ObjectMeta](#atomix-runtime-meta-v1-ObjectMeta) |  |  |
| value | [bytes](#bytes) |  |  |






<a name="atomix-runtime-value-v1-Precondition"></a>

### Precondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [atomix.runtime.meta.v1.ObjectMeta](#atomix-runtime-meta-v1-ObjectMeta) |  |  |






<a name="atomix-runtime-value-v1-SetRequest"></a>

### SetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |
| value | [Object](#atomix-runtime-value-v1-Object) |  |  |
| preconditions | [Precondition](#atomix-runtime-value-v1-Precondition) | repeated |  |






<a name="atomix-runtime-value-v1-SetResponse"></a>

### SetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| value | [Object](#atomix-runtime-value-v1-Object) |  |  |





 


<a name="atomix-runtime-value-v1-Event-Type"></a>

### Event.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| UPDATE | 1 |  |


 

 


<a name="atomix-runtime-value-v1-Value"></a>

### Value
Value is a service for a value primitive

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Set | [SetRequest](#atomix-runtime-value-v1-SetRequest) | [SetResponse](#atomix-runtime-value-v1-SetResponse) | Set sets the value |
| Get | [GetRequest](#atomix-runtime-value-v1-GetRequest) | [GetResponse](#atomix-runtime-value-v1-GetResponse) | Get gets the value |
| Events | [EventsRequest](#atomix-runtime-value-v1-EventsRequest) | [EventsResponse](#atomix-runtime-value-v1-EventsResponse) stream | Events listens for value change events |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

