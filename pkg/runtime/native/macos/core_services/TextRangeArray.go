//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TextRangeArray
*/

type TextRangeArray struct {
  FNumOfRanges int16 `v8:"fNumOfRanges"`
  FRange [1]TextRange `v8:"fRange"`
}
