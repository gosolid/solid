//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntVariationHeader
*/

type SfntVariationHeader struct {
  Version int `v8:"version"`
  OffsetToData uint16 `v8:"offsetToData"`
  CountSizePairs uint16 `v8:"countSizePairs"`
  AxisCount uint16 `v8:"axisCount"`
  AxisSize uint16 `v8:"axisSize"`
  InstanceCount uint16 `v8:"instanceCount"`
  InstanceSize uint16 `v8:"instanceSize"`
  Axis [1]SfntVariationAxis `v8:"axis"`
  Instance [1]SfntInstance `v8:"instance"`
}
