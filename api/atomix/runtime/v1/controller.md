# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/v1/controller.proto](#atomix_runtime_v1_controller-proto)
    - [Agent](#atomix-runtime-v1-Agent)
    - [ConfigureAgentRequest](#atomix-runtime-v1-ConfigureAgentRequest)
    - [ConfigureAgentResponse](#atomix-runtime-v1-ConfigureAgentResponse)
    - [CreateProxyRequest](#atomix-runtime-v1-CreateProxyRequest)
    - [CreateProxyResponse](#atomix-runtime-v1-CreateProxyResponse)
    - [DestroyProxyRequest](#atomix-runtime-v1-DestroyProxyRequest)
    - [DestroyProxyResponse](#atomix-runtime-v1-DestroyProxyResponse)
    - [Driver](#atomix-runtime-v1-Driver)
    - [Proxy](#atomix-runtime-v1-Proxy)
    - [StartAgentRequest](#atomix-runtime-v1-StartAgentRequest)
    - [StartAgentResponse](#atomix-runtime-v1-StartAgentResponse)
    - [StopAgentRequest](#atomix-runtime-v1-StopAgentRequest)
    - [StopAgentResponse](#atomix-runtime-v1-StopAgentResponse)
  
    - [Controller](#atomix-runtime-v1-Controller)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_v1_controller-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/v1/controller.proto



<a name="atomix-runtime-v1-Agent"></a>

### Agent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |






<a name="atomix-runtime-v1-ConfigureAgentRequest"></a>

### ConfigureAgentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| driver | [Driver](#atomix-runtime-v1-Driver) |  |  |
| agent | [Agent](#atomix-runtime-v1-Agent) |  |  |






<a name="atomix-runtime-v1-ConfigureAgentResponse"></a>

### ConfigureAgentResponse







<a name="atomix-runtime-v1-CreateProxyRequest"></a>

### CreateProxyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proxy | [Proxy](#atomix-runtime-v1-Proxy) |  |  |






<a name="atomix-runtime-v1-CreateProxyResponse"></a>

### CreateProxyResponse







<a name="atomix-runtime-v1-DestroyProxyRequest"></a>

### DestroyProxyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [Proxy](#atomix-runtime-v1-Proxy) |  |  |






<a name="atomix-runtime-v1-DestroyProxyResponse"></a>

### DestroyProxyResponse







<a name="atomix-runtime-v1-Driver"></a>

### Driver



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="atomix-runtime-v1-Proxy"></a>

### Proxy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |






<a name="atomix-runtime-v1-StartAgentRequest"></a>

### StartAgentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| driver | [Driver](#atomix-runtime-v1-Driver) |  |  |
| agent | [Agent](#atomix-runtime-v1-Agent) |  |  |






<a name="atomix-runtime-v1-StartAgentResponse"></a>

### StartAgentResponse







<a name="atomix-runtime-v1-StopAgentRequest"></a>

### StopAgentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| driver | [Driver](#atomix-runtime-v1-Driver) |  |  |
| agent | [Agent](#atomix-runtime-v1-Agent) |  |  |






<a name="atomix-runtime-v1-StopAgentResponse"></a>

### StopAgentResponse






 

 

 


<a name="atomix-runtime-v1-Controller"></a>

### Controller
The runtime controller service provides functions for the core controller to manage
the runtime drivers based on external configurations.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StartAgent | [StartAgentRequest](#atomix-runtime-v1-StartAgentRequest) | [StartAgentResponse](#atomix-runtime-v1-StartAgentResponse) |  |
| StopAgent | [StopAgentRequest](#atomix-runtime-v1-StopAgentRequest) | [StopAgentResponse](#atomix-runtime-v1-StopAgentResponse) |  |
| CreateProxy | [CreateProxyRequest](#atomix-runtime-v1-CreateProxyRequest) | [CreateProxyResponse](#atomix-runtime-v1-CreateProxyResponse) |  |
| DestroyProxy | [DestroyProxyRequest](#atomix-runtime-v1-DestroyProxyRequest) | [DestroyProxyResponse](#atomix-runtime-v1-DestroyProxyResponse) |  |

 



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

