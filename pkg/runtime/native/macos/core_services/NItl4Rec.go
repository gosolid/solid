//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.NItl4Rec
*/

type NItl4Rec struct {
  Flags int16 `v8:"flags"`
  ResourceType int `v8:"resourceType"`
  ResourceNum int16 `v8:"resourceNum"`
  Version int16 `v8:"version"`
  Format int16 `v8:"format"`
  ResHeader int16 `v8:"resHeader"`
  ResHeader2 int `v8:"resHeader2"`
  NumTables int16 `v8:"numTables"`
  MapOffset int `v8:"mapOffset"`
  StrOffset int `v8:"strOffset"`
  FetchOffset int `v8:"fetchOffset"`
  UnTokenOffset int `v8:"unTokenOffset"`
  DefPartsOffset int `v8:"defPartsOffset"`
  WhtSpListOffset int `v8:"whtSpListOffset"`
  ResOffset7 int `v8:"resOffset7"`
  ResOffset8 int `v8:"resOffset8"`
  ResLength1 int16 `v8:"resLength1"`
  ResLength2 int16 `v8:"resLength2"`
  ResLength3 int16 `v8:"resLength3"`
  UnTokenLength int16 `v8:"unTokenLength"`
  DefPartsLength int16 `v8:"defPartsLength"`
  WhtSpListLength int16 `v8:"whtSpListLength"`
  ResLength7 int16 `v8:"resLength7"`
  ResLength8 int16 `v8:"resLength8"`
}
