//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.Itl4Rec
*/

type Itl4Rec struct {
  Flags int16 `v8:"flags"`
  ResourceType int `v8:"resourceType"`
  ResourceNum int16 `v8:"resourceNum"`
  Version int16 `v8:"version"`
  ResHeader1 int `v8:"resHeader1"`
  ResHeader2 int `v8:"resHeader2"`
  NumTables int16 `v8:"numTables"`
  MapOffset int `v8:"mapOffset"`
  StrOffset int `v8:"strOffset"`
  FetchOffset int `v8:"fetchOffset"`
  UnTokenOffset int `v8:"unTokenOffset"`
  DefPartsOffset int `v8:"defPartsOffset"`
  ResOffset6 int `v8:"resOffset6"`
  ResOffset7 int `v8:"resOffset7"`
  ResOffset8 int `v8:"resOffset8"`
}
