//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntDescriptorHeader
*/

type SfntDescriptorHeader struct {
  Version int `v8:"version"`
  DescriptorCount int `v8:"descriptorCount"`
  Descriptor [1]SfntFontDescriptor `v8:"descriptor"`
}
