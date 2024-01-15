//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWebView : UIKit.UIView
*/

type UIWebView struct {
  *UIView

}

func (r *UIWebView) StringByEvaluatingJavaScriptFromString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetMediaPlaybackRequiresUserAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetMediaPlaybackAllowsAirPlay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SuppressesIncrementalRendering() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) GapBetweenPages() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) Reload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) StopLoading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) ScrollView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetGapBetweenPages() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) PaginationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetPaginationBreakingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetPageLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) AllowsPictureInPictureMediaPlayback() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) GoForward() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetDetectsPhoneNumbers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) GoBack() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetAllowsInlineMediaPlayback() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) MediaPlaybackRequiresUserAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetAllowsPictureInPictureMediaPlayback() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) LoadRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) LoadHTMLString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) DataDetectorTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetDataDetectorTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetPaginationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetAllowsLinkPreview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) LoadData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) Request() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetSuppressesIncrementalRendering() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) KeyboardDisplayRequiresUserAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetKeyboardDisplayRequiresUserAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) PaginationBreakingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) SetScalesPageToFit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) AllowsInlineMediaPlayback() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) IsLoading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) ScalesPageToFit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) PageCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) CanGoBack() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) CanGoForward() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) PageLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) AllowsLinkPreview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) DetectsPhoneNumbers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWebView) MediaPlaybackAllowsAirPlay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

