//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/http/cdn_loop/v3/cdn_loop.proto

package cdn_loopv3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *CdnLoopConfig) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CdnLoopConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *CdnLoopConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.MaxAllowedOccurrences != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.MaxAllowedOccurrences))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CdnId) > 0 {
		i -= len(m.CdnId)
		copy(dAtA[i:], m.CdnId)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.CdnId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CdnLoopConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CdnId)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.MaxAllowedOccurrences != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.MaxAllowedOccurrences))
	}
	n += len(m.unknownFields)
	return n
}