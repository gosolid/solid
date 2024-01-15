//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIViewController : UIKit.UIResponder
*/

type UIViewController struct {
  *UIResponder

}

func (r *UIViewController) ViewDidDisappear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetProvidesPresentationContextTransitionStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) IsBeingDismissed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) WantsFullScreenLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewIsAppearing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PerformSegueWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) IsViewLoaded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ParentViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) FocusGroupIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ModalTransitionStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PreferredUserInterfaceStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) OverrideUserInterfaceStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) LoadView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) IsBeingPresented() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetContentScrollView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ChildViewControllerContainingSegueSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) NibBundle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ModalPresentationStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PreferredStatusBarStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) InitWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetModalPresentationCapturesStatusBarAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PreferredStatusBarUpdateAnimation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ContentScrollViewForEdge() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ShouldPerformSegueWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) UnwindForSegue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SegueForUnwindingToViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewDidLayoutSubviews() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewIfLoaded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) NibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PreferredContentSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) LoadViewIfNeeded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PresentModalViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetFocusGroupIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ExtendedLayoutIncludesOpaqueBars() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) AllowedChildViewControllersForUnwindingFromSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PrepareForSegue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) DidReceiveMemoryWarning() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) View() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ModalViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) RestoresFocusAfterTransition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PrefersStatusBarHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetNeedsStatusBarAppearanceUpdate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) DismissModalViewControllerAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetRestoresFocusAfterTransition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetInteractionActivityTrackingBaseName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewWillLayoutSubviews() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewWillAppear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) TargetViewControllerForAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ModalPresentationCapturesStatusBarAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetExtendedLayoutIncludesOpaqueBars() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) CanPerformUnwindSegueAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) IsMovingFromParentViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) Storyboard() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PresentViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ShowDetailViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ProvidesPresentationContextTransitionStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) DisablesAutomaticKeyboardDismissal() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) AutomaticallyAdjustsScrollViewInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetAutomaticallyAdjustsScrollViewInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewDidUnload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewDidAppear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PresentedViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetDefinesPresentationContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetWantsFullScreenLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewDidLoad() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewControllerForUnwindSegueAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewWillDisappear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) DismissViewControllerAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetNeedsUserInterfaceAppearanceUpdate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) InteractionActivityTrackingBaseName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetModalTransitionStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetModalPresentationStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ViewWillUnload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetPreferredContentSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetOverrideUserInterfaceStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) SetEdgesForExtendedLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) PresentingViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) DefinesPresentationContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) IsMovingToParentViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) EdgesForExtendedLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewController) ShowViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

