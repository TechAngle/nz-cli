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

// updates unread value
func (c *TUI) updateUnreadQty() {
	// updating quantity of unread news
	q, err := c.client.UnreadNotifications()
	if err != nil {
		log.Println("failed to get unread news quantity:", err)
		return
	}

	c.mainState.unreadNewsQty.SetText(fmt.Sprintf("Unread news: [%s]%d[%s]", thirdCode, q, thirdCode))
}

// update news lists on main and news pages
func (c *TUI) updateNewsLists() {
	// updating news list
	news, err := c.client.Notifications()
	if err != nil {
		log.Println("failed to get notifications:", err)
		c.modalState.message.SetText(fmt.Sprintf("failed to get notifications: %s", err))
		c.pages.SwitchToPage("modal")
		return
	}

	// clearing news list, just that news were not doubled
	c.mainState.shortNewsList.Clear()
	c.newsState.newsList.Clear()

	// parsing first 5 news and add them to short list.
	// then add all news to general news list
	for i, data := range news.Data {
		var mainText, subText strings.Builder

		// formatting main text
		fmt.Fprintf(&mainText,
			// format: time, subject
			"- [Time: %s] (%s)",
			data.SentAt,
			data.NotificationData.LessonName,
		)

		// formatting subtext
		fmt.Fprintf(&subText,
			"\t%s %s %s %s",
			// body, lesson type (optional) mark (optional), comment (optional)
			data.Body,
			data.NotificationData.LessonType,
			data.NotificationData.MarkValue,
			data.NotificationData.Comment,
		)

		log.Println(mainText.String(), subText.String())

		// ignoring more than 5 news for main page
		if i < 4 {
			c.mainState.shortNewsList.AddItem(mainText.String(), subText.String(), rune(0), func() {
				c.pages.SwitchToPage(NewsPage)
			})
		}

		// adding item to news list on news page
		c.newsState.newsList.AddItem(mainText.String(), subText.String(), rune(0), nil)
	}
}

// Start notifications status updating
func (c *TUI) StartNotificationsUpdater() {
	go func() {
		for {
			c.app.QueueUpdateDraw(func() {
				c.updateUnreadQty()

				// sleeping a bit to not trigger rate limit
				time.Sleep(RequestDelay)

				c.updateNewsLists()

				c.app.Sync()
			})

			time.Sleep(NotificationsUpdateDelay)
		}
	}()
}

// Despite that account lays in AppData files, state saves in client structure
// and if user tried to update account - access token can be empty, so we i wanna check it
// and show him that he won't be able to send requests to nz.ua.
func (c *TUI) StartAccountStateUpdater() {
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

// Start clock updater
func (c *TUI) StartClockUpdater() {
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
