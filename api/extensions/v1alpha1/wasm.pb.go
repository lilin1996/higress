// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extensions/v1alpha1/wasm.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// The phase in the filter chain where the plugin will be injected.
type PluginPhase int32

const (
	// Control plane decides where to insert the plugin. This will generally
	// be at the end of the filter chain, right before the Router.
	// Do not specify `PluginPhase` if the plugin is independent of others.
	PluginPhase_UNSPECIFIED_PHASE PluginPhase = 0
	// Insert plugin before Istio authentication filters.
	PluginPhase_AUTHN PluginPhase = 1
	// Insert plugin before Istio authorization filters and after Istio authentication filters.
	PluginPhase_AUTHZ PluginPhase = 2
	// Insert plugin before Istio stats filters and after Istio authorization filters.
	PluginPhase_STATS PluginPhase = 3
)

var PluginPhase_name = map[int32]string{
	0: "UNSPECIFIED_PHASE",
	1: "AUTHN",
	2: "AUTHZ",
	3: "STATS",
}

var PluginPhase_value = map[string]int32{
	"UNSPECIFIED_PHASE": 0,
	"AUTHN":             1,
	"AUTHZ":             2,
	"STATS":             3,
}

func (x PluginPhase) String() string {
	return proto.EnumName(PluginPhase_name, int32(x))
}

func (PluginPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d60b240916c4e18, []int{0}
}

// The pull behaviour to be applied when fetching an OCI image,
// mirroring K8s behaviour.
//
// <!--
// buf:lint:ignore ENUM_VALUE_UPPER_SNAKE_CASE
// -->
type PullPolicy int32

const (
	// Defaults to IfNotPresent, except for OCI images with tag `latest`, for which
	// the default will be Always.
	PullPolicy_UNSPECIFIED_POLICY PullPolicy = 0
	// If an existing version of the image has been pulled before, that
	// will be used. If no version of the image is present locally, we
	// will pull the latest version.
	PullPolicy_IfNotPresent PullPolicy = 1
	// We will always pull the latest version of an image when applying
	// this plugin.
	PullPolicy_Always PullPolicy = 2
)

var PullPolicy_name = map[int32]string{
	0: "UNSPECIFIED_POLICY",
	1: "IfNotPresent",
	2: "Always",
}

var PullPolicy_value = map[string]int32{
	"UNSPECIFIED_POLICY": 0,
	"IfNotPresent":       1,
	"Always":             2,
}

func (x PullPolicy) String() string {
	return proto.EnumName(PullPolicy_name, int32(x))
}

func (PullPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d60b240916c4e18, []int{1}
}

