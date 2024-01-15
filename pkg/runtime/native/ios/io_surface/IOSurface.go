//js:package native/ios/io-surface
package io_surface

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface IOSurface.IOSurface : objc.NSObject
*/

type IOSurface struct {
  *objc.NSObject

}

func (r *IOSurface) BaseAddress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) AllAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) IncrementUseCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) PlaneCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) LocalUseCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) LockWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) PixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) InitWithProperties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) Seed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) RemoveAllAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) AllocationSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) BytesPerRow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) ElementWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) IsInUse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) ElementHeightOfPlaneAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) AttachmentForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) AllowsPixelSizeCasting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) WidthOfPlaneAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) RemoveAttachmentForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) BaseAddressOfPlaneAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) BytesPerElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) UnlockWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) HeightOfPlaneAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) SetPurgeable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) ElementHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) BytesPerElementOfPlaneAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) SetAllAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) SetAttachment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) DecrementUseCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) Width() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) Height() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) BytesPerRowOfPlaneAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) ElementWidthOfPlaneAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

