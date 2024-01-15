//js:package native/macos/color-sync
package color_sync

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ColorSync.ColorSyncMD5
*/

type ColorSyncMD5 struct {
  Digest [16]byte `v8:"digest"`
}
