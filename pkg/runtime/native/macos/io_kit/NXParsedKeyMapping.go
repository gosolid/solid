//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXParsedKeyMapping
*/

type NXParsedKeyMapping struct {
  Shorts int16 `v8:"shorts"`
  KeyBits [256]byte `v8:"keyBits"`
  MaxMod int `v8:"maxMod"`
  ModDefs [16]*byte `v8:"modDefs"`
  NumDefs int `v8:"numDefs"`
  KeyDefs [256]*byte `v8:"keyDefs"`
  NumSeqs int `v8:"numSeqs"`
  SeqDefs [128]*byte `v8:"seqDefs"`
  NumSpecialKeys int `v8:"numSpecialKeys"`
  SpecialKeys [24]uint16 `v8:"specialKeys"`
  Mapping byte `v8:"mapping"`
  MappingLen int `v8:"mappingLen"`
}
