# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/primitive/list/v1/service.proto](#atomix_primitive_list_v1_service-proto)
    - [AppendRequest](#atomix-primitive-list-v1-AppendRequest)
    - [AppendResponse](#atomix-primitive-list-v1-AppendResponse)
    - [ClearRequest](#atomix-primitive-list-v1-ClearRequest)
    - [ClearResponse](#atomix-primitive-list-v1-ClearResponse)
    - [ContainsRequest](#atomix-primitive-list-v1-ContainsRequest)
    - [ContainsResponse](#atomix-primitive-list-v1-ContainsResponse)
    - [ElementsRequest](#atomix-primitive-list-v1-ElementsRequest)
    - [ElementsResponse](#atomix-primitive-list-v1-ElementsResponse)
    - [Event](#atomix-primitive-list-v1-Event)
    - [EventsRequest](#atomix-primitive-list-v1-EventsRequest)
    - [EventsResponse](#atomix-primitive-list-v1-EventsResponse)
    - [GetRequest](#atomix-primitive-list-v1-GetRequest)
    - [GetResponse](#atomix-primitive-list-v1-GetResponse)
    - [InsertRequest](#atomix-primitive-list-v1-InsertRequest)
    - [InsertResponse](#atomix-primitive-list-v1-InsertResponse)
    - [Item](#atomix-primitive-list-v1-Item)
    - [Precondition](#atomix-primitive-list-v1-Precondition)
    - [RemoveRequest](#atomix-primitive-list-v1-RemoveRequest)
    - [RemoveResponse](#atomix-primitive-list-v1-RemoveResponse)
    - [SetRequest](#atomix-primitive-list-v1-SetRequest)
    - [SetResponse](#atomix-primitive-list-v1-SetResponse)
    - [SizeRequest](#atomix-primitive-list-v1-SizeRequest)
    - [SizeResponse](#atomix-primitive-list-v1-SizeResponse)
    - [Value](#atomix-primitive-list-v1-Value)
  
    - [Event.Type](#atomix-primitive-list-v1-Event-Type)
  
    - [List](#atomix-primitive-list-v1-List)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_primitive_list_v1_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/primitive/list/v1/service.proto



<a name="atomix-primitive-list-v1-AppendRequest"></a>

### AppendRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [Value](#atomix-primitive-list-v1-Value) |  |  |






<a name="atomix-primitive-list-v1-AppendResponse"></a>

### AppendResponse







<a name="atomix-primitive-list-v1-ClearRequest"></a>

### ClearRequest







<a name="atomix-primitive-list-v1-ClearResponse"></a>

### ClearResponse







<a name="atomix-primitive-list-v1-ContainsRequest"></a>

### ContainsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [Value](#atomix-primitive-list-v1-Value) |  |  |






<a name="atomix-primitive-list-v1-ContainsResponse"></a>

### ContainsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contains | [bool](#bool) |  |  |






<a name="atomix-primitive-list-v1-ElementsRequest"></a>

### ElementsRequest







<a name="atomix-primitive-list-v1-ElementsResponse"></a>

### ElementsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [Item](#atomix-primitive-list-v1-Item) |  |  |






<a name="atomix-primitive-list-v1-Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Event.Type](#atomix-primitive-list-v1-Event-Type) |  |  |
| item | [Item](#atomix-primitive-list-v1-Item) |  |  |






<a name="atomix-primitive-list-v1-EventsRequest"></a>

### EventsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| replay | [bool](#bool) |  |  |






<a name="atomix-primitive-list-v1-EventsResponse"></a>

### EventsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#atomix-primitive-list-v1-Event) |  |  |






<a name="atomix-primitive-list-v1-GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint32](#uint32) |  |  |






<a name="atomix-primitive-list-v1-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [Item](#atomix-primitive-list-v1-Item) |  |  |






<a name="atomix-primitive-list-v1-InsertRequest"></a>

### InsertRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [Item](#atomix-primitive-list-v1-Item) |  |  |
| preconditions | [Precondition](#atomix-primitive-list-v1-Precondition) | repeated |  |






<a name="atomix-primitive-list-v1-InsertResponse"></a>

### InsertResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [Item](#atomix-primitive-list-v1-Item) |  |  |






<a name="atomix-primitive-list-v1-Item"></a>

### Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint32](#uint32) |  |  |
| value | [Value](#atomix-primitive-list-v1-Value) |  |  |






<a name="atomix-primitive-list-v1-Precondition"></a>

### Precondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [atomix.primitive.meta.v1.ObjectMeta](#atomix-primitive-meta-v1-ObjectMeta) |  |  |






<a name="atomix-primitive-list-v1-RemoveRequest"></a>

### RemoveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint32](#uint32) |  |  |
| preconditions | [Precondition](#atomix-primitive-list-v1-Precondition) | repeated |  |






<a name="atomix-primitive-list-v1-RemoveResponse"></a>

### RemoveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [Item](#atomix-primitive-list-v1-Item) |  |  |






<a name="atomix-primitive-list-v1-SetRequest"></a>

### SetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [Item](#atomix-primitive-list-v1-Item) |  |  |
| preconditions | [Precondition](#atomix-primitive-list-v1-Precondition) | repeated |  |






<a name="atomix-primitive-list-v1-SetResponse"></a>

### SetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [Item](#atomix-primitive-list-v1-Item) |  |  |






<a name="atomix-primitive-list-v1-SizeRequest"></a>

### SizeRequest







<a name="atomix-primitive-list-v1-SizeResponse"></a>

### SizeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| size | [uint32](#uint32) |  |  |






<a name="atomix-primitive-list-v1-Value"></a>

### Value



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [atomix.primitive.meta.v1.ObjectMeta](#atomix-primitive-meta-v1-ObjectMeta) |  |  |
| value | [string](#string) |  |  |





 


<a name="atomix-primitive-list-v1-Event-Type"></a>

### Event.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ADD | 1 |  |
| REMOVE | 2 |  |
| REPLAY | 3 |  |


 

 


<a name="atomix-primitive-list-v1-List"></a>

### List
List is a service for a list primitive

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Size | [SizeRequest](#atomix-primitive-list-v1-SizeRequest) | [SizeResponse](#atomix-primitive-list-v1-SizeResponse) | Size gets the number of elements in the list |
| Append | [AppendRequest](#atomix-primitive-list-v1-AppendRequest) | [AppendResponse](#atomix-primitive-list-v1-AppendResponse) | Append appends a value to the list |
| Insert | [InsertRequest](#atomix-primitive-list-v1-InsertRequest) | [InsertResponse](#atomix-primitive-list-v1-InsertResponse) | Insert inserts a value at a specific index in the list |
| Get | [GetRequest](#atomix-primitive-list-v1-GetRequest) | [GetResponse](#atomix-primitive-list-v1-GetResponse) | Get gets the value at an index in the list |
| Set | [SetRequest](#atomix-primitive-list-v1-SetRequest) | [SetResponse](#atomix-primitive-list-v1-SetResponse) | Set sets the value at an index in the list |
| Remove | [RemoveRequest](#atomix-primitive-list-v1-RemoveRequest) | [RemoveResponse](#atomix-primitive-list-v1-RemoveResponse) | Remove removes an element from the list |
| Clear | [ClearRequest](#atomix-primitive-list-v1-ClearRequest) | [ClearResponse](#atomix-primitive-list-v1-ClearResponse) | Clear removes all elements from the list |
| Events | [EventsRequest](#atomix-primitive-list-v1-EventsRequest) | [EventsResponse](#atomix-primitive-list-v1-EventsResponse) stream | Events listens for change events |
| Elements | [ElementsRequest](#atomix-primitive-list-v1-ElementsRequest) | [ElementsResponse](#atomix-primitive-list-v1-ElementsResponse) stream | Elements streams all elements in the list |

 



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

