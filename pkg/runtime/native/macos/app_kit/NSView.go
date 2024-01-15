//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/quartz_core"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSView : AppKit.NSResponder
*/

type NSView struct {
  *NSResponder
  *NSAnimatablePropertyContainer
  *NSUserInterfaceItemIdentification
  *NSDraggingDestination
  *NSAppearanceCustomization
  *NSAccessibilityElement
  *NSAccessibility
}

func (r *NSView) ScaleUnitSquareToSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) CacheDisplayInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetWantsRestingTouches() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) EnclosingScrollView() (*NSScrollView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) NeedsToDrawRect() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) ViewDidHide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ViewDidUnhide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ScrollPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetFrameSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) LockFocusIfCanDraw() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) LockFocusIfCanDrawInContext() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) RotateByAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) DrawRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetUserInterfaceLayoutDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) IsCompatibleWithResponsiveScrolling() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) AncestorSharedWithView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) GetRectsBeingDrawn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetBoundsOrigin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) FrameRotation() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) SetAlphaValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) PostsBoundsChangedNotifications() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) UpdateLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) OpaqueAncestor() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) SetPostsBoundsChangedNotifications() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) UserInterfaceLayoutDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) ViewDidChangeBackingProperties() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) CenterScanRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) HitTest() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) Layer() (*quartz_core.CALayer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertSizeToBacking() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertSizeFromBacking() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSView) SetPostsFrameChangedNotifications() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertRectFromLayer() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) ScrollRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) AllowsVibrancy() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) AddSubview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) RemoveFromSuperviewWithoutNeedingDisplay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetBoundsSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) DisplayRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) MouseDownCanMoveWindow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) Subviews() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) SetAutoresizingMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) LayerUsesCoreImageFilters() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertRectToBacking() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertSizeFromLayer() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSView) Mouse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) SetBackgroundFilters() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetFrameOrigin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) UnlockFocus() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) AdjustScroll() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) SetShadow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) InLiveResize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) PreservesContentDuringLiveResize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) RemoveToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) Bounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) AcceptsTouchEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) PrepareContentInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) Tag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) SetCanDrawSubviewsIntoLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetNeedsLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetLayerUsesCoreImageFilters() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ViewDidMoveToSuperview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ReplaceSubview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) DisplayRectIgnoringOpacity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetNeedsDisplay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) RemoveFromSuperview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) DisplayIfNeededInRectIgnoringOpacity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetBoundsRotation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetCanDrawConcurrently() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetContentFilters() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) RectPreservedDuringLiveResize() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertPointFromBacking() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSView) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) FrameCenterRotation() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) SetAcceptsTouchEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) AlphaValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) ViewWillMoveToWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) VisibleRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) BitmapImageRepForCachingDisplayInRect() (*NSBitmapImageRep, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) MenuForEvent() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) ViewDidChangeEffectiveAppearance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) WantsUpdateLayer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) SetToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) IsDescendantOf() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) ResizeWithOldSuperviewSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertPointFromLayer() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSView) LayerContentsRedrawPolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) PreparedContentRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) ViewDidMoveToWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ViewWillMoveToSuperview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ViewWithTag() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) PerformKeyEquivalent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) DidCloseMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) WantsLayer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) WillRemoveSubview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertPointToLayer() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSView) TranslateRectsNeedingDisplayInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertRectToLayer() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) ShouldDelayWindowOrderingForEvent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) WillOpenMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) IsRotatedOrScaledFromBase() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) DidAddSubview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ResizeSubviewsWithOldSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertSizeToLayer() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSView) SetLayerContentsPlacement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ClipsToBounds() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) SetLayerContentsRedrawPolicy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) Shadow() (*NSShadow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) DefaultMenu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) SetPreparedContentRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) DisplayIfNeededIgnoringOpacity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) DisplayIfNeededInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) Window() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) IsHiddenOrHasHiddenAncestor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) CanDrawConcurrently() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) CompositingFilter() (*CIFilter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) ContentFilters() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) SortSubviewsUsingFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) Display() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetSubviews() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) PostsFrameChangedNotifications() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) CanDraw() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) SetWantsLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ToolTip() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) RemoveAllToolTips() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) ViewDidEndLiveResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetAutoresizesSubviews() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) BackingAlignedRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) LayerContentsPlacement() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) SetCompositingFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) NeedsPanelToBecomeKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) InputContext() (*NSTextInputContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) DisplayIfNeeded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) MakeBackingLayer() (*quartz_core.CALayer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) WantsDefaultClipping() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) WantsRestingTouches() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) NeedsLayout() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) LockFocus() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) RectForSmartMagnificationAtPoint() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) IsOpaque() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) GetRectsExposedDuringLiveResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) PrepareForReuse() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) AutoresizesSubviews() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) SetFrameRotation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) IsFlipped() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) TranslateOriginToPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) AcceptsFirstMouse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) AddToolTipRect() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) ScrollRectToVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) NeedsDisplay() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) SetClipsToBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) SetFrameCenterRotation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) BackgroundFilters() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertRectFromBacking() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSView) ViewWillStartLiveResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) Superview() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) Autoscroll() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) BoundsRotation() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) FocusView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertPointToBacking() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSView) ViewWillDraw() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) Layout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) AutoresizingMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSView) IsRotatedFromBase() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) CanDrawSubviewsIntoLayer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSView) ConvertSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSView) SetNeedsDisplayInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSView) LayoutSubtreeIfNeeded() error {
  return fmt.Errorf("unimplemented")
}