// <!-- crd generation tags
// +cue-gen:WasmPlugin:groupName:extensions.higress.io
// +cue-gen:WasmPlugin:version:v1alpha1
// +cue-gen:WasmPlugin:storageVersion
// +cue-gen:WasmPlugin:annotations:helm.sh/resource-policy=keep
// +cue-gen:WasmPlugin:subresource:status
// +cue-gen:WasmPlugin:scope:Namespaced
// +cue-gen:WasmPlugin:resource:categories=higress-io,extensions-higress-io
// +cue-gen:WasmPlugin:preserveUnknownFields:pluginConfig,defaultConfig,matchRules.[].config
// +cue-gen:WasmPlugin:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=extensions.higress.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type WasmPlugin struct {
	// URL of a Wasm module or OCI container. If no scheme is present,
	// defaults to `oci://`, referencing an OCI image. Other valid schemes
	// are `file://` for referencing .wasm module files present locally
	// within the proxy container, and `http[s]://` for .wasm module files
	// hosted remotely.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// SHA256 checksum that will be used to verify Wasm module or OCI container.
	// If the `url` field already references a SHA256 (using the `@sha256:`
	// notation), it must match the value of this field. If an OCI image is
	// referenced by tag and this field is set, its checksum will be verified
	// against the contents of this field after pulling.
	Sha256 string `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
	// The pull behaviour to be applied when fetching an OCI image. Only
	// relevant when images are referenced by tag instead of SHA. Defaults
	// to IfNotPresent, except when an OCI image is referenced in the `url`
	// and the `latest` tag is used, in which case `Always` is the default,
	// mirroring K8s behaviour.
	// Setting is ignored if `url` field is referencing a Wasm module directly
	// using `file://` or `http[s]://`
	ImagePullPolicy PullPolicy `protobuf:"varint,4,opt,name=image_pull_policy,json=imagePullPolicy,proto3,enum=higress.extensions.v1alpha1.PullPolicy" json:"image_pull_policy,omitempty"`
	// Credentials to use for OCI image pulling.
	// Name of a K8s Secret in the same namespace as the `WasmPlugin` that
	// contains a docker pull secret which is to be used to authenticate
	// against the registry when pulling the image.
	ImagePullSecret string `protobuf:"bytes,5,opt,name=image_pull_secret,json=imagePullSecret,proto3" json:"image_pull_secret,omitempty"`
	// Public key that will be used to verify signatures of signed OCI images
	// or Wasm modules. Must be supplied in PEM format.
	VerificationKey string `protobuf:"bytes,6,opt,name=verification_key,json=verificationKey,proto3" json:"verification_key,omitempty"`
	// The configuration that will be passed on to the plugin.
	PluginConfig *types.Struct `protobuf:"bytes,7,opt,name=plugin_config,json=pluginConfig,proto3" json:"plugin_config,omitempty"`
	// The plugin name to be used in the Envoy configuration (used to be called
	// `rootID`). Some .wasm modules might require this value to select the Wasm
	// plugin to execute.
	PluginName string `protobuf:"bytes,8,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	// Determines where in the filter chain this `WasmPlugin` is to be injected.
	Phase PluginPhase `protobuf:"varint,9,opt,name=phase,proto3,enum=higress.extensions.v1alpha1.PluginPhase" json:"phase,omitempty"`
	// Determines ordering of `WasmPlugins` in the same `phase`.
	// When multiple `WasmPlugins` are applied to the same workload in the
	// same `phase`, they will be applied by priority, in descending order.
	// If `priority` is not set, or two `WasmPlugins` exist with the same
	// value, the ordering will be deterministically derived from name and
	// namespace of the `WasmPlugins`. Defaults to `0`.
	Priority *types.Int64Value `protobuf:"bytes,10,opt,name=priority,proto3" json:"priority,omitempty"`
	// Extended by Higress, the default configuration takes effect globally
	DefaultConfig *types.Struct `protobuf:"bytes,101,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	// Extended by Higress, matching rules take effect
	MatchRules           []*MatchRule `protobuf:"bytes,102,rep,name=match_rules,json=matchRules,proto3" json:"match_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WasmPlugin) Reset()         { *m = WasmPlugin{} }
func (m *WasmPlugin) String() string { return proto.CompactTextString(m) }
func (*WasmPlugin) ProtoMessage()    {}
func (*WasmPlugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d60b240916c4e18, []int{0}
}
func (m *WasmPlugin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WasmPlugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WasmPlugin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WasmPlugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmPlugin.Merge(m, src)
}
func (m *WasmPlugin) XXX_Size() int {
	return m.Size()
}
func (m *WasmPlugin) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmPlugin.DiscardUnknown(m)
}

var xxx_messageInfo_WasmPlugin proto.InternalMessageInfo

func (m *WasmPlugin) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *WasmPlugin) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *WasmPlugin) GetImagePullPolicy() PullPolicy {
	if m != nil {
		return m.ImagePullPolicy
	}
	return PullPolicy_UNSPECIFIED_POLICY
}

func (m *WasmPlugin) GetImagePullSecret() string {
	if m != nil {
		return m.ImagePullSecret
	}
	return ""
}

func (m *WasmPlugin) GetVerificationKey() string {
	if m != nil {
		return m.VerificationKey
	}
	return ""
}

func (m *WasmPlugin) GetPluginConfig() *types.Struct {
	if m != nil {
		return m.PluginConfig
	}
	return nil
}

func (m *WasmPlugin) GetPluginName() string {
	if m != nil {
		return m.PluginName
	}
	return ""
}

func (m *WasmPlugin) GetPhase() PluginPhase {
	if m != nil {
		return m.Phase
	}
	return PluginPhase_UNSPECIFIED_PHASE
}

func (m *WasmPlugin) GetPriority() *types.Int64Value {
	if m != nil {
		return m.Priority
	}
	return nil
}

func (m *WasmPlugin) GetDefaultConfig() *types.Struct {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func (m *WasmPlugin) GetMatchRules() []*MatchRule {
	if m != nil {
		return m.MatchRules
	}
	return nil
}

// Extended by Higress
type MatchRule struct {
	Ingress              []string      `protobuf:"bytes,1,rep,name=ingress,proto3" json:"ingress,omitempty"`
	Domain               []string      `protobuf:"bytes,2,rep,name=domain,proto3" json:"domain,omitempty"`
	Config               *types.Struct `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MatchRule) Reset()         { *m = MatchRule{} }
