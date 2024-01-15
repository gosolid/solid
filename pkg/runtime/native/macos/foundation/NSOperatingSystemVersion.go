//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSOperatingSystemVersion
*/

type NSOperatingSystemVersion struct {
  MajorVersion int64 `v8:"majorVersion"`
  MinorVersion int64 `v8:"minorVersion"`
  PatchVersion int64 `v8:"patchVersion"`
}
