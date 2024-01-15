//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAsyncCompletionContent
*/

type IOAsyncCompletionContent struct {
  Result int `v8:"result"`
  Args []*void `v8:"args"`
}
