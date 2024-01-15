//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IONamedValue
*/

type IONamedValue struct {
  Value int `v8:"value"`
  Name byte `v8:"name"`
}
