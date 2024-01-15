//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECConverterContextRec
*/

type TECConverterContextRec struct {
  PluginRec *byte `v8:"pluginRec"`
  SourceEncoding uint `v8:"sourceEncoding"`
  DestEncoding uint `v8:"destEncoding"`
  Reserved1 uint `v8:"reserved1"`
  Reserved2 uint `v8:"reserved2"`
  BufferContext TECBufferContextRec `v8:"bufferContext"`
  ContextRefCon *void `v8:"contextRefCon"`
  ConversionProc *func(...any) (any, error) `v8:"conversionProc"`
  FlushProc *func(...any) (any, error) `v8:"flushProc"`
  ClearContextInfoProc *func(...any) (any, error) `v8:"clearContextInfoProc"`
  Options1 uint `v8:"options1"`
  Options2 uint `v8:"options2"`
  PluginState TECPluginStateRec `v8:"pluginState"`
}
