//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIApplication : UIKit.UIResponder
*/

type UIApplication struct {
  *UIResponder

}

func (r *UIApplication) IsStatusBarHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) ConnectedScenes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) OpenSessions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) UserInterfaceLayoutDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) EndBackgroundTask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) KeyWindow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) ApplicationSupportsShakeToEdit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) IsProtectedDataAvailable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SetApplicationIconBadgeNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SetApplicationSupportsShakeToEdit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SupportedInterfaceOrientationsForWindow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) BeginBackgroundTaskWithExpirationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) ActivateSceneSessionForRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) RequestSceneSessionActivation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) StatusBarStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) StatusBarOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SetMinimumBackgroundFetchInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) IsIgnoringInteractionEvents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) IsNetworkActivityIndicatorVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SetNetworkActivityIndicatorVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) StatusBarFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) BackgroundRefreshStatus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SendAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SetIdleTimerDisabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) StatusBarOrientationAnimationDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SendEvent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) RequestSceneSessionDestruction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) IsIdleTimerDisabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) Windows() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) BeginIgnoringInteractionEvents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) EndIgnoringInteractionEvents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) OpenURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) CanOpenURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) PreferredContentSizeCategory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) BeginBackgroundTaskWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) RequestSceneSessionRefresh() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) ApplicationState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SupportsMultipleScenes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SharedApplication() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) ApplicationIconBadgeNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplication) BackgroundTimeRemaining() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

