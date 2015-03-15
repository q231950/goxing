// oauthconsumer.go
package xingapi

import (
	"fmt"	
	"net/http"
	"net/url"
	"io"
)

type ResponseHandler func(io.Reader, error)
type AuthenticateHandler func()

type OAuthConsumer struct {
	Authenticated bool
	oauthAuthenticator *OAuthAuthenticator
	AuthenticateHandlers []AuthenticateHandler
}

func (consumer *OAuthConsumer) authenticate(handler AuthenticateHandler) {
	
	if consumer.Authenticated {
		handler()
	} else {
		credentialStore := new(CredentialStore)
		storedCredentials, localeCredentialsError := credentialStore.Credentials()

		if (localeCredentialsError == nil) {
			consumer.oauthAuthenticator = new(OAuthAuthenticator)
			consumer.oauthAuthenticator.AuthenticateUsingStoredCredentials(storedCredentials, func(err error) {
				if err == nil {
					consumer.handleAuthentication()
				}
			})
		} else {
			consumer.requestCredentials(handler)
		}
	}
}

func (consumer *OAuthConsumer) requestCredentials(handler func()) {
	consumer.oauthAuthenticator = new(OAuthAuthenticator)
	consumer.oauthAuthenticator.Authenticate(func(err error) {
		if err == nil {
			handler()
		}
	})
}

func (consumer *OAuthConsumer) Get(path string, parameters url.Values, handler ResponseHandler) {	
	consumer.addAuthenticationHandler(func () {
			httpClient := new(http.Client)
			url := "https://api.xing.com" + path
			credentials := consumer.oauthAuthenticator.OAuthCredentials
			resp, err := consumer.oauthAuthenticator.Client.Get(httpClient, &credentials, url, parameters)
			PrintCommand(fmt.Sprintf("GET %s\n", path))
			PrintResponse(resp)
			if resp.StatusCode == 200 {
				handler(resp.Body, err)
			} 
			defer resp.Body.Close()
		})

	if !consumer.Authenticated {
		consumer.authenticate(func (){
			consumer.handleAuthentication()
		})
	}
}

func (consumer *OAuthConsumer) addAuthenticationHandler(handler AuthenticateHandler) {
	if consumer.AuthenticateHandlers == nil {
		consumer.AuthenticateHandlers = []AuthenticateHandler{}
	}
	
	if consumer.Authenticated {
		handler()
	} else {
		consumer.AuthenticateHandlers = append(consumer.AuthenticateHandlers, handler)
	}
}

func (consumer *OAuthConsumer) handleAuthentication() {
	consumer.Authenticated = true
	for _, authenticateHandler := range consumer.AuthenticateHandlers {
		authenticateHandler()
	}
}
