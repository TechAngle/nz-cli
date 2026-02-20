package api

import (
	"fmt"
	"strconv"
)

// Get notifications.
// Method: GET
func (c *NZAPIClient) Notifications() (*NotificationsResponse, error) {
	response, err := sendRequest[NotificationsResponse](c, GetMethod, notificationsEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send notifications request: %v", err)
	}

	// their shitty problems
	if isNZError(response.ErrorMessage) {
		return nil, fmt.Errorf("nz.ua notifications problem: %s", response.ErrorMessage)
	}

	return response, nil
}

// Get unread notifications integer instead of their string.
// If error occured - returns -1 (just because count of unread notifications cannot be negative)
// Method: GET
func (c *NZAPIClient) UnreadNotifications() (int, error) {
	response, err := sendRequest[UnreadNotificationsResponse](c, GetMethod, unreadNotificationsEndpoint, nil)
	if err != nil {
		return -1, fmt.Errorf("failed to get qty of unread notifications: %v", err)
	}

	if isNZError(response.ErrorMessage) {
		return -1, fmt.Errorf("nz.ua notifications qty problem: %s", response.ErrorMessage)
	}

	qty, err := strconv.Atoi(response.Qty)
	if err != nil {
		return -1, fmt.Errorf("failed to convert qty to integer: %v", err)
	}

	return qty, nil
}
