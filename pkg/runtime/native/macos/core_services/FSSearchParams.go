//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FSSearchParams
*/

type FSSearchParams struct {
  SearchTime int `v8:"searchTime"`
  SearchBits uint `v8:"searchBits"`
  SearchNameLength uint64 `v8:"searchNameLength"`
  SearchName uint16 `v8:"searchName"`
  SearchInfo1 FSCatalogInfo `v8:"searchInfo1"`
  SearchInfo2 FSCatalogInfo `v8:"searchInfo2"`
}
