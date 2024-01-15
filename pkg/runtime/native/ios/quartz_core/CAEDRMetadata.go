//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface QuartzCore.CAEDRMetadata : objc.NSObject
*/

type CAEDRMetadata struct {
  *objc.NSObject

}

func (r *CAEDRMetadata) HLGMetadata() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) IsAvailable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) HDR10MetadataWithDisplayInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) HDR10MetadataWithMinLuminance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) HLGMetadataWithAmbientViewingEnvironment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

