//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.sigvec
*/

type Sigvec struct {
  SvHandler func(...any) (any, error) `v8:"svHandler"`
  SvMask int `v8:"svMask"`
  SvFlags int `v8:"svFlags"`
}
