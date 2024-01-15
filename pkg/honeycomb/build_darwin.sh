#!/bin/bash

set -e

cd "$(dirname "$(realpath "$0")")"

export GOOS=${GOOS:-ios}
export GOARCH=${GOARCH:-amd64}
export CGO_ENABLED=1
export SDK=${SDK:-iphonesimulator}
export CGO_CFLAGS="-fembed-bitcode"


export SDK_PATH=$(xcrun --sdk "$SDK" --show-sdk-path)

if [ "$GOARCH" = "amd64" ]; then
    CARCH="x86_64"
elif [ "$GOARCH" = "arm64" ]; then
    CARCH="arm64"
fi

if [ "$SDK" = "iphoneos" ]; then
  export MIN_VERSION=15
  export TARGET="$CARCH-apple-ios$MIN_VERSION"
  export APPICON="AppIcon"
elif [ "$SDK" = "macosx" ]; then
  export MIN_VERSION=14
  export TARGET="$CARCH-apple-macosx$MIN_VERSION"
elif [ "$SDK" = "iphonesimulator" ]; then
  export MIN_VERSION=15
  export TARGET="$CARCH-apple-ios$MIN_VERSION-simulator"
  export APPICON="AppIcon"
fi

export CGO_LDFLAGS="-target ${TARGET} -isysroot \"${SDK_PATH}\""
export CC="$(pwd)/clangwrap.sh"
export CXX="$(pwd)/clangwrap.sh"
export CLANG=$(xcrun --sdk "$SDK" --find clang)

export BUNDLE_DIR="$(realpath "../../")/build/${SDK}/Honeycomb.app"

rm -Rf "${BUNDLE_DIR}/"
mkdir -p "${BUNDLE_DIR}/"


if [ "$SDK" = "macosx" ]; then
  export BUNDLE_DIR="$(realpath "../../")/build/${SDK}/Honeycomb.app/Contents"
  mkdir -p "${BUNDLE_DIR}/"
fi

cp info_${GOOS}.plist "${BUNDLE_DIR}/Info.plist"
if ! [ "$SDK" = "macosx" ]; then
  cp embedded.mobileprovision "${BUNDLE_DIR}/embedded.mobileprovision"
else
  cp app.entitlements "${BUNDLE_DIR}/app.entitlements"
fi
echo "APPL????" >"${BUNDLE_DIR}/PkgInfo"
if ! [ "$SDK" = "macosx" ]; then
  mkdir -p "${BUNDLE_DIR}/assets"
  go build -a -o "${BUNDLE_DIR}/Honeycomb" -tags ${NODE_ENV:-development} .
  xcrun actool Assets.xcassets --compile "${BUNDLE_DIR}/" --app-icon ${APPICON} --platform ${SDK} --output-partial-info-plist "${BUNDLE_DIR}/assets.plist" --minimum-deployment-target 8.0
  /usr/libexec/PlistBuddy -x -c "Merge \"${BUNDLE_DIR}/assets.plist\"" "${BUNDLE_DIR}/Info.plist"
  rm "${BUNDLE_DIR}/assets.plist"
else
  mkdir -p "${BUNDLE_DIR}/MacOS" "${BUNDLE_DIR}/Resources" "${BUNDLE_DIR}/Frameworks"
  go build -a -o "${BUNDLE_DIR}/MacOS/Honeycomb" -tags ${NODE_ENV:-development} .
  cp -r "/usr/local/lib/chromium/Chromium Framework.framework" "${BUNDLE_DIR}/Frameworks/Chromium Framework.framework"
  cp /usr/local/lib/chromium/*.dylib "${BUNDLE_DIR}/Frameworks/Chromium Framework.framework"
  png2icns "${BUNDLE_DIR}/Resources/Icon.icns" "appicon_darwin.png"
fi



plutil -convert binary1 "${BUNDLE_DIR}/Info.plist"

export BUNDLE_DIR="$(realpath "../../")/build/${SDK}/Honeycomb.app"
codesign --force --deep --sign - "${BUNDLE_DIR}"

touch "${BUNDLE_DIR}"