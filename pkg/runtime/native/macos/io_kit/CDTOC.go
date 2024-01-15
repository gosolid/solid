//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.CDTOC
*/

type CDTOC struct {
  Length uint16 `v8:"length"`
  SessionFirst byte `v8:"sessionFirst"`
  SessionLast byte `v8:"sessionLast"`
  Descriptors [0]CDTOCDescriptor `v8:"descriptors"`
}
