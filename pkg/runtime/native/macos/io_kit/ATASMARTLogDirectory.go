//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.ATASMARTLogDirectory
*/

type ATASMARTLogDirectory struct {
  SMARTLoggingVersion [2]byte `v8:"sMARTLoggingVersion"`
  Entries [255]ATASMARTLogEntry `v8:"entries"`
}
