//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLTextureSwizzleChannels
*/

type MTLTextureSwizzleChannels struct {
  Red int `v8:"red"`
  Green int `v8:"green"`
  Blue int `v8:"blue"`
  Alpha int `v8:"alpha"`
}
