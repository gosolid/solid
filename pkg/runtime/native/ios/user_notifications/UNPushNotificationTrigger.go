//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

/*
interface UserNotifications.UNPushNotificationTrigger : UserNotifications.UNNotificationTrigger
*/

type UNPushNotificationTrigger struct {
  *UNNotificationTrigger

}

