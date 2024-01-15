//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.CGPSConverterCallbacks
*/

type CGPSConverterCallbacks struct {
  Version uint `v8:"version"`
  BeginDocument *func(...any) (any, error) `v8:"beginDocument"`
  EndDocument *func(...any) (any, error) `v8:"endDocument"`
  BeginPage *func(...any) (any, error) `v8:"beginPage"`
  EndPage *func(...any) (any, error) `v8:"endPage"`
  NoteProgress *func(...any) (any, error) `v8:"noteProgress"`
  NoteMessage *func(...any) (any, error) `v8:"noteMessage"`
  ReleaseInfo *func(...any) (any, error) `v8:"releaseInfo"`
}
