//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSImageRep : objc.NSObject
*/

type NSImageRep struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSImageRep) CanInitWithData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepWithContentsOfFile() (*NSImageRep, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) SetOpaque() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) DrawAtPoint() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepClassForType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepClassForData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) PixelsWide() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) LayoutDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) RegisterImageRepClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) SetSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) Draw() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) BitsPerSample() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) SetPixelsWide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) CGImageForProposedRect() (*core_graphics.CGImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ColorSpaceName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) PixelsHigh() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) SetPixelsHigh() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageUnfilteredFileTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepsWithContentsOfURL() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepsWithPasteboard() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) HasAlpha() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) SetColorSpaceName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) SetBitsPerSample() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepClassForFileType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImagePasteboardTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepsWithContentsOfFile() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) Size() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) SetAlpha() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) IsOpaque() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) RegisteredImageRepClasses() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageUnfilteredTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) UnregisterImageRepClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) CanInitWithPasteboard() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepWithPasteboard() (*NSImageRep, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageFileTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepWithContentsOfURL() (*NSImageRep, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) SetLayoutDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageRep) DrawInRect() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageRepClassForPasteboardType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageRep) ImageUnfilteredPasteboardTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

