//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface QuartzCore.CAEDRMetadata : objc.NSObject
*/

type CAEDRMetadata struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CAEDRMetadata) HLGMetadata() (*CAEDRMetadata, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) IsAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) HDR10MetadataWithDisplayInfo() (*CAEDRMetadata, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) HDR10MetadataWithMinLuminance() (*CAEDRMetadata, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEDRMetadata) HLGMetadataWithAmbientViewingEnvironment() (*CAEDRMetadata, error) {
  return nil, fmt.Errorf("unimplemented")
}

