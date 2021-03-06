# protoc-gen-jsonshim

`protoc-gen-jsonshim` is a plugin for protoc which generates `MarshalJSON()` and
`UnmarshalJSON()` functions for `.pb.go` types that contain `oneof` fields.  The
implementations simply redirect to `MarshalJSONPB()` and `UnmarshalJSONPB()`
respectively.  This allows the types to be used with native Go JSON
serialization.

## Usage

Add the executable to your system's PATH, for example:

`$ go install istio.io/tools/cmd/protoc-gen-jsonshim`

Add the `jsonshim_out` option to your `protoc` command line, for example:

`$ protoc --jsonshim_out=:output/path ...`

## Quirks

The generator is built off `protoc-gen-gogo`, which adds some imports that are
not used, along with copying over file comments from the source file.  See below
for an example of the genrated code.

## Examples Of Generated Code

```go
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/destination_rule.proto

package v1alpha3 // import "istio.io/api/networking/v1alpha3"

import bytes "bytes"
import github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler supporting oneof fields for LoadBalancerSettings
func (this *LoadBalancerSettings) MarshalJSON() ([]byte, error) {
    str, err := DestinationRuleMarshaler.MarshalToString(this)
    return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for LoadBalancerSettings
func (this *LoadBalancerSettings) UnmarshalJSON(b []byte) error {
    return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for LoadBalancerSettings_ConsistentHashLB
func (this *LoadBalancerSettings_ConsistentHashLB) MarshalJSON() ([]byte, error) {
    str, err := DestinationRuleMarshaler.MarshalToString(this)
    return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for LoadBalancerSettings_ConsistentHashLB
func (this *LoadBalancerSettings_ConsistentHashLB) UnmarshalJSON(b []byte) error {
    return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
    DestinationRuleMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
    DestinationRuleUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
```
