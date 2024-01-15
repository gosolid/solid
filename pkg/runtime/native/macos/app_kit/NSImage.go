//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSImage : objc.NSObject
*/

type NSImage struct {
  *objc.NSObject

}

func (r *NSImage) CGImageForProposedRect() (*core_graphics.CGImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetMatchesOnMultipleResolution() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) SetCacheMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) SetAlignmentRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) SymbolConfiguration() (*NSImageSymbolConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) ImageWithSystemSymbolName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) Recache() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) InitWithCGImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetCapInsets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) InitByReferencingFile() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) DrawAtPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) AddRepresentations() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) UsesEPSOnResolutionMismatch() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) MatchesOnMultipleResolution() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) ImageUnfilteredTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) Locale() (*foundation.NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) AddRepresentation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) InitWithDataIgnoringOrientation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) Name() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) RemoveRepresentation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) ImageTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) AccessibilityDescription() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) TIFFRepresentationUsingCompression() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) RecommendedLayerContentsScale() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImage) IsTemplate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) DrawRepresentation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetResizingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) ImageWithSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) DrawInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) HitTestRect() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) PrefersColorMatch() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetAccessibilityDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) ImageNamed() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) InitWithPasteboard() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) MatchesOnlyOnBestFittingAxis() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) ResizingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImage) ImageWithSymbolConfiguration() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetPrefersColorMatch() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) CacheMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImage) AlignmentRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetTemplate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) InitWithSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) InitWithContentsOfFile() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) Size() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSImage) ImageWithSymbolName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) LayerContentsForContentsScale() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) SetMatchesOnlyOnBestFittingAxis() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) TIFFRepresentation() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) Representations() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) IsValid() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) InitByReferencingURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) SetUsesEPSOnResolutionMismatch() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImage) CapInsets() (foundation.NSEdgeInsets, error) {
  return foundation.NSEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSImage) CanInitWithPasteboard() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImage) BestRepresentationForRect() (*NSImageRep, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImage) ImageWithLocale() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

