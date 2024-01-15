//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
interface AppKit.NSColor : objc.NSObject
*/

type NSColor struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
  *NSPasteboardReading
  *NSPasteboardWriting
}

func (r *NSColor) ColorWithAlphaComponent() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) PlaceholderTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) PatternImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) KeyboardFocusIndicatorColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemYellowColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemGrayColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) CatalogNameComponent() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) BrightnessComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorUsingType() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemPinkColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) TertiarySystemFillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) QuinaryLabelColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) WindowBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SaturationComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithRed() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithCalibratedHue() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) UnemphasizedSelectedTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) CGColor() (*core_graphics.CGColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithDisplayP3Red() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) MagentaColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) UnemphasizedSelectedContentBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) LightGrayColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) HeaderTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) TextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemFillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SetIgnoresAlpha() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithCalibratedRed() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorUsingColorSpace() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) Set() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) AlphaComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithCalibratedWhite() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemGreenColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemMintColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) DarkGrayColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) WindowFrameTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SeparatorColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) AlternatingContentBackgroundColors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SelectedControlColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) GetComponents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorFromPasteboard() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorNameComponent() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) GreenComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithDeviceCyan() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) WhiteComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) YellowComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) HighlightWithLevel() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) GetRed() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) BlueColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithSystemEffect() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ControlBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithDeviceWhite() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorForControlTint() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ShadowWithLevel() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SelectedContentBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) BlendedColorWithFraction() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) OrangeColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) QuaternaryLabelColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithSRGBRed() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithGenericGamma22White() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) CyanComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) TextInsertionPointColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SecondarySystemFillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) MagentaComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) ControlAccentColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ShadowColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) PurpleColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) TextBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ScrubberTexturedBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) HighlightColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorNamed() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) GrayColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SelectedTextBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemOrangeColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithWhite() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) TertiaryLabelColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SelectedTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ControlColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ControlTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) BlackComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) CurrentControlTint() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) BlueComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) GetCyan() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) WriteToPasteboard() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) ClearColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) GreenColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) AlternateSelectedControlTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemBlueColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) LocalizedColorNameComponent() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) GetHue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithCGColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) WhiteColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) NumberOfComponents() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithPatternImage() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SetStroke() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) DisabledControlTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) LocalizedCatalogNameComponent() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithCatalogName() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) BlackColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemCyanColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) HueComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) IgnoresAlpha() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColor) SetFill() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) GridColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) QuaternarySystemFillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithColorSpace() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) FindHighlightColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemBrownColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithDeviceHue() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) DrawSwatchInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) UnemphasizedSelectedTextBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) QuinarySystemFillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) RedComponent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColor) CyanColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) LinkColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemIndigoColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SecondaryLabelColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SelectedMenuItemTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithHue() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithName() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) BrownColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemTealColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) RedColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) UnderPageBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemPurpleColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) ColorWithDeviceRed() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) YellowColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SystemRedColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) GetWhite() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColor) LabelColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColor) SelectedControlTextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

