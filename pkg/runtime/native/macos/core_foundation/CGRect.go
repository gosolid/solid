//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CGRect
*/

type CGRect struct {
  Origin CGPoint `v8:"origin"`
  Size CGSize `v8:"size"`
}
