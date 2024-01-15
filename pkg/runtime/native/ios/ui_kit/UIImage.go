//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIImage : objc.NSObject
*/

type UIImage struct {
  *objc.NSObject

}

func (r *UIImage) SystemImageNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) CGImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithAlignmentRectInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageRestrictedToStandardDynamicRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageRendererFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) AnimatedResizableImageNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithoutBaseline() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) PrepareForDisplayWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) TraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) SymbolConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) IsHighDynamicRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ResizableImageWithCapInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageByPreparingThumbnailOfSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) DrawAsPatternInRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) PrepareThumbnailOfSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) HasBaseline() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) DrawAtPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageFlippedForRightToLeftLayoutDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithHorizontallyFlippedOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithBaselineOffsetFromBottom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) RenderingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) InitWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) AnimatedImageWithImages() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) CapInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithContentsOfFile() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithCGImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) AlignmentRectInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) InitWithContentsOfFile() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) CIImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) FlipsForRightToLeftLayoutDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) InitWithCIImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageByApplyingSymbolConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageByPreparingForDisplay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ResizingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageAsset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) AnimatedImageNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) Size() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) DrawInRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) Images() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) BaselineOffsetFromBottom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithCIImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) InitWithCGImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) IsSymbolImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) Duration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) ImageWithRenderingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImage) Scale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

