<!--
SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/counter/v1/primitive.proto](#atomix_runtime_counter_v1_primitive-proto)
    - [DecrementRequest](#atomix-runtime-counter-v1-DecrementRequest)
    - [DecrementResponse](#atomix-runtime-counter-v1-DecrementResponse)
    - [GetRequest](#atomix-runtime-counter-v1-GetRequest)
    - [GetResponse](#atomix-runtime-counter-v1-GetResponse)
    - [IncrementRequest](#atomix-runtime-counter-v1-IncrementRequest)
    - [IncrementResponse](#atomix-runtime-counter-v1-IncrementResponse)
    - [Precondition](#atomix-runtime-counter-v1-Precondition)
    - [SetRequest](#atomix-runtime-counter-v1-SetRequest)
    - [SetResponse](#atomix-runtime-counter-v1-SetResponse)
    - [Value](#atomix-runtime-counter-v1-Value)
  
    - [Counter](#atomix-runtime-counter-v1-Counter)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_counter_v1_primitive-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/counter/v1/primitive.proto



<a name="atomix-runtime-counter-v1-DecrementRequest"></a>

### DecrementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |
| delta | [int64](#int64) |  |  |






<a name="atomix-runtime-counter-v1-DecrementResponse"></a>

### DecrementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| value | [int64](#int64) |  |  |






<a name="atomix-runtime-counter-v1-GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |






<a name="atomix-runtime-counter-v1-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| value | [int64](#int64) |  |  |






<a name="atomix-runtime-counter-v1-IncrementRequest"></a>

### IncrementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |
| delta | [int64](#int64) |  |  |






<a name="atomix-runtime-counter-v1-IncrementResponse"></a>

### IncrementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| value | [int64](#int64) |  |  |






<a name="atomix-runtime-counter-v1-Precondition"></a>

### Precondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |






<a name="atomix-runtime-counter-v1-SetRequest"></a>

### SetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |
| value | [int64](#int64) |  |  |
| preconditions | [Precondition](#atomix-runtime-counter-v1-Precondition) | repeated |  |






<a name="atomix-runtime-counter-v1-SetResponse"></a>

### SetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| value | [int64](#int64) |  |  |






<a name="atomix-runtime-counter-v1-Value"></a>

### Value



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |





 

 

 


<a name="atomix-runtime-counter-v1-Counter"></a>

### Counter
Counter is a service for a counter primitive

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Set | [SetRequest](#atomix-runtime-counter-v1-SetRequest) | [SetResponse](#atomix-runtime-counter-v1-SetResponse) | Set sets the counter value |
| Get | [GetRequest](#atomix-runtime-counter-v1-GetRequest) | [GetResponse](#atomix-runtime-counter-v1-GetResponse) | Get gets the current counter value |
| Increment | [IncrementRequest](#atomix-runtime-counter-v1-IncrementRequest) | [IncrementResponse](#atomix-runtime-counter-v1-IncrementResponse) | Increment increments the counter value |
| Decrement | [DecrementRequest](#atomix-runtime-counter-v1-DecrementRequest) | [DecrementResponse](#atomix-runtime-counter-v1-DecrementResponse) | Decrement decrements the counter value |

 



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

