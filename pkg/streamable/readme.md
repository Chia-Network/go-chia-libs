# Streamable

This package implements the chia streamable format. Not all aspects of the streamable format are fully implemented, and 
support for more types are added as protocol messages are added to this package. This is not intended to be used in
consensus critical applications and there may be unexpected errors for untested streamable objects.

For more information on the streamable format, see the [streamable docs](https://docs.chia.net/serialization-protocol?_highlight=strea#streamable-format)

## How to Use

When defining structs that are streamable, the order of the fields is extremely important, and should match the order
of the fields in [chia-blockchain](https://github.com/chia-network/chia-blockchain). To support struct fields that are
not defined in chia-blockchain, streamable objects require a `streamable` tag on each field of the struct that should be
streamed.

**Example Type Definition:**

```go
// TimestampedPeerInfo contains information about peers with timestamps
type TimestampedPeerInfo struct {
	Host      string `streamable:""`
	Port      uint16 `streamable:""`
	Timestamp uint64 `streamable:""`
}
```

For a given streamable object, the interface is very similar to json marshal/unmarshal. 

**Encode to Bytes**

```go
peerInfo := &TimestampedPeerInfo{....}
bytes, err := streamable.Marshal(peerInfo)
```

**Decode to Struct**

```go
peerInfo := &TimestampedPeerInfo{}
err := streamable.Unmarshal(bytes, peerInfo)
```
