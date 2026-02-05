package visuals

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	// Clock display format
	ClockTimeFormat = "Monday - 03:04:05 pm - 2006-01-02"

	// 30 second for account checking.
	AccountUpdateDelay = 30 * time.Second

	// 2 minutes for notifications update
	NotificationsUpdateDelay = 2 * time.Minute

	// Delay between requests
	RequestDelay = 5 * time.Second
)

// Start notifications status updating
func (c *CLI) startNotificationsUpdater() {
	go func() {
		for {
			c.app.QueueUpdateDraw(func() {
				// updating quantity of unread news
				q, err := c.client.UnreadNotifications()
				if err != nil {
					log.Println("failed to get unread news quantity:", err)
					return
				}

				c.mainState.unreadNewsQty.SetText(fmt.Sprintf("Unread news: %d", q))

				// sleeping a bit to not trigger rate limit
				time.Sleep(RequestDelay)

				// updating news list
				news, err := c.client.Notifications()
				if err != nil {
					log.Panicln("failed to get notifications:", err)
					c.modalState.message.SetText(fmt.Sprintf("failed to get notifications: %s", err))
					c.pages.SwitchToPage("modal")
					return
				}

				c.mainState.newsList.Clear()
				// parsing first 5 news
				for i, data := range news.Data {
					// ignoring more than 5 news
					if i > 4 {
						break
					}

					var mainText strings.Builder
					var subText strings.Builder

					fmt.Fprintf(&mainText,
						// format: time, subject
						"[Time: %s] (%s)",
						data.SentAt,
						data.NotificationData.LessonName,
					)

					fmt.Fprintf(&subText,
						"%s %s %s",
						// body, mark, comment
						data.Body,
						data.NotificationData.MarkValue,
						data.NotificationData.Comment,
					)

					log.Println(mainText.String(), subText.String())

					c.mainState.newsList.AddItem(mainText.String(), subText.String(), rune(i), nil)
				}
			})

			time.Sleep(NotificationsUpdateDelay)
		}
	}()
}

// Despite that account lays in AppData files, state saves in client structure
// and if user tried to update account - access token can be empty, so we i wanna check it
// and show him that he won't be able to send requests to nz.ua.
func (c *CLI) startAccountStateUpdater() {
	go func() {
		// creating a cycle with 30 second delay
		for {
			c.app.QueueUpdateDraw(func() {
				// if we have an account - show it
				if c.client.Authorized() {
					acc := c.client.Account()
					// TODO: Format it with styles
					c.mainState.loggedAccountLabel.SetText(fmt.Sprintf("Logged in as %s (Student ID: %d)", acc.FIO, acc.StudentID))
				} else {
					c.mainState.loggedAccountLabel.SetText("No account!")
				}
			})

			time.Sleep(AccountUpdateDelay)
		}
	}()
}

// start clock updater
func (c *CLI) startClockUpdater() {
	go func() {
		// update clock every second because we don't need less than it
		for {
			c.app.QueueUpdateDraw(func() {
				now := time.Now()

				c.mainState.clockLabel.SetText(now.Format(ClockTimeFormat))
			})

			time.Sleep(1 * time.Second)
		}
	}()
}
