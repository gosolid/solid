//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLDispatchThreadgroupsIndirectArguments
*/

type MTLDispatchThreadgroupsIndirectArguments struct {
  ThreadgroupsPerGrid [3]uint `v8:"threadgroupsPerGrid"`
}
