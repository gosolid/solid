#define __cplus_plus
#include "Chromium.h"

#include "content/public/browser/web_contents.h"
#include "chrome/app/chrome_main_mac.h"
#include "chrome/browser/chrome_browser_application_mac.h"

extern "C"
{
  __attribute__((visibility("default"))) int NO_STACK_PROTECTOR
  ChromeMain(int argc, const char **argv);

  void Chromium_Initialize()
  {
    chrome_browser_application_mac::RegisterBrowserCrApp();
  }

  WebContentsPtr Chromium_CreateWebContents()
  {
    std::shared_ptr<content::WebContents> webContents(content::WebContents::Create(
        content::WebContents::CreateParams(NULL)));
    return webContents.get();
  }

  NativeViewPtr Chromium_WebContents_GetNativeView(WebContentsPtr pWebContents)
  {
    NSView *nativeView = ((content::WebContents *)pWebContents)->GetNativeView().GetNativeNSView();
    return nativeView;
  }

  void Chromium_WebContents_OpenURL(const char *url)
  {
    std::string sURL(url);
    // GURL gURL(sURL);
  }
}