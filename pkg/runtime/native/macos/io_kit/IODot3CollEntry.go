//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODot3CollEntry
*/

type IODot3CollEntry struct {
  CollFrequencies [16]uint `v8:"collFrequencies"`
}
