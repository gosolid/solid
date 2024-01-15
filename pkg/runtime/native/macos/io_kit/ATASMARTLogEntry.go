//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.ATASMARTLogEntry
*/

type ATASMARTLogEntry struct {
  NumberOfSectors byte `v8:"numberOfSectors"`
  Reserved byte `v8:"reserved"`
}
