//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.gpt_ent
*/

type GptEnt struct {
  EntType [16]byte `v8:"entType"`
  EntUuid [16]byte `v8:"entUuid"`
  EntLbaStart uint64 `v8:"entLbaStart"`
  EntLbaEnd uint64 `v8:"entLbaEnd"`
  EntAttr uint64 `v8:"entAttr"`
  EntName [36]uint16 `v8:"entName"`
}
