
typedef void *WebContentsPtr;
typedef void *NativeViewPtr;

#ifdef __cplus_plus
extern "C"
{
#endif

  void Chromium_Initialize();
  WebContentsPtr Chromium_CreateWebContents();
  NativeViewPtr Chromium_WebContents_GetNativeView(WebContentsPtr pWebContents);
  void Chromium_WebContents_OpenURL(const char *url);

#ifdef __cplus_plus
}
#endif