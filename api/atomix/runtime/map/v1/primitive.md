# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/map/v1/primitive.proto](#atomix_runtime_map_v1_primitive-proto)
    - [ClearRequest](#atomix-runtime-map-v1-ClearRequest)
    - [ClearResponse](#atomix-runtime-map-v1-ClearResponse)
    - [EntriesRequest](#atomix-runtime-map-v1-EntriesRequest)
    - [EntriesResponse](#atomix-runtime-map-v1-EntriesResponse)
    - [Entry](#atomix-runtime-map-v1-Entry)
    - [Event](#atomix-runtime-map-v1-Event)
    - [EventsRequest](#atomix-runtime-map-v1-EventsRequest)
    - [EventsResponse](#atomix-runtime-map-v1-EventsResponse)
    - [GetRequest](#atomix-runtime-map-v1-GetRequest)
    - [GetResponse](#atomix-runtime-map-v1-GetResponse)
    - [Key](#atomix-runtime-map-v1-Key)
    - [Precondition](#atomix-runtime-map-v1-Precondition)
    - [PutRequest](#atomix-runtime-map-v1-PutRequest)
    - [PutResponse](#atomix-runtime-map-v1-PutResponse)
    - [RemoveRequest](#atomix-runtime-map-v1-RemoveRequest)
    - [RemoveResponse](#atomix-runtime-map-v1-RemoveResponse)
    - [SizeRequest](#atomix-runtime-map-v1-SizeRequest)
    - [SizeResponse](#atomix-runtime-map-v1-SizeResponse)
    - [Value](#atomix-runtime-map-v1-Value)
  
    - [Event.Type](#atomix-runtime-map-v1-Event-Type)
  
    - [Map](#atomix-runtime-map-v1-Map)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_map_v1_primitive-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/map/v1/primitive.proto



<a name="atomix-runtime-map-v1-ClearRequest"></a>

### ClearRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |






<a name="atomix-runtime-map-v1-ClearResponse"></a>

### ClearResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |






<a name="atomix-runtime-map-v1-EntriesRequest"></a>

### EntriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |






<a name="atomix-runtime-map-v1-EntriesResponse"></a>

### EntriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| entry | [Entry](#atomix-runtime-map-v1-Entry) |  |  |






<a name="atomix-runtime-map-v1-Entry"></a>

### Entry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [Key](#atomix-runtime-map-v1-Key) |  |  |
| value | [Value](#atomix-runtime-map-v1-Value) |  |  |






<a name="atomix-runtime-map-v1-Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Event.Type](#atomix-runtime-map-v1-Event-Type) |  |  |
| entry | [Entry](#atomix-runtime-map-v1-Entry) |  |  |






<a name="atomix-runtime-map-v1-EventsRequest"></a>

### EventsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |
| key | [string](#string) |  |  |
| replay | [bool](#bool) |  |  |






<a name="atomix-runtime-map-v1-EventsResponse"></a>

### EventsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| event | [Event](#atomix-runtime-map-v1-Event) |  |  |






<a name="atomix-runtime-map-v1-GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |
| key | [string](#string) |  |  |






<a name="atomix-runtime-map-v1-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| entry | [Entry](#atomix-runtime-map-v1-Entry) |  |  |






<a name="atomix-runtime-map-v1-Key"></a>

### Key



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [atomix.runtime.meta.v1.ObjectMeta](#atomix-runtime-meta-v1-ObjectMeta) |  |  |
| key | [string](#string) |  |  |






<a name="atomix-runtime-map-v1-Precondition"></a>

### Precondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [atomix.runtime.meta.v1.ObjectMeta](#atomix-runtime-meta-v1-ObjectMeta) |  |  |






<a name="atomix-runtime-map-v1-PutRequest"></a>

### PutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |
| entry | [Entry](#atomix-runtime-map-v1-Entry) |  |  |
| preconditions | [Precondition](#atomix-runtime-map-v1-Precondition) | repeated |  |






<a name="atomix-runtime-map-v1-PutResponse"></a>

### PutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| entry | [Entry](#atomix-runtime-map-v1-Entry) |  |  |






<a name="atomix-runtime-map-v1-RemoveRequest"></a>

### RemoveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |
| key | [Key](#atomix-runtime-map-v1-Key) |  |  |
| preconditions | [Precondition](#atomix-runtime-map-v1-Precondition) | repeated |  |






<a name="atomix-runtime-map-v1-RemoveResponse"></a>

### RemoveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| entry | [Entry](#atomix-runtime-map-v1-Entry) |  |  |






<a name="atomix-runtime-map-v1-SizeRequest"></a>

### SizeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.RequestHeaders](#atomix-runtime-meta-v1-RequestHeaders) |  |  |






<a name="atomix-runtime-map-v1-SizeResponse"></a>

### SizeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.runtime.meta.v1.ResponseHeaders](#atomix-runtime-meta-v1-ResponseHeaders) |  |  |
| size | [uint32](#uint32) |  |  |






<a name="atomix-runtime-map-v1-Value"></a>

### Value



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bytes](#bytes) |  |  |
| ttl | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |





 


<a name="atomix-runtime-map-v1-Event-Type"></a>

### Event.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| INSERT | 1 |  |
| UPDATE | 2 |  |
| REMOVE | 3 |  |
| REPLAY | 4 |  |


 

 


<a name="atomix-runtime-map-v1-Map"></a>

### Map
Map is a service for a map primitive

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Size | [SizeRequest](#atomix-runtime-map-v1-SizeRequest) | [SizeResponse](#atomix-runtime-map-v1-SizeResponse) | Size returns the size of the map |
| Put | [PutRequest](#atomix-runtime-map-v1-PutRequest) | [PutResponse](#atomix-runtime-map-v1-PutResponse) | Put puts an entry into the map |
| Get | [GetRequest](#atomix-runtime-map-v1-GetRequest) | [GetResponse](#atomix-runtime-map-v1-GetResponse) | Get gets the entry for a key |
| Remove | [RemoveRequest](#atomix-runtime-map-v1-RemoveRequest) | [RemoveResponse](#atomix-runtime-map-v1-RemoveResponse) | Remove removes an entry from the map |
| Clear | [ClearRequest](#atomix-runtime-map-v1-ClearRequest) | [ClearResponse](#atomix-runtime-map-v1-ClearResponse) | Clear removes all entries from the map |
| Events | [EventsRequest](#atomix-runtime-map-v1-EventsRequest) | [EventsResponse](#atomix-runtime-map-v1-EventsResponse) stream | Events listens for change events |
| Entries | [EntriesRequest](#atomix-runtime-map-v1-EntriesRequest) | [EntriesResponse](#atomix-runtime-map-v1-EntriesResponse) stream | Entries lists all entries in the map |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers ??? if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers ??? if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
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

