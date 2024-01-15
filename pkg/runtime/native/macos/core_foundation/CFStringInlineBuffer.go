//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFStringInlineBuffer
*/

type CFStringInlineBuffer struct {
  Buffer [64]uint16 `v8:"buffer"`
  TheString *CFString `v8:"theString"`
  DirectUniCharBuffer uint16 `v8:"directUniCharBuffer"`
  DirectCStringBuffer byte `v8:"directCStringBuffer"`
  RangeToBuffer CFRange `v8:"rangeToBuffer"`
  BufferedRangeStart int64 `v8:"bufferedRangeStart"`
  BufferedRangeEnd int64 `v8:"bufferedRangeEnd"`
}