func (m *MatchRule) String() string { return proto.CompactTextString(m) }
func (*MatchRule) ProtoMessage()    {}
func (*MatchRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d60b240916c4e18, []int{1}
}
func (m *MatchRule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MatchRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MatchRule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MatchRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchRule.Merge(m, src)
}
func (m *MatchRule) XXX_Size() int {
	return m.Size()
}
func (m *MatchRule) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchRule.DiscardUnknown(m)
}

var xxx_messageInfo_MatchRule proto.InternalMessageInfo

func (m *MatchRule) GetIngress() []string {
	if m != nil {
		return m.Ingress
	}
	return nil
}

func (m *MatchRule) GetDomain() []string {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *MatchRule) GetConfig() *types.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterEnum("higress.extensions.v1alpha1.PluginPhase", PluginPhase_name, PluginPhase_value)
	proto.RegisterEnum("higress.extensions.v1alpha1.PullPolicy", PullPolicy_name, PullPolicy_value)
	proto.RegisterType((*WasmPlugin)(nil), "higress.extensions.v1alpha1.WasmPlugin")
	proto.RegisterType((*MatchRule)(nil), "higress.extensions.v1alpha1.MatchRule")
}

func init() { proto.RegisterFile("extensions/v1alpha1/wasm.proto", fileDescriptor_4d60b240916c4e18) }

