//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_video"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface CoreImage.CIImage : objc.NSObject
*/

type CIImage struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CIImage) ImageTransformForOrientation() (core_foundation.CGAffineTransform, error) {
  return core_foundation.CGAffineTransform{}, fmt.Errorf("unimplemented")
}

func (r *CIImage) WhiteImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCGLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingTransform() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithIOSurface() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithColor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByPremultiplyingAlpha() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) BlueImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCGImageSource() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithContentsOfURL() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) CyanImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) YellowImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingOrientation() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) GrayImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCVImageBuffer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCGImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ClearImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageBySettingAlphaOneInExtent() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) RedImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByCompositingOverImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByInsertingIntermediate() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) Url() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCGLayer() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCGImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithBitmapData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageBySettingProperties() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageBySamplingNearest() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) MagentaImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithTexture() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithMTLTexture() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByCroppingToRect() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByColorMatchingColorSpaceToWorkingSpace() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByColorMatchingWorkingSpaceToColorSpace() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithColor() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCGImageSource() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCVPixelBuffer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingGaussianBlurWithSigma() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) Definition() (*CIFilterShape, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ColorSpace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) CGImage() (*core_graphics.CGImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithTexture() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithIOSurface() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageBySamplingLinear() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) RegionOfInterestForImage() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageTransformForCGOrientation() (core_foundation.CGAffineTransform, error) {
  return core_foundation.CGAffineTransform{}, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByUnpremultiplyingAlpha() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByClampingToExtent() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByClampingToRect() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) BlackImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) PixelBuffer() (*core_video.CVBuffer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCVPixelBuffer() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingCGOrientation() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) GreenImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCVImageBuffer() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithMTLTexture() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) EmptyImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingFilter() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithBitmapData() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithData() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) Properties() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

