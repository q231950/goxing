// oauthconsumer.go
package xingapi

import (
	"github.com/garyburd/go-oauth/oauth"
	"github.com/str1ngs/ansi/color"
	"github.com/etix/stoppableListener"
	"net/http"
	"net/url"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

type AuthenticateHandler func()
type ResponseHandler func(io.Reader)

type OAuthConsumer struct {
	Client oauth.Client
	TemporaryCredentials oauth.Credentials
	OAuthCredentials oauth.Credentials
	Authenticated bool
	AuthenticateHandlers []AuthenticateHandler
	listener *stoppableListener.StoppableListener
}

func (consumer *OAuthConsumer) Connect(handler AuthenticateHandler) {
	
	credentialStore := new(CredentialStore)
	storedCredentials, _ := credentialStore.Credentials()
	
	if consumer.Authenticated {
		consumer.HandleAuthentication()

	} else if (0 < len(storedCredentials.Token)) {
		println("Authenticating with " + storedCredentials.Token + " " + storedCredentials.Secret + ".")

		consumer.Client.Credentials = consumer.OAuthConsumerCredentials()
		consumer.Authenticated = true
		consumer.OAuthCredentials = oauth.Credentials{storedCredentials.Token, storedCredentials.Secret}
		
		consumer.HandleAuthentication()
	} else {
		err := consumer.getRequestToken()
		if err != nil {
			log.Fatal("Connect:", err)
		}

		listener, err := net.Listen("tcp", "127.0.0.1:8080")

		consumer.listener = stoppableListener.Handle(listener)

		/* Handle SIGTERM (Ctrl+C) */
		k := make(chan os.Signal, 1)
		signal.Notify(k, os.Interrupt)
		go func() {
			<-k
			consumer.listener.Stop <- true
		}()

		http.HandleFunc("/", consumer.onReceiveTemporaryVerifierAndToken)

		err = http.Serve(consumer.listener, nil)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (consumer *OAuthConsumer)getRequestToken() error {
	println("Get request token...")
	
	consumer.Client = consumer.OauthClient()
	httpClient := new(http.Client)
	cred, err := consumer.Client.RequestTemporaryCredentials(httpClient, "http://localhost:8080/", nil)
	if err == nil {
		consumer.TemporaryCredentials = *cred
		println("Received token:%s", consumer.TemporaryCredentials.Token, " - secret:%s", consumer.TemporaryCredentials.Secret)

		tc := consumer.TemporaryCredentials
		accessUrl := consumer.Client.AuthorizationURL(&tc, nil)
		print("Please paste this url into your browser:")
		color.Print("m", fmt.Sprintf("%s%s", accessUrl, "\n"))
	}

	return err
}

func (consumer *OAuthConsumer)OAuthConsumerCredentials() (oauth.Credentials){
	credentials := NewCredentials()
	return oauth.Credentials {credentials.Token, credentials.Secret}
}

func (consumer *OAuthConsumer)OauthClient() (oauth.Client) {
	client := new(oauth.Client)
	client.Credentials = consumer.OAuthConsumerCredentials()
	client.TemporaryCredentialRequestURI = "https://api.xing.com/v1/request_token"
	client.ResourceOwnerAuthorizationURI = "https://api.xing.com/v1/authorize"
	client.TokenRequestURI = "https://api.xing.com/v1/access_token"
	client.SignatureMethod = oauth.PLAINTEXT
	return *client
}

func (consumer *OAuthConsumer)onReceiveTemporaryVerifierAndToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	verifier, _ := r.Form["oauth_verifier"]
	token, _ := r.Form["oauth_token"]

	fmt.Fprintf(w, "<html><head></head><body>")
	if (0 < len(verifier) && 0 < len(token)) {
		println("Temporary verfifier:", verifier[0])
		println("Temporary token:", token[0])
		
		httpClient := new(http.Client)
		tc := consumer.TemporaryCredentials
		credentials, _, _ := consumer.Client.RequestToken(httpClient, &tc, verifier[0])
		
		credentialStore := new(CredentialStore)
		credentialStore.SaveCredentials(Credentials{credentials.Token, credentials.Secret})

		consumer.OAuthCredentials = *credentials
		consumer.HandleAuthentication()
		consumer.listener.Stop <- true
		fmt.Fprintf(w, "<h3><span style=\"font-family:'Helvetica'; color:#777\">Success, you are authenticated.</span></h3>")
	} else {
		fmt.Fprintf(w, "<h3>Failure, something went wrong</h3>")
	}
    fmt.Fprintf(w, "</body></html>")
}

func (consumer *OAuthConsumer) Get(path string, parameters url.Values, handler ResponseHandler) {	
	consumer.AddAuthenticationHandler(func () {
			httpClient := new(http.Client)
			url := "https://api.xing.com" + path
			credentials := consumer.OAuthCredentials
			resp, _ := consumer.Client.Get(httpClient, &credentials, url, parameters)
			color.Printf("c", fmt.Sprintf("GET %s\n", path))
			consumer.PrintResponse(resp)
			if resp.StatusCode == 200 {
				handler(resp.Body)
			} 

			defer resp.Body.Close()
		})

	if !consumer.Authenticated {
		consumer.Connect(func (){
			consumer.HandleAuthentication()
		})
	}
}

func (consumer *OAuthConsumer) PrintResponse(response *http.Response) {
	var colorCode string
	if (response.StatusCode == 200) {
		colorCode = "g"
	} else {
		colorCode = "r"
	}
	color.Printf(colorCode, fmt.Sprintf("%s\n", response.Status))
}

func (consumer *OAuthConsumer) AddAuthenticationHandler(handler AuthenticateHandler) {
	if consumer.AuthenticateHandlers == nil {
		consumer.AuthenticateHandlers = []AuthenticateHandler{}
	}
	
	if consumer.Authenticated {
		handler()
	} else {
		consumer.AuthenticateHandlers = append(consumer.AuthenticateHandlers, handler)
	}
}

func (consumer *OAuthConsumer) HandleAuthentication() {
	consumer.Authenticated = true
	println("handle authentication")
	for _, authenticateHandler := range consumer.AuthenticateHandlers {
		authenticateHandler()
	}
}
