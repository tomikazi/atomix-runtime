<!--
SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/controller/v1/controller.proto](#atomix_controller_v1_controller-proto)
    - [Configuration](#atomix-controller-v1-Configuration)
    - [GetConfigurationRequest](#atomix-controller-v1-GetConfigurationRequest)
    - [GetConfigurationResponse](#atomix-controller-v1-GetConfigurationResponse)
    - [WatchConfigurationRequest](#atomix-controller-v1-WatchConfigurationRequest)
    - [WatchConfigurationResponse](#atomix-controller-v1-WatchConfigurationResponse)
  
    - [Controller](#atomix-controller-v1-Controller)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_controller_v1_controller-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/controller/v1/controller.proto



<a name="atomix-controller-v1-Configuration"></a>

### Configuration



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| store | [string](#string) |  |  |
| driver | [string](#string) |  |  |
| version | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="atomix-controller-v1-GetConfigurationRequest"></a>

### GetConfigurationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| store | [string](#string) |  |  |






<a name="atomix-controller-v1-GetConfigurationResponse"></a>

### GetConfigurationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration | [Configuration](#atomix-controller-v1-Configuration) |  |  |






<a name="atomix-controller-v1-WatchConfigurationRequest"></a>

### WatchConfigurationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| store | [string](#string) |  |  |






<a name="atomix-controller-v1-WatchConfigurationResponse"></a>

### WatchConfigurationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration | [Configuration](#atomix-controller-v1-Configuration) |  |  |





 

 

 


<a name="atomix-controller-v1-Controller"></a>

### Controller
The controller service provides control functions for the runtime.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetConfiguration | [GetConfigurationRequest](#atomix-controller-v1-GetConfigurationRequest) | [GetConfigurationResponse](#atomix-controller-v1-GetConfigurationResponse) |  |
| WatchConfiguration | [WatchConfigurationRequest](#atomix-controller-v1-WatchConfigurationRequest) | [WatchConfigurationResponse](#atomix-controller-v1-WatchConfigurationResponse) stream |  |

 



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

