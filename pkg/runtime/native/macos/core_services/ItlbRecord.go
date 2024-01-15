//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ItlbRecord
*/

type ItlbRecord struct {
  ItlbNumber int16 `v8:"itlbNumber"`
  ItlbDate int16 `v8:"itlbDate"`
  ItlbSort int16 `v8:"itlbSort"`
  ItlbFlags int16 `v8:"itlbFlags"`
  ItlbToken int16 `v8:"itlbToken"`
  ItlbEncoding int16 `v8:"itlbEncoding"`
  ItlbLang int16 `v8:"itlbLang"`
  ItlbNumRep byte `v8:"itlbNumRep"`
  ItlbDateRep byte `v8:"itlbDateRep"`
  ItlbKeys int16 `v8:"itlbKeys"`
  ItlbIcon int16 `v8:"itlbIcon"`
}
