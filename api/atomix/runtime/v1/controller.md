# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/v1/controller.proto](#atomix_runtime_v1_controller-proto)
    - [ConfigureDriverRequest](#atomix-runtime-v1-ConfigureDriverRequest)
    - [ConfigureDriverResponse](#atomix-runtime-v1-ConfigureDriverResponse)
    - [Driver](#atomix-runtime-v1-Driver)
    - [DriverId](#atomix-runtime-v1-DriverId)
    - [StartDriverRequest](#atomix-runtime-v1-StartDriverRequest)
    - [StartDriverResponse](#atomix-runtime-v1-StartDriverResponse)
    - [StopDriverRequest](#atomix-runtime-v1-StopDriverRequest)
    - [StopDriverResponse](#atomix-runtime-v1-StopDriverResponse)
  
    - [Controller](#atomix-runtime-v1-Controller)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_v1_controller-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/v1/controller.proto



<a name="atomix-runtime-v1-ConfigureDriverRequest"></a>

### ConfigureDriverRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| driver | [Driver](#atomix-runtime-v1-Driver) |  |  |






<a name="atomix-runtime-v1-ConfigureDriverResponse"></a>

### ConfigureDriverResponse







<a name="atomix-runtime-v1-Driver"></a>

### Driver



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [DriverId](#atomix-runtime-v1-DriverId) |  |  |






<a name="atomix-runtime-v1-DriverId"></a>

### DriverId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |






<a name="atomix-runtime-v1-StartDriverRequest"></a>

### StartDriverRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| driver | [Driver](#atomix-runtime-v1-Driver) |  |  |






<a name="atomix-runtime-v1-StartDriverResponse"></a>

### StartDriverResponse







<a name="atomix-runtime-v1-StopDriverRequest"></a>

### StopDriverRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [DriverId](#atomix-runtime-v1-DriverId) |  |  |






<a name="atomix-runtime-v1-StopDriverResponse"></a>

### StopDriverResponse






 

 

 


<a name="atomix-runtime-v1-Controller"></a>

### Controller


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StartDriver | [StartDriverRequest](#atomix-runtime-v1-StartDriverRequest) | [StartDriverResponse](#atomix-runtime-v1-StartDriverResponse) |  |
| ConfigureDriver | [ConfigureDriverRequest](#atomix-runtime-v1-ConfigureDriverRequest) | [ConfigureDriverResponse](#atomix-runtime-v1-ConfigureDriverResponse) |  |
| StopDriver | [StopDriverRequest](#atomix-runtime-v1-StopDriverRequest) | [StopDriverResponse](#atomix-runtime-v1-StopDriverResponse) |  |

 



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

