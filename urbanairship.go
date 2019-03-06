package urbanairship

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// ua_server_url fcm server url
	ua_server_url = "https://go.urbanairship.com/api/push/"
)

var (
	// uaServerUrl for testing purposes
	uaServerUrl = ua_server_url
)

// UrbanAirshipMsg represents usrbanairship request message
type UAMsg struct {
	NamedUser    string      `json:"nameduser,omitempty"`
	Tag          string      `json:"tag,omitempty"`
	ChannelId    string      `json:"channelid,omitempty"`
	DeviceTypes  []string    `json:"devicetypes,omitempty"`
	Notification interface{} `json:"notification,omitempty"`
}

// UrbanAirshipResponseStatus represents urban airship response message
type UAResponseStatus struct {
	Ok            bool
	StatusCode    int
	Operation_id  string      `json:"operation_id"`
	Push_ids      []string    `json:"push_ids"`
	Message_ids   []string    `json:"message_ids,omitempty"`
	Content_urls  []string    `json:"content_urls,omitempty"`
	Localized_ids []string    `json:"localized_ids,omitempty"`
	Error         string      `json:"error,omitempty"`
	Error_code    int         `json:"error_code,omitempty"`
	Details       interface{} `json:"details,omitempty"`
}

// UrbanAirshipClient struct
type UAClient struct {
	ApiKey        string
	MasterKey     string
	Authorization string
	Message       UAMsg
}

// NewUAClient generates the value of the Authorization key
func NewUAClient(apiKey string, masterKey string) *UAClient {
	ua := new(UAClient)
	ua.ApiKey = apiKey
	ua.MasterKey = masterKey
	generateAuth := apiKey + masterKey

	ua.Authorization = base64.StdEncoding.EncodeToString([]byte(generateAuth))

	return ua
}

// authorizationHeader generates the value of the Authorization key
func (this *UAClient) authorizationHeader() string {
	return fmt.Sprintf("key=%v", this.Authorization)
}

// NewUATagsMsg sets the targeted tagged devices
func (this *UAClient) NewUATagsMsg(authorizationKey string, tag string, devicetypes []string, notification interface{}) *UAClient {

	this.NewSendTagMsg(authorizationKey, tag, devicetypes, notification)

	return this
}

// NewUANamedUserMsg sets the targeted nameuser
func (this *UAClient) NewUANamedUserMsg(authorizationKey string, nameduser string, devicetypes []string, notification interface{}) *UAClient {

	this.NewSendnamedUserMsg(authorizationKey, nameduser, devicetypes, notification)

	return this
}

// NewUAChannelIdMsg sets the targeted to channelid
func (this *UAClient) NewUAChannelIdMsg(authorizationKey string, channelid string, devicetypes []string, notification interface{}) *UAClient {

	this.NewSendChannelIdMsg(authorizationKey, channelid, devicetypes, notification)

	return this
}

// NewSendTagMsg sets the targeted tag and the data payload
func (this *UAClient) NewSendTagMsg(authorizationKey string, tag string, devicetypes []string, notification interface{}) *UAClient {

	this.Authorization = authorizationKey
	this.Message.Tag = tag
	this.Message.DeviceTypes = devicetypes
	this.Message.Notification = notification

	return this
}

// NewSendnamedUserMsg sets the targeted nameduser and the data payload
func (this *UAClient) NewSendnamedUserMsg(authorizationKey string, nameduser string, devicetypes []string, notification interface{}) *UAClient {

	this.Authorization = authorizationKey
	this.Message.NamedUser = nameduser
	this.Message.DeviceTypes = devicetypes
	this.Message.Notification = notification

	return this
}

// NewSendChannelIdMsg sets the targeted channelid and the data payload
func (this *UAClient) NewSendChannelIdMsg(authorizationKey string, channelid string, devicetypes []string, notification interface{}) *UAClient {

	this.Authorization = authorizationKey
	this.Message.ChannelId = channelid
	this.Message.DeviceTypes = devicetypes
	this.Message.Notification = notification

	return this
}

// toJsonByte converts uaMsg to a json byte
func (this *UAMsg) toJsonByte() ([]byte, error) {

	return json.Marshal(this)

}

// parseStatusBody parse UA response body
func (this *UAResponseStatus) parseStatusBody(body []byte) error {

	if err := json.Unmarshal([]byte(body), &this); err != nil {
		return err
	}
	return nil

}

// sendOnce send a single request to ua
func (this *UAClient) sendOnce() (*UAResponseStatus, error) {

	uaRespStatus := new(UAResponseStatus)

	jsonByte, err := this.Message.toJsonByte()
	if err != nil {
		return uaRespStatus, err
	}

	request, err := http.NewRequest("POST", uaServerUrl, bytes.NewBuffer(jsonByte))
	request.Header.Set("Authorization", "Basic X2kzWkh3b1VTeEtKekRfb0ExUXVDUXJQT1pwOVdzUTFpLWJRVjZuWUpwU0E=")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/vnd.urbanairship+json; version=3")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return uaRespStatus, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uaRespStatus, err
	}

	uaRespStatus.StatusCode = response.StatusCode

	//uaRespStatus.RetryAfter = response.Header.Get(retry_after_header)

	if response.StatusCode != 200 {
		return uaRespStatus, nil
	}

	err = uaRespStatus.parseStatusBody(body)
	if err != nil {
		return uaRespStatus, err
	}
	uaRespStatus.Ok = true

	return uaRespStatus, nil
}

// Send to ua
func (this *UAClient) Send() (*UAResponseStatus, error) {
	return this.sendOnce()

}
