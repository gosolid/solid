//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECBufferContextRec
*/

type TECBufferContextRec struct {
  TextInputBuffer *byte `v8:"textInputBuffer"`
  TextInputBufferEnd *byte `v8:"textInputBufferEnd"`
  TextOutputBuffer *byte `v8:"textOutputBuffer"`
  TextOutputBufferEnd *byte `v8:"textOutputBufferEnd"`
  EncodingInputBuffer *TextEncodingRun `v8:"encodingInputBuffer"`
  EncodingInputBufferEnd *TextEncodingRun `v8:"encodingInputBufferEnd"`
  EncodingOutputBuffer *TextEncodingRun `v8:"encodingOutputBuffer"`
  EncodingOutputBufferEnd *TextEncodingRun `v8:"encodingOutputBufferEnd"`
}
