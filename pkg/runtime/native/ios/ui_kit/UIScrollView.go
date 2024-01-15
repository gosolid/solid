//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIScrollView : UIKit.UIView
*/

type UIScrollView struct {
  *UIView

}

func (r *UIScrollView) AllowsKeyboardScrolling() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IsDecelerating() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) CanCancelContentTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetBouncesZoom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetContentOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ShowsVerticalScrollIndicator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IsDragging() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IndicatorStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetDelaysContentTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) BouncesZoom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) TouchesShouldBegin() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetZoomScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetContentInsetAdjustmentBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ZoomToRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) RefreshControl() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetBounces() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ContentInset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetContentInset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ContentInsetAdjustmentBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) AlwaysBounceHorizontal() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetPagingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IsTracking() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) DelaysContentTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetMaximumZoomScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) AdjustedContentInset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetDirectionalLockEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) AlwaysBounceVertical() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) DirectionalPressGestureRecognizer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetIndicatorStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) PanGestureRecognizer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetAllowsKeyboardScrolling() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) FlashScrollIndicators() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) TouchesShouldCancelInContentView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) FrameLayoutGuide() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetIndexDisplayMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IsZoomBouncing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) AutomaticallyAdjustsScrollIndicatorInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetHorizontalScrollIndicatorInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IndexDisplayMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetMinimumZoomScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) KeyboardDismissMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) VerticalScrollIndicatorInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetScrollIndicatorInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) MaximumZoomScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ZoomScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ScrollsToTop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) AdjustedContentInsetDidChange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetContentSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) Bounces() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ContentLayoutGuide() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IsPagingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ContentSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetKeyboardDismissMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) PinchGestureRecognizer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ContentOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) HorizontalScrollIndicatorInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetCanCancelContentTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetAlwaysBounceHorizontal() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IsScrollEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ShowsHorizontalScrollIndicator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IsZooming() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetRefreshControl() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ScrollRectToVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) IsDirectionalLockEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetAlwaysBounceVertical() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetVerticalScrollIndicatorInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetScrollsToTop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetAutomaticallyAdjustsScrollIndicatorInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetScrollEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetShowsHorizontalScrollIndicator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetDecelerationRate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) MinimumZoomScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) ScrollIndicatorInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) SetShowsVerticalScrollIndicator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScrollView) DecelerationRate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

