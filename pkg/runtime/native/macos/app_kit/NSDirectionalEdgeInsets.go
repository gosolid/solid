//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct AppKit.NSDirectionalEdgeInsets
*/

type NSDirectionalEdgeInsets struct {
  Top float64 `v8:"top"`
  Leading float64 `v8:"leading"`
  Bottom float64 `v8:"bottom"`
  Trailing float64 `v8:"trailing"`
}
