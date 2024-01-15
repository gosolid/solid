//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CQDProcs
*/

type CQDProcs struct {
  TextProc *func(...any) (any, error) `v8:"textProc"`
  LineProc *func(...any) (any, error) `v8:"lineProc"`
  RectProc *func(...any) (any, error) `v8:"rectProc"`
  RRectProc *func(...any) (any, error) `v8:"rRectProc"`
  OvalProc *func(...any) (any, error) `v8:"ovalProc"`
  ArcProc *func(...any) (any, error) `v8:"arcProc"`
  PolyProc *func(...any) (any, error) `v8:"polyProc"`
  RgnProc *func(...any) (any, error) `v8:"rgnProc"`
  BitsProc *func(...any) (any, error) `v8:"bitsProc"`
  CommentProc *func(...any) (any, error) `v8:"commentProc"`
  TxMeasProc *func(...any) (any, error) `v8:"txMeasProc"`
  GetPicProc *func(...any) (any, error) `v8:"getPicProc"`
  PutPicProc *func(...any) (any, error) `v8:"putPicProc"`
  OpcodeProc *func(...any) (any, error) `v8:"opcodeProc"`
  NewProc1 *func(...any) (any, error) `v8:"newProc1"`
  GlyphsProc *func(...any) (any, error) `v8:"glyphsProc"`
  PrinterStatusProc *func(...any) (any, error) `v8:"printerStatusProc"`
  NewProc4 *func(...any) (any, error) `v8:"newProc4"`
  NewProc5 *func(...any) (any, error) `v8:"newProc5"`
  NewProc6 *func(...any) (any, error) `v8:"newProc6"`
}
