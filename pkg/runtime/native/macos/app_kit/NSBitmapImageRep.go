//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
interface AppKit.NSBitmapImageRep : AppKit.NSImageRep
*/

type NSBitmapImageRep struct {
  *NSImageRep
  *foundation.NSSecureCoding
}

func (r *NSBitmapImageRep) BytesPerPlane() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) InitWithCGImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) GetBitmapDataPlanes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) GetPixel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) InitWithBitmapDataPlanes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) InitWithCIImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) LocalizedNameForTIFFCompressionType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) IncrementalLoadFromData() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) BitsPerPixel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) ImageRepWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) ColorizeByMappingGray() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) BitmapData() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) ImageRepsWithData() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) BitmapImageRepByConvertingToColorSpace() (*NSBitmapImageRep, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) CGImage() (*core_graphics.CGImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) TIFFRepresentationUsingCompression() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) BytesPerRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) BitmapFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) ColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) NumberOfPlanes() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) InitWithFocusedViewRect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) GetTIFFCompressionTypes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) BitmapImageRepByRetaggingWithColorSpace() (*NSBitmapImageRep, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) SamplesPerPixel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) IsPlanar() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) SetCompression() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) TIFFRepresentationOfImageRepsInArray() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) CanBeCompressedUsing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) SetPixel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) GetCompression() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) InitForIncrementalLoad() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) ColorAtX() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBitmapImageRep) TIFFRepresentation() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

