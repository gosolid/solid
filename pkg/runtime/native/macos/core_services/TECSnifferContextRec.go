//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECSnifferContextRec
*/

type TECSnifferContextRec struct {
  PluginRec *byte `v8:"pluginRec"`
  Encoding uint `v8:"encoding"`
  MaxErrors uint64 `v8:"maxErrors"`
  MaxFeatures uint64 `v8:"maxFeatures"`
  TextInputBuffer *byte `v8:"textInputBuffer"`
  TextInputBufferEnd *byte `v8:"textInputBufferEnd"`
  NumFeatures uint64 `v8:"numFeatures"`
  NumErrors uint64 `v8:"numErrors"`
  ContextRefCon *void `v8:"contextRefCon"`
  SniffProc *func(...any) (any, error) `v8:"sniffProc"`
  ClearContextInfoProc *func(...any) (any, error) `v8:"clearContextInfoProc"`
  PluginState TECPluginStateRec `v8:"pluginState"`
}
