//js:package native/macos/open-gl
package open_gl

//go:generate go run github.com/grexie/isolates/codegen

/*
struct OpenGL._CGLContextObject
*/

type CGLContextObject struct {
  Rend *GLIContextRec `v8:"rend"`
  Disp GLIFunctionDispatch `v8:"disp"`
  Priv *CGLPrivateObject `v8:"priv"`
  Stak void `v8:"stak"`
}
