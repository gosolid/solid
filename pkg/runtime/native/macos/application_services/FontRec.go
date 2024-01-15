//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.FontRec
*/

type FontRec struct {
  FontType int16 `v8:"fontType"`
  FirstChar int16 `v8:"firstChar"`
  LastChar int16 `v8:"lastChar"`
  WidMax int16 `v8:"widMax"`
  KernMax int16 `v8:"kernMax"`
  NDescent int16 `v8:"nDescent"`
  FRectWidth int16 `v8:"fRectWidth"`
  FRectHeight int16 `v8:"fRectHeight"`
  OwTLoc uint16 `v8:"owTLoc"`
  Ascent int16 `v8:"ascent"`
  Descent int16 `v8:"descent"`
  Leading int16 `v8:"leading"`
  RowWords int16 `v8:"rowWords"`
}
