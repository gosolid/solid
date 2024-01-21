//go:build ios

package main

/*
#cgo ios CFLAGS: -x objective-c
#cgo ios LDFLAGS: -lobjc -framework UIKit -framework Foundation -framework CoreGraphics -framework WebKit
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>

#include <Foundation/Foundation.h>
#include <UIKit/UIKit.h>
#include <WebKit/WebKit.h>

@interface AppDelegate : UIResponder <UIApplicationDelegate>
@property(strong, nonatomic) UIWindow *window;
@property(strong, nonatomic) WKWebView *webView;
@end

@implementation AppDelegate

- (BOOL)application:(UIApplication *)application
    didFinishLaunchingWithOptions:(NSDictionary *)launchOptions {
	NSLog(@"Hello from Grexie Cloud");
  self.window = [[UIWindow alloc] initWithFrame:[[UIScreen mainScreen] bounds]];
	UIViewController *rootViewController = [UIViewController new];
	self.webView = [[WKWebView alloc] initWithFrame:[[UIScreen mainScreen] bounds]];
	self.webView.autoresizingMask = UIViewAutoresizingFlexibleHeight | UIViewAutoresizingFlexibleWidth;
	[rootViewController.view addSubview:self.webView];
	self.window.rootViewController = rootViewController;

	NSURL *url = [NSURL URLWithString:@"http://localhost:3086/"];
	NSURLRequest *urlReq = [NSURLRequest requestWithURL:url];

	[self.webView loadRequest:urlReq];

	[self.window makeKeyAndVisible];
  return YES;
}

@end

int iosmain(int argc, char* argv[]) {
	@autoreleasepool {
		return UIApplicationMain(argc, argv, nil, NSStringFromClass([AppDelegate class]));
	}
}

*/
import "C"

import (
	"runtime"
)

func main() {
	runtime.LockOSThread()

	// startWorker()

	C.iosmain(0, nil)
}
