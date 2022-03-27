# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/v1/runtime.proto](#atomix_runtime_v1_runtime-proto)
    - [CloseProxyRequest](#atomix-runtime-v1-CloseProxyRequest)
    - [CloseProxyResponse](#atomix-runtime-v1-CloseProxyResponse)
    - [CreatePrimitiveRequest](#atomix-runtime-v1-CreatePrimitiveRequest)
    - [CreatePrimitiveResponse](#atomix-runtime-v1-CreatePrimitiveResponse)
    - [CreateProxyRequest](#atomix-runtime-v1-CreateProxyRequest)
    - [CreateProxyResponse](#atomix-runtime-v1-CreateProxyResponse)
    - [DeletePrimitiveRequest](#atomix-runtime-v1-DeletePrimitiveRequest)
    - [DeletePrimitiveResponse](#atomix-runtime-v1-DeletePrimitiveResponse)
    - [Primitive](#atomix-runtime-v1-Primitive)
    - [Proxy](#atomix-runtime-v1-Proxy)
  
    - [Runtime](#atomix-runtime-v1-Runtime)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_v1_runtime-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/v1/runtime.proto



<a name="atomix-runtime-v1-CloseProxyRequest"></a>

### CloseProxyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proxy_id | [uint64](#uint64) |  |  |






<a name="atomix-runtime-v1-CloseProxyResponse"></a>

### CloseProxyResponse







<a name="atomix-runtime-v1-CreatePrimitiveRequest"></a>

### CreatePrimitiveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive | [Primitive](#atomix-runtime-v1-Primitive) |  |  |






<a name="atomix-runtime-v1-CreatePrimitiveResponse"></a>

### CreatePrimitiveResponse







<a name="atomix-runtime-v1-CreateProxyRequest"></a>

### CreateProxyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proxy | [Proxy](#atomix-runtime-v1-Proxy) |  |  |






<a name="atomix-runtime-v1-CreateProxyResponse"></a>

### CreateProxyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proxy_id | [uint64](#uint64) |  |  |






<a name="atomix-runtime-v1-DeletePrimitiveRequest"></a>

### DeletePrimitiveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive_id | [string](#string) |  |  |
| store_id | [string](#string) |  |  |






<a name="atomix-runtime-v1-DeletePrimitiveResponse"></a>

### DeletePrimitiveResponse







<a name="atomix-runtime-v1-Primitive"></a>

### Primitive



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| version | [string](#string) |  |  |
| options | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="atomix-runtime-v1-Proxy"></a>

### Proxy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive_id | [string](#string) |  |  |
| options | [google.protobuf.Any](#google-protobuf-Any) |  |  |





 

 

 


<a name="atomix-runtime-v1-Runtime"></a>

### Runtime
The runtime service provides functions for applications to create and manage primitives
at runtime.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreatePrimitive | [CreatePrimitiveRequest](#atomix-runtime-v1-CreatePrimitiveRequest) | [CreatePrimitiveResponse](#atomix-runtime-v1-CreatePrimitiveResponse) |  |
| DeletePrimitive | [DeletePrimitiveRequest](#atomix-runtime-v1-DeletePrimitiveRequest) | [DeletePrimitiveResponse](#atomix-runtime-v1-DeletePrimitiveResponse) |  |
| CreateProxy | [CreateProxyRequest](#atomix-runtime-v1-CreateProxyRequest) | [CreateProxyResponse](#atomix-runtime-v1-CreateProxyResponse) |  |
| CloseProxy | [CloseProxyRequest](#atomix-runtime-v1-CloseProxyRequest) | [CloseProxyResponse](#atomix-runtime-v1-CloseProxyResponse) |  |

 



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

