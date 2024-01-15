//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOHIDCompletion
*/

type IOHIDCompletion struct {
  Target void `v8:"target"`
  Action *func(...any) (any, error) `v8:"action"`
  Parameter void `v8:"parameter"`
}
