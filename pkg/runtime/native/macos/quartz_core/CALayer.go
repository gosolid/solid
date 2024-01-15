//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
interface QuartzCore.CALayer : objc.NSObject
*/

type CALayer struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *CAMediaTiming
}

func (r *CALayer) ContentsScale() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) AddSublayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetGeometryFlipped() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMinificationFilterBias() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) Style() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) LayoutIfNeeded() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsFormat() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) DrawsAsynchronously() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetDrawsAsynchronously() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBackgroundFilters() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) ActionForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) IsOpaque() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) LayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetOpaque() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetCornerCurve() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) BorderWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowRadius() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) ZPosition() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsGravity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) RenderInContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) MaskedCorners() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) CornerCurve() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) LayoutSublayers() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) ResizeSublayersWithOldSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) InsertSublayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) CornerRadius() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowColor() (*core_graphics.CGColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAnchorPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetSublayers() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetActions() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMaskedCorners() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) AutoresizingMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) Mask() (*CALayer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMasksToBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMinificationFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) DefaultActionForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Sublayers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBorderWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowOffset() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetTransform() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) AllowsEdgeAntialiasing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) Opacity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) DisplayIfNeeded() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetOpacity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) AnchorPointZ() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBorderColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SublayerTransform() (CATransform3D, error) {
  return CATransform3D{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAllowsGroupOpacity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetFilters() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) RasterizationScale() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) Display() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetZPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) Contents() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetCornerRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowOpacity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetNeedsDisplay() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) DrawInContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) MasksToBounds() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetRasterizationScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowPath() (*core_graphics.CGPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) RemoveAllAnimations() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) NeedsDisplayOnBoundsChange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetEdgeAntialiasingMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetLayoutManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) Position() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) WantsExtendedDynamicRangeContent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) Transform() (CATransform3D, error) {
  return CATransform3D{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetWantsExtendedDynamicRangeContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) CompositingFilter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) InitWithLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ResizeWithOldSuperlayerSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) MinificationFilter() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetNeedsLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) IsGeometryFlipped() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowOpacity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) PreferredFrameSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetDoubleSided() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) Superlayer() (*CALayer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) MinificationFilterBias() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetCompositingFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ConvertPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) Filters() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetNeedsDisplayOnBoundsChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) AllowsGroupOpacity() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) ShouldRasterize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) Layer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) DefaultValueForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) CornerCurveExpansionFactor() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) AddAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) IsDoubleSided() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsGravity() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ConvertTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) HitTest() (*CALayer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AnimationForKey() (*CAAnimation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AnchorPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetSublayerTransform() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) MagnificationFilter() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) EdgeAntialiasingMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CALayer) BackgroundColor() (*core_graphics.CGColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AffineTransform() (core_foundation.CGAffineTransform, error) {
  return core_foundation.CGAffineTransform{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) ReplaceSublayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) BackgroundFilters() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShouldRasterize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetNeedsDisplayInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAllowsEdgeAntialiasing() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) BorderColor() (*core_graphics.CGColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Actions() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ModelLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) RemoveFromSuperlayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) AnimationKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Bounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsCenter() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAffineTransform() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) ConvertRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAnchorPointZ() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMagnificationFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAutoresizingMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayer) PresentationLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) RemoveAnimationForKey() error {
  return fmt.Errorf("unimplemented")
}

