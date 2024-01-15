//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSAccessibility
*/

type NSAccessibility struct {

}

func (r *NSAccessibility) SetAccessibilityFocused() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityActivationPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySelectedChildren() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityWarningValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySharedFocusElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityLayoutPointForScreenPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPlaceholderValue() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityFilename() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMinValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityMarkerGroupUIElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityScreenSizeForLayoutSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySplitters() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityWindows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMarkerGroupUIElement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySelectedColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityStringForRange() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformCancel() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityVisibleChildren() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityProtectedContent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityCustomActions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformConfirm() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityApplicationFocusedUIElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityCustomActions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySelectedText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySelectedTextRanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityCancelButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityExtrasMenuBar() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityLabelValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityColumnCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityApplicationFocusedUIElement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityMain() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityIncrementButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityColumnCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityFocused() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityUserInputLabels() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityCriticalValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityInsertionPointLineNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityScreenPointForLayoutPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityStyleRangeForIndex() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityClearButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMenuBar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityVerticalUnits() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityMarkerTypeDescription() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRTFForRange() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformRaise() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySelected() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityTitleUIElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySearchButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMarkerTypeDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityLineForIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityMaxValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMarkerUIElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityTopLevelUIElement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityGrowArea() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityUnits() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityFullScreenButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityDisclosureLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityAllowedValues() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityInsertionPointLineNumber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySortDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRowHeaderUIElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityVerticalUnitDescription() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMarkerValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityUnits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySplitters() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityVerticalScrollBar() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityExpanded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMainWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityColumnTitles() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySelectedTextRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRangeForLine() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySubrole() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySharedFocusElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityVisibleCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityDisclosedByRow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityDisclosedRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityUnitDescription() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityDecrementButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityCellForColumn() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMaxValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityOrderedByRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityDisclosureLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityMinimized() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformPick() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityValueDescription() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityZoomButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPreviousContents() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityCustomRotors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityFocusedWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMain() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityHorizontalScrollBar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityChildren() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityProtectedContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityHandles() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityRoleDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityCustomRotors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityRowIndexRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityVisibleCharacterRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityDisclosedRows() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityHelp() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityRequired() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityVerticalUnitDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityNumberOfCharacters() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityModal() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityColumnTitles() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySharedTextUIElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityTopLevelUIElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityValueDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityOverflowButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityShownMenu() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySelectedText() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityServesAsTitleForUIElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityHorizontalUnitDescription() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityHorizontalUnitDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityRulerMarkerType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySearchMenu() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityChildrenInNavigationOrder() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityHorizontalUnits() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityDisclosed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformDecrement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformDelete() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformPress() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityVisibleChildren() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityIncrementButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityParent() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityServesAsTitleForUIElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityFocusedWindow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityZoomButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityPlaceholderValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRoleDescription() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityTabs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityToolbarButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRowCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityRowCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySubrole() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityShownMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityWindows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityUnitDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformShowMenu() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityHelp() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityColumnHeaderUIElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityDocument() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityChildren() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityHeader() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityVisibleCells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityTabs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityHeader() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityVisibleColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySelectedCells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityFrame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityNextContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityDisclosed() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityVisibleColumns() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformIncrement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityFrontmost() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityCriticalValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySelectedRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityFrameForRange() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityLinkedUIElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityEdited() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityCloseButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRangeForPosition() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityNextContents() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityLabelUIElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilitySelected() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityAttributedUserInputLabels() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityColumnIndexRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityToolbarButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityGrowArea() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMinimizeButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityExtrasMenuBar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityWarningValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySelectedTextRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityDefaultButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityOrientation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityChildrenInNavigationOrder() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityMarkerValues() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityLayoutSizeForScreenSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityAttributedStringForRange() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilitySelectorAllowed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityTitleUIElement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityDisclosedByRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityRowHeaderUIElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityDecrementButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityColumns() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityCloseButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRangeForIndex() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityElement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityAlternateUIVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityElement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityEdited() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityVisibleRows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityVisibleRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityMainWindow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySelectedColumns() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityProxy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityPreviousContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySelectedChildren() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityExpanded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityClearButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityLabelValue() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySelectedRows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformShowAlternateUI() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityPerformShowDefaultUI() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityVerticalUnits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityAllowedValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityAttributedUserInputLabels() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityMarkerUIElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySharedCharacterRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityProxy() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityHorizontalUnits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityOverflowButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityParent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySearchButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityMenuBar() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySharedTextUIElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySelectedTextRanges() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRole() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityHorizontalScrollBar() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySelectedCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityColumnIndexRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityRequired() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityFrontmost() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityFilename() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityUserInputLabels() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityMinimized() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityWindow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRowIndexRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityModal() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityRulerMarkerType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityLinkedUIElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityAlternateUIVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) IsAccessibilityOrderedByRow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityDefaultButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityMinimizeButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityMinValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityRole() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityVerticalScrollBar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityColumnHeaderUIElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityFullScreenButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityHandles() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityLabelUIElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySortDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilitySearchMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityContents() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityVisibleCharacterRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityCancelButton() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilitySharedCharacterRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) SetAccessibilityNumberOfCharacters() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibility) AccessibilityActivationPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

