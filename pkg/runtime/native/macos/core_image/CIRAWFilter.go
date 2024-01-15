//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreImage.CIRAWFilter : CoreImage.CIFilter
*/

type CIRAWFilter struct {
  *CIFilter

}

func (r *CIRAWFilter) Exposure() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) MoireReductionAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SemanticSegmentationTeethMatte() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) ScaleFactor() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) NeutralTint() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) DetailAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetDetailAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) ExtendedDynamicRangeAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetNeutralChromaticity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SemanticSegmentationHairMatte() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SemanticSegmentationGlassesMatte() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsDraftModeEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetScaleFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetShadowBias() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsContrastSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) PortraitEffectsMatte() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) BoostAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) Orientation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetBaselineExposure() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetBoostAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetNeutralTemperature() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) FilterWithImageURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetDraftModeEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsColorNoiseReductionSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetSharpnessAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetLensCorrectionEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) LuminanceNoiseReductionAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetExtendedDynamicRangeAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) FilterWithCVPixelBuffer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) BaselineExposure() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsLensCorrectionSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsLensCorrectionEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) NeutralLocation() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetNeutralTint() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SemanticSegmentationSkinMatte() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SupportedDecoderVersions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) NativeSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetNeutralLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) PreviewImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetLuminanceNoiseReductionAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetLocalToneMapAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SupportedCameraModels() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) DecoderVersion() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetDecoderVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) BoostShadowAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetGamutMappingEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) LinearSpaceFilter() (*CIFilter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) ColorNoiseReductionAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetLinearSpaceFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SemanticSegmentationSkyMatte() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) ShadowBias() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetColorNoiseReductionAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsSharpnessSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) NeutralChromaticity() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsMoireReductionSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetMoireReductionAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsLocalToneMapSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) FilterWithImageData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetBoostShadowAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsLuminanceNoiseReductionSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetContrastAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsDetailSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) NeutralTemperature() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) LocalToneMapAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) Properties() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SetExposure() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) IsGamutMappingEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) SharpnessAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRAWFilter) ContrastAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

