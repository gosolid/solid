//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSEdgeInsets
*/

type NSEdgeInsets struct {
  Top float64 `v8:"top"`
  Left float64 `v8:"left"`
  Bottom float64 `v8:"bottom"`
  Right float64 `v8:"right"`
}
