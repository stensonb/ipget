// Code generated by protoc-gen-gogo.
// source: message.proto
// DO NOT EDIT!

/*
Package bitswap_message_pb is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
*/
package bitswap_message_pb

import proto "github.com/noffle/ipget/Godeps/_workspace/src/github.com/gogo/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Message struct {
	Wantlist         *Message_Wantlist `protobuf:"bytes,1,opt,name=wantlist" json:"wantlist,omitempty"`
	Blocks           [][]byte          `protobuf:"bytes,2,rep,name=blocks" json:"blocks,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func (m *Message) GetWantlist() *Message_Wantlist {
	if m != nil {
		return m.Wantlist
	}
	return nil
}

func (m *Message) GetBlocks() [][]byte {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type Message_Wantlist struct {
	Entries          []*Message_Wantlist_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	Full             *bool                     `protobuf:"varint,2,opt,name=full" json:"full,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Message_Wantlist) Reset()         { *m = Message_Wantlist{} }
func (m *Message_Wantlist) String() string { return proto.CompactTextString(m) }
func (*Message_Wantlist) ProtoMessage()    {}

func (m *Message_Wantlist) GetEntries() []*Message_Wantlist_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *Message_Wantlist) GetFull() bool {
	if m != nil && m.Full != nil {
		return *m.Full
	}
	return false
}

type Message_Wantlist_Entry struct {
	Block            *string `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
	Priority         *int32  `protobuf:"varint,2,opt,name=priority" json:"priority,omitempty"`
	Cancel           *bool   `protobuf:"varint,3,opt,name=cancel" json:"cancel,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Message_Wantlist_Entry) Reset()         { *m = Message_Wantlist_Entry{} }
func (m *Message_Wantlist_Entry) String() string { return proto.CompactTextString(m) }
func (*Message_Wantlist_Entry) ProtoMessage()    {}

func (m *Message_Wantlist_Entry) GetBlock() string {
	if m != nil && m.Block != nil {
		return *m.Block
	}
	return ""
}

func (m *Message_Wantlist_Entry) GetPriority() int32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

func (m *Message_Wantlist_Entry) GetCancel() bool {
	if m != nil && m.Cancel != nil {
		return *m.Cancel
	}
	return false
}

func init() {
}