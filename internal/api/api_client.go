package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/enetx/surf"
	cookiejar "github.com/juju/persistent-cookiejar"
)

const (
	SessionCookiesBase = "cookies.json"
	AccountStateBase   = "account.json"
)

var cookiesJar *cookiejar.Jar

type NZAPIClient struct {
	// account which we use
	account *AccountState
	client  *http.Client
	cookies []*http.Cookie
}

// is user authorized and has access token
func (c NZAPIClient) Authorized() bool {
	return c.account != nil && c.account.AccessToken != ""
}

// get current account state
func (c NZAPIClient) Account() AccountState {
	return *c.account
}

// save current session
// returns error if something has gone wrong
func (c NZAPIClient) SaveSession() error {
	err := cookiesJar.Save()
	err = c.account.Save(AccountStateBase)
	return err
}

// Create new api client
func NewApiClient() (apiClient *NZAPIClient, err error) {
	// New func loads from filename in private method newFromTime
	// that's why it has not Load() method ._.
	// p.s. i hate its dev
	cookiesJar, err = cookiejar.New(&cookiejar.Options{
		Filename: SessionCookiesBase,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %v", err)
	}

	// creating client with session, impersonating android chrome
	// because we use nz.ua mobile api
	client := surf.NewClient().
		Builder().
		Session().
		Impersonate().
		Chrome().
		Build().
		Unwrap()
	// log.Println("client loaded")

	stdClient := client.Std()
	stdClient.Jar = cookiesJar // setting 'fake' cookies jar with saving ability!

	state := AccountState{}
	err = state.Load(AccountStateBase)
	if err != nil {
		log.Println("failed to load account state:", err)
	} // else {
	// 	// log.Println(state)
	// 	// log.Println("account state loaded, logged in as:", state.FIO, state.StudentID)
	// }

	return &NZAPIClient{
		account: &state,
		client:  stdClient,
	}, nil
}
