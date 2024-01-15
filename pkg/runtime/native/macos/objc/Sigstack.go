//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.sigstack
*/

type Sigstack struct {
  SsSp byte `v8:"ssSp"`
  SsOnstack int `v8:"ssOnstack"`
}
