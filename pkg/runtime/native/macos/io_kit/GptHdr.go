//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.gpt_hdr
*/

type GptHdr struct {
  HdrSig [8]byte `v8:"hdrSig"`
  HdrRevision uint `v8:"hdrRevision"`
  HdrSize uint `v8:"hdrSize"`
  HdrCrcSelf uint `v8:"hdrCrcSelf"`
  Reserved uint `v8:"reserved"`
  HdrLbaSelf uint64 `v8:"hdrLbaSelf"`
  HdrLbaAlt uint64 `v8:"hdrLbaAlt"`
  HdrLbaStart uint64 `v8:"hdrLbaStart"`
  HdrLbaEnd uint64 `v8:"hdrLbaEnd"`
  HdrUuid [16]byte `v8:"hdrUuid"`
  HdrLbaTable uint64 `v8:"hdrLbaTable"`
  HdrEntries uint `v8:"hdrEntries"`
  HdrEntsz uint `v8:"hdrEntsz"`
  HdrCrcTable uint `v8:"hdrCrcTable"`
  Padding uint `v8:"padding"`
}
