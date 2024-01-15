//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIImageReaderConfiguration : objc.NSObject
*/

type UIImageReaderConfiguration struct {
  *objc.NSObject

}

func (r *UIImageReaderConfiguration) SetPreferredThumbnailSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReaderConfiguration) PixelsPerInch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReaderConfiguration) SetPixelsPerInch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReaderConfiguration) PrefersHighDynamicRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReaderConfiguration) SetPrefersHighDynamicRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReaderConfiguration) PreparesImagesForDisplay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReaderConfiguration) SetPreparesImagesForDisplay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReaderConfiguration) PreferredThumbnailSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

