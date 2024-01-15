//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface QuartzCore.CALayer : objc.NSObject
*/

type CALayer struct {
  *objc.NSObject

}

func (r *CALayer) IsGeometryFlipped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsGravity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AllowsEdgeAntialiasing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) MasksToBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetWantsExtendedDynamicRangeContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetNeedsDisplayOnBoundsChange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ConvertRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) LayoutIfNeeded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) WantsExtendedDynamicRangeContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Filters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetActions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) BorderColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) DisplayIfNeeded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetDoubleSided() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMinificationFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) MagnificationFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) DefaultValueForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ReplaceSublayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) NeedsLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AddSublayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) DrawInContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAnchorPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) MinificationFilterBias() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetRasterizationScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) InitWithLayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) InsertSublayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AnimationKeys() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShouldArchiveValueForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) RemoveAnimationForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) IsOpaque() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMaskedCorners() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShouldRasterize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) RemoveAllAnimations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) PresentationLayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) DefaultActionForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AnchorPointZ() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AddAnimation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SublayerTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetDrawsAsynchronously() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) BackgroundFilters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) CornerCurve() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBorderWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsAreFlipped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ConvertPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetZPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShouldRasterize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetNeedsDisplay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) RenderInContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Sublayers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetSublayers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Position() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMagnificationFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetCornerRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) NeedsDisplay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetNeedsLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMasksToBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Style() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ModelLayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) NeedsDisplayForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) IsDoubleSided() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMinificationFilterBias() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetCornerCurve() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ConvertTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ActionForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Frame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Superlayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) RasterizationScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAffineTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ContainsPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ZPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AllowsGroupOpacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetEdgeAntialiasingMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) CornerRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBorderColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowOpacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Layer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) MinificationFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAllowsGroupOpacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ShadowColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Actions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) RemoveFromSuperlayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAnchorPointZ() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Mask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Opacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) CornerCurveExpansionFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetSublayerTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Contents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) HitTest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AnchorPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetOpaque() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) NeedsDisplayOnBoundsChange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) CompositingFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) PreferredFrameSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Transform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) IsHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsGravity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsCenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) BorderWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetShadowOpacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBackgroundFilters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) LayoutSublayers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AnimationForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) EdgeAntialiasingMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetFilters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) Display() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) DrawsAsynchronously() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetCompositingFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetOpacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetNeedsDisplayInRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetGeometryFlipped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) MaskedCorners() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) AffineTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) ContentsScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetContentsCenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayer) SetAllowsEdgeAntialiasing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

