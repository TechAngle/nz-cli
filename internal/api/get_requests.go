package api

import (
	"fmt"
	"nz-cli/internal/models"
	"strconv"
)

// Get notifications.
// Method: GET
func (c *NZAPIClient) Notifications() (*models.Notifications, error) {
	var response models.Notifications
	err := c.SendRequest(GetMethod, ApiEndpoint+NotificationsEndpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("faield to send notifications request: %v", err)
	}

	// their shitty problems
	if IsNZError(response.ErrorMessage) {
		return nil, fmt.Errorf("nz.ua notifications problem: %s", response.ErrorMessage)
	}

	return &response, nil
}

// get unread notifications integer instead of their string.
// if error occured - returns -1 (just because count of unread notifications cannot be negative)
// Method: GET
func (c *NZAPIClient) UnreadNotifications() (int, error) {
	var response models.UnreadNotificationsResponse
	err := c.SendRequest(GetMethod, ApiEndpoint+UnreadNotificationsEndpoint, nil, &response)
	if err != nil {
		return -1, fmt.Errorf("failed to get qty of unread notifications: %v", err)
	}

	if IsNZError(response.ErrorMessage) {
		return -1, fmt.Errorf("nz.ua notifications qty problem: %s", response.ErrorMessage)
	}

	qty, err := strconv.Atoi(response.Qty)
	if err != nil {
		return -1, fmt.Errorf("failed to convert qty to integer: %v", err)
	}

	return qty, nil
}
