//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.timeval
*/

type Timeval struct {
  TvSec int64 `v8:"tvSec"`
  TvUsec int `v8:"tvUsec"`
}
