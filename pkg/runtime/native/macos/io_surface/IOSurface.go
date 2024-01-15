//js:package native/macos/io-surface
package io_surface

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface IOSurface.IOSurface : objc.NSObject
*/

type IOSurface struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *IOSurface) ElementWidth() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) SetAllAttachments() error {
  return fmt.Errorf("unimplemented")
}

func (r *IOSurface) Width() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) RemoveAttachmentForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *IOSurface) AllowsPixelSizeCasting() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *IOSurface) ElementHeightOfPlaneAtIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) BaseAddressOfPlaneAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *IOSurface) BytesPerElementOfPlaneAtIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) RemoveAllAttachments() error {
  return fmt.Errorf("unimplemented")
}

func (r *IOSurface) PixelFormat() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) UnlockWithOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) HeightOfPlaneAtIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) IncrementUseCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *IOSurface) BytesPerRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) LocalUseCount() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) WidthOfPlaneAtIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) SetAttachment() error {
  return fmt.Errorf("unimplemented")
}

func (r *IOSurface) BaseAddress() error {
  return fmt.Errorf("unimplemented")
}

func (r *IOSurface) BytesPerRowOfPlaneAtIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) AllAttachments() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) PlaneCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) InitWithProperties() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) Seed() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) SetPurgeable() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) Height() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) BytesPerElement() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) ElementHeight() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) ElementWidthOfPlaneAtIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) AttachmentForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *IOSurface) AllocationSize() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) IsInUse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *IOSurface) LockWithOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *IOSurface) DecrementUseCount() error {
  return fmt.Errorf("unimplemented")
}