var fileDescriptor_4d60b240916c4e18 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x76, 0xeb, 0xd6, 0xd3, 0x6d, 0x64, 0x96, 0x18, 0xd6, 0x86, 0x4a, 0xb5, 0x0b,
	0x28, 0xbb, 0x48, 0xb4, 0x02, 0xe3, 0x06, 0x4d, 0x74, 0xa3, 0xb0, 0x0a, 0x28, 0x51, 0xb2, 0x81,
	0xd8, 0x4d, 0xe5, 0x66, 0x6e, 0x6a, 0xe1, 0xc4, 0x91, 0xed, 0x6c, 0xf4, 0xf9, 0xb8, 0xe1, 0x92,
	0x47, 0x40, 0x7b, 0x12, 0x54, 0x27, 0xd9, 0x07, 0xa0, 0xde, 0x9d, 0x8f, 0xdf, 0x39, 0xf9, 0xff,
	0x8f, 0x1c, 0x68, 0xd2, 0xef, 0x9a, 0x26, 0x8a, 0x89, 0x44, 0xb9, 0x17, 0x7b, 0x84, 0xa7, 0x13,
	0xb2, 0xe7, 0x5e, 0x12, 0x15, 0x3b, 0xa9, 0x14, 0x5a, 0xa0, 0xed, 0x09, 0x8b, 0x24, 0x55, 0xca,
	0xb9, 0xe1, 0x9c, 0x92, 0xdb, 0x6a, 0x46, 0x42, 0x44, 0x9c, 0xba, 0x06, 0x1d, 0x65, 0x63, 0xf7,
	0x52, 0x92, 0x34, 0xa5, 0x52, 0xe5, 0xc3, 0x5b, 0x0f, 0xff, 0xee, 0x2b, 0x2d, 0xb3, 0x50, 0xe7,
	0xdd, 0x9d, 0x1f, 0x8b, 0x00, 0x5f, 0x88, 0x8a, 0x3d, 0x9e, 0x45, 0x2c, 0x41, 0x36, 0x54, 0x33,
	0xc9, 0x71, 0xa5, 0x65, 0xb5, 0xeb, 0xfe, 0x2c, 0x44, 0x9b, 0x50, 0x53, 0x13, 0xd2, 0x79, 0xb1,
	0x8f, 0xab, 0xa6, 0x58, 0x64, 0x28, 0x80, 0x0d, 0x16, 0x93, 0x88, 0x0e, 0xd3, 0x8c, 0xf3, 0x61,
	0x2a, 0x38, 0x0b, 0xa7, 0x78, 0xb1, 0x65, 0xb5, 0xd7, 0x3b, 0x4f, 0x9c, 0x39, 0x7a, 0x1d, 0x2f,
	0xe3, 0xdc, 0x33, 0xb8, 0x7f, 0xcf, 0x6c, 0xb8, 0x29, 0xa0, 0xdd, 0x3b, 0x4b, 0x15, 0x0d, 0x25,
	0xd5, 0x78, 0xc9, 0x7c, 0xf7, 0x86, 0x0d, 0x4c, 0x19, 0x3d, 0x05, 0xfb, 0x82, 0x4a, 0x36, 0x66,
	0x21, 0xd1, 0x4c, 0x24, 0xc3, 0x6f, 0x74, 0x8a, 0x6b, 0x39, 0x7a, 0xbb, 0xfe, 0x9e, 0x4e, 0xd1,
	0x2b, 0x58, 0x4b, 0x8d, 0xbf, 0x61, 0x28, 0x92, 0x31, 0x8b, 0xf0, 0x72, 0xcb, 0x6a, 0x37, 0x3a,
	0x0f, 0x9c, 0xfc, 0x34, 0x4e, 0x79, 0x1a, 0x27, 0x30, 0xa7, 0xf1, 0x57, 0x73, 0xfa, 0xc8, 0xc0,
	0xe8, 0x11, 0x34, 0x8a, 0xe9, 0x84, 0xc4, 0x14, 0xaf, 0x98, 0x6f, 0x40, 0x5e, 0x1a, 0x90, 0x98,
	0xa2, 0x03, 0x58, 0x4a, 0x27, 0x44, 0x51, 0x5c, 0x37, 0xf6, 0xdb, 0xf3, 0xed, 0x9b, 0x39, 0x6f,
	0xc6, 0xfb, 0xf9, 0x18, 0x7a, 0x09, 0x2b, 0xa9, 0x64, 0x42, 0x32, 0x3d, 0xc5, 0x60, 0x94, 0x6d,
	0xff, 0xa3, 0xac, 0x9f, 0xe8, 0xfd, 0xe7, 0x9f, 0x09, 0xcf, 0xa8, 0x7f, 0x0d, 0xa3, 0x03, 0x58,
	0x3f, 0xa7, 0x63, 0x92, 0x71, 0x5d, 0x1a, 0xa3, 0xf3, 0x8d, 0xad, 0x15, 0x78, 0xe1, 0xec, 0x1d,
	0x34, 0x62, 0xa2, 0xc3, 0xc9, 0x50, 0x66, 0x9c, 0x2a, 0x3c, 0x6e, 0x55, 0xdb, 0x8d, 0xce, 0xe3,
	0xb9, 0xf2, 0x3f, 0xce, 0x78, 0x3f, 0xe3, 0xd4, 0x87, 0xb8, 0x0c, 0xd5, 0x4e, 0x02, 0xf5, 0xeb,
	0x06, 0xc2, 0xb0, 0xcc, 0x12, 0xb3, 0x01, 0x5b, 0xad, 0x6a, 0xbb, 0xee, 0x97, 0xe9, 0xec, 0x2d,
	0x9d, 0x8b, 0x98, 0xb0, 0x04, 0x57, 0x4c, 0xa3, 0xc8, 0x90, 0x0b, 0xb5, 0x42, 0x7f, 0x75, 0xbe,
	0xfe, 0x02, 0xdb, 0xed, 0x41, 0xe3, 0xd6, 0x1d, 0xd1, 0x7d, 0xd8, 0x38, 0x1d, 0x04, 0x5e, 0xef,
	0xa8, 0xff, 0xb6, 0xdf, 0x7b, 0x33, 0xf4, 0x8e, 0xbb, 0x41, 0xcf, 0x5e, 0x40, 0x75, 0x58, 0xea,
	0x9e, 0x9e, 0x1c, 0x0f, 0x6c, 0xab, 0x0c, 0xcf, 0xec, 0xca, 0x2c, 0x0c, 0x4e, 0xba, 0x27, 0x81,
	0x5d, 0xdd, 0x3d, 0x04, 0xb8, 0xf5, 0xf8, 0x36, 0x01, 0xdd, 0xd9, 0xf2, 0xe9, 0x43, 0xff, 0xe8,
	0xab, 0xbd, 0x80, 0x6c, 0x58, 0xed, 0x8f, 0x07, 0x42, 0x7b, 0x92, 0x2a, 0x9a, 0x68, 0xdb, 0x42,
	0x00, 0xb5, 0x2e, 0xbf, 0x24, 0x53, 0x65, 0x57, 0x0e, 0x5f, 0xff, 0xbc, 0x6a, 0x5a, 0xbf, 0xae,
	0x9a, 0xd6, 0xef, 0xab, 0xa6, 0x75, 0xd6, 0x89, 0x98, 0x9e, 0x64, 0x23, 0x27, 0x14, 0xb1, 0x4b,
	0x38, 0x1b, 0x91, 0x11, 0x71, 0x8b, 0x73, 0xba, 0x24, 0x65, 0xee, 0x7f, 0x7e, 0xf4, 0x51, 0xcd,
	0xb8, 0x7c, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xe0, 0x3d, 0x06, 0x06, 0x04, 0x00, 0x00,
}

