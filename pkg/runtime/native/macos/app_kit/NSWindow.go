//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSWindow : AppKit.NSResponder
*/

type NSWindow struct {
  *NSResponder
  *NSAnimatablePropertyContainer
  *NSMenuItemValidation
  *NSUserInterfaceValidations
  *NSUserInterfaceItemIdentification
  *NSAppearanceCustomization
  *NSAccessibilityElement
  *NSAccessibility
}

func (r *NSWindow) OrderWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ConvertRectFromScreen() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetFrameAutosaveName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) StandardWindowButton() (*NSButton, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ViewsNeedDisplay() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) MinFullScreenContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Display() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) TitlebarAccessoryViewControllers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) RepresentedURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsKeyWindow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetAppearanceSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) InitWithContentRect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Miniaturize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) InLiveResize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsMiniaturized() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetCanHide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetContentMaxSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) WindowController() (*NSWindowController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ChildWindows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) BeginSheet() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetRepresentedURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) BackingType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Sheets() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) MaxFullScreenContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) AnimationResizeTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) AutorecalculatesContentBorderThicknessForEdge() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetFrameFromString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ParentWindow() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) WindowTitlebarLayoutDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) RemoveTitlebarAccessoryViewControllerAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) Deminiaturize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetSubtitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentResizeIncrements() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetAlphaValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentBorderThicknessForEdge() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) OrderFront() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetDynamicDepthLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) EnableKeyEquivalentForDefaultButtonCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetTitlebarSeparatorStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ConstrainFrameRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) MergeAllWindows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) UserTabbingPreference() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Toolbar() (*NSToolbar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) TabbingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsMainWindow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) MakeMainWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ConvertPointFromScreen() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) TabGroup() (*NSWindowTabGroup, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ConvertRectToBacking() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) RemoveChildWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) TransferWindowSharingToWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentMaxSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) HasActiveWindowSharingSession() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetMaxSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SheetParent() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) TabbedWindows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) AddTitlebarAccessoryViewController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) InsertTitlebarAccessoryViewController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) WindowWithContentViewController() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetHasShadow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetAllowsConcurrentViewDrawing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetContentMinSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) DataWithPDFInsideRect() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SelectKeyViewPrecedingView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) Subtitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetHidesOnDeactivate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentRectForFrameRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) BecomeKeyWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) AddChildWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) TitlebarAppearsTransparent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ResizeIncrements() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsZoomed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) PreventsApplicationTerminationWhenModal() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) DisplaysWhenScreenProfileChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetShowsToolbarButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) OrderOut() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetStyleMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) AspectRatio() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) AddTabbedWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) MiniwindowImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ConvertPointFromBacking() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetToolbarStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) StringWithSavedFrame() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ToggleTabBar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsMovable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsOnActiveSpace() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) DefaultButtonCell() (*NSButtonCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) BackingAlignedRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) WorksWhenModal() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetAllowsToolTipsWhenApplicationIsInactive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) PerformMiniaturize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) WindowNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) AlphaValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetFrameTopLeftPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) MakeFirstResponder() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentLayoutRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Level() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetCollectionBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) FrameAutosaveName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) PerformClose() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ToggleToolbarShown() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetCanBecomeVisibleWithoutLogin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) DeviceDescription() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ResignMainWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ConvertRectToScreen() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) DataWithEPSInsideRect() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentLayoutGuide() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) CanBecomeVisibleWithoutLogin() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetDepthLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) AttachedSheet() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetParentWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) Tab() (*NSWindowTab, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ConvertRectFromBacking() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) PerformZoom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SelectKeyViewFollowingView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetContentView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) DeepestScreen() (*NSScreen, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetTitleWithRepresentedFilename() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) Update() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ValidRequestorForSendType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) OrderBack() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) DisableScreenUpdatesUntilFlush() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetFrameUsingName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SelectNextTab() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetReleasedWhenClosed() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) HasDynamicDepthLimit() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) InitialFirstResponder() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) RepresentedFilename() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsExcludedFromWindowsMenu() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ResizeFlags() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) TryToPerform() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Center() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) RunToolbarCustomizationPalette() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsReleasedWhenClosed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) HidesOnDeactivate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) MiniwindowTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) MinFrameWidthWithTitle() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) FirstResponder() (*NSResponder, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetWindowController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetColorSpace() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) KeyViewSelectionDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) CanRepresentDisplayGamut() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetTitleVisibility() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetExcludedFromWindowsMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetContentResizeIncrements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsMovableByWindowBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentViewController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetContentViewController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) EndSheet() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) CollectionBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) AppearanceSource() (*objc.NSObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetInitialFirstResponder() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) InvalidateShadow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ToolbarStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) CanHide() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) AllowsConcurrentViewDrawing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ToggleTabOverview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetMovable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) FrameRectForContentRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ResignKeyWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentAspectRatio() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetPreventsApplicationTerminationWhenModal() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) BackingScaleFactor() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetAutorecalculatesContentBorderThickness() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) HasShadow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetContentBorderThickness() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetTitlebarAppearsTransparent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetResizeIncrements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetViewsNeedDisplay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) AutorecalculatesKeyViewLoop() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) BecomeMainWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) WindowNumbersWithOptions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SharingType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetFrameOrigin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsDocumentEdited() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) ConvertPointToScreen() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetTitlebarAccessoryViewControllers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) StyleMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) MakeKeyAndOrderFront() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetContentAspectRatio() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetPreservesContentDuringLiveResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) BeginCriticalSheet() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetMovableByWindowBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetMiniwindowTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) DockTile() (*NSDockTile, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) AllowsToolTipsWhenApplicationIsInactive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetBackingType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetMinSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetMinFullScreenContentSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) EndEditingFor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) Close() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SelectPreviousKeyView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetDisplaysWhenScreenProfileChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) OcclusionState() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) RemoveFrameUsingName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) DefaultDepthLimit() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) TitleVisibility() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) MaxSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetAutorecalculatesKeyViewLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) TabbingIdentifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SaveFrameUsingName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetToolbar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ShowsToolbarButton() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) OrderFrontRegardless() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) PerformWindowDragWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) CanBecomeMainWindow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Screen() (*NSScreen, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetTabbingIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) MakeKeyWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) FieldEditor() (*NSText, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetOpaque() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) MinSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetAllowsAutomaticWindowTabbing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ConvertPointToBacking() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SelectNextKeyView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) DisableKeyEquivalentForDefaultButtonCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsSheet() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Zoom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) WindowNumberAtPoint() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetMaxFullScreenContentSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetContentSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) DisplayIfNeeded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SelectPreviousTab() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetRepresentedFilename() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetAnimationBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetDefaultButtonCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) AllowsAutomaticWindowTabbing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetDocumentEdited() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ContentMinSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) RecalculateKeyViewLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetMiniwindowImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) DepthLimit() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetTabbingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) CascadeTopLeftFromPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetAspectRatio() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) TitlebarSeparatorStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWindow) Print() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) ToggleFullScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) MoveTabToNewWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) PreservesContentDuringLiveResize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) CanBecomeKeyWindow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) IsOpaque() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindow) SetSharingType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindow) AnimationBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

