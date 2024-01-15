//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLClearColor
*/

type MTLClearColor struct {
  Red float64 `v8:"red"`
  Green float64 `v8:"green"`
  Blue float64 `v8:"blue"`
  Alpha float64 `v8:"alpha"`
}