func (m *WasmPlugin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WasmPlugin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WasmPlugin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MatchRules) > 0 {
		for iNdEx := len(m.MatchRules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MatchRules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWasm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6
			i--
			dAtA[i] = 0xb2
		}
	}
	if m.DefaultConfig != nil {
		{
			size, err := m.DefaultConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWasm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xaa
	}
	if m.Priority != nil {
		{
			size, err := m.Priority.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWasm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.Phase != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.Phase))
		i--
		dAtA[i] = 0x48
	}
	if len(m.PluginName) > 0 {
		i -= len(m.PluginName)
		copy(dAtA[i:], m.PluginName)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.PluginName)))
		i--
		dAtA[i] = 0x42
	}
	if m.PluginConfig != nil {
		{
			size, err := m.PluginConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWasm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.VerificationKey) > 0 {
		i -= len(m.VerificationKey)
		copy(dAtA[i:], m.VerificationKey)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.VerificationKey)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ImagePullSecret) > 0 {
		i -= len(m.ImagePullSecret)
		copy(dAtA[i:], m.ImagePullSecret)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.ImagePullSecret)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ImagePullPolicy != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.ImagePullPolicy))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Sha256) > 0 {
		i -= len(m.Sha256)
		copy(dAtA[i:], m.Sha256)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.Sha256)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *MatchRule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MatchRule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MatchRule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWasm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Domain) > 0 {
		for iNdEx := len(m.Domain) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Domain[iNdEx])
			copy(dAtA[i:], m.Domain[iNdEx])
			i = encodeVarintWasm(dAtA, i, uint64(len(m.Domain[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Ingress) > 0 {
		for iNdEx := len(m.Ingress) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Ingress[iNdEx])
			copy(dAtA[i:], m.Ingress[iNdEx])
			i = encodeVarintWasm(dAtA, i, uint64(len(m.Ingress[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintWasm(dAtA []byte, offset int, v uint64) int {
	offset -= sovWasm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WasmPlugin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	l = len(m.Sha256)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	if m.ImagePullPolicy != 0 {
		n += 1 + sovWasm(uint64(m.ImagePullPolicy))
	}
	l = len(m.ImagePullSecret)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	l = len(m.VerificationKey)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	if m.PluginConfig != nil {
		l = m.PluginConfig.Size()
		n += 1 + l + sovWasm(uint64(l))
	}
	l = len(m.PluginName)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	if m.Phase != 0 {
		n += 1 + sovWasm(uint64(m.Phase))
	}
	if m.Priority != nil {
		l = m.Priority.Size()
		n += 1 + l + sovWasm(uint64(l))
	}
	if m.DefaultConfig != nil {
		l = m.DefaultConfig.Size()
		n += 2 + l + sovWasm(uint64(l))
	}
	if len(m.MatchRules) > 0 {
		for _, e := range m.MatchRules {
			l = e.Size()
			n += 2 + l + sovWasm(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MatchRule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ingress) > 0 {
		for _, s := range m.Ingress {
			l = len(s)
			n += 1 + l + sovWasm(uint64(l))
		}
	}
	if len(m.Domain) > 0 {
		for _, s := range m.Domain {
			l = len(s)
			n += 1 + l + sovWasm(uint64(l))
		}
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovWasm(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWasm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWasm(x uint64) (n int) {
	return sovWasm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WasmPlugin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WasmPlugin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WasmPlugin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha256", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha256 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImagePullPolicy", wireType)
			}
			m.ImagePullPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ImagePullPolicy |= PullPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImagePullSecret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImagePullSecret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PluginConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PluginConfig == nil {
				m.PluginConfig = &types.Struct{}
			}
			if err := m.PluginConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PluginName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PluginName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phase", wireType)
			}
			m.Phase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Phase |= PluginPhase(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Priority == nil {
				m.Priority = &types.Int64Value{}
			}
			if err := m.Priority.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 101:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultConfig == nil {
				m.DefaultConfig = &types.Struct{}
			}
			if err := m.DefaultConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 102:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchRules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchRules = append(m.MatchRules, &MatchRule{})
			if err := m.MatchRules[len(m.MatchRules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MatchRule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MatchRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MatchRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ingress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ingress = append(m.Ingress, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = append(m.Domain, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &types.Struct{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWasm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWasm
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWasm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWasm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWasm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWasm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWasm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWasm = fmt.Errorf("proto: unexpected end of group")
)