// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsUpdateNotificationDisplayOption undocumented
type WindowsUpdateNotificationDisplayOption int

const (
	// WindowsUpdateNotificationDisplayOptionVNotConfigured undocumented
	WindowsUpdateNotificationDisplayOptionVNotConfigured WindowsUpdateNotificationDisplayOption = 0
	// WindowsUpdateNotificationDisplayOptionVDefaultNotifications undocumented
	WindowsUpdateNotificationDisplayOptionVDefaultNotifications WindowsUpdateNotificationDisplayOption = 1
	// WindowsUpdateNotificationDisplayOptionVRestartWarningsOnly undocumented
	WindowsUpdateNotificationDisplayOptionVRestartWarningsOnly WindowsUpdateNotificationDisplayOption = 2
	// WindowsUpdateNotificationDisplayOptionVDisableAllNotifications undocumented
	WindowsUpdateNotificationDisplayOptionVDisableAllNotifications WindowsUpdateNotificationDisplayOption = 3
)

// WindowsUpdateNotificationDisplayOptionPNotConfigured returns a pointer to WindowsUpdateNotificationDisplayOptionVNotConfigured
func WindowsUpdateNotificationDisplayOptionPNotConfigured() *WindowsUpdateNotificationDisplayOption {
	v := WindowsUpdateNotificationDisplayOptionVNotConfigured
	return &v
}

// WindowsUpdateNotificationDisplayOptionPDefaultNotifications returns a pointer to WindowsUpdateNotificationDisplayOptionVDefaultNotifications
func WindowsUpdateNotificationDisplayOptionPDefaultNotifications() *WindowsUpdateNotificationDisplayOption {
	v := WindowsUpdateNotificationDisplayOptionVDefaultNotifications
	return &v
}

// WindowsUpdateNotificationDisplayOptionPRestartWarningsOnly returns a pointer to WindowsUpdateNotificationDisplayOptionVRestartWarningsOnly
func WindowsUpdateNotificationDisplayOptionPRestartWarningsOnly() *WindowsUpdateNotificationDisplayOption {
	v := WindowsUpdateNotificationDisplayOptionVRestartWarningsOnly
	return &v
}

// WindowsUpdateNotificationDisplayOptionPDisableAllNotifications returns a pointer to WindowsUpdateNotificationDisplayOptionVDisableAllNotifications
func WindowsUpdateNotificationDisplayOptionPDisableAllNotifications() *WindowsUpdateNotificationDisplayOption {
	v := WindowsUpdateNotificationDisplayOptionVDisableAllNotifications
	return &v
}
