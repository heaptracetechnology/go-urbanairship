package urbanairship

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	// ua_server_url fcm server url
	UA_SERVER_URL = "https://go.urbanairship.com/api/push/"
)

type Audiance struct {
	Tag              string `json:"tag,omitempty"`
	AndroidChannelId string `json:"android_channel,omitempty"`
	IOSChannelId     string `json:"ios_channel,omitempty"`
	NamedUser        string `json:"named_user,omitempty"`
}

type Notification struct {
	Alert string `json:"alert,omitempty"`
}

type UAMessage struct {
	Audience     Audiance     `json:"audience,omitempty"`
	DeviceTypes  []string     `json:"device_types,omitempty"`
	Notification Notification `json:"notification,omitempty"`
}

// UrbanAirshipResponseStatus represents urban airship response message
type UAResponseStatus struct {
	Ok           bool        `json:"ok"`
	OperationId  string      `json:"operation_id"`
	PushIds      []string    `json:"push_ids"`
	MessageIds   []string    `json:"message_ids,omitempty"`
	ContentURLs  []string    `json:"content_urls,omitempty"`
	LocalizedIds []string    `json:"localized_ids,omitempty"`
	Error        string      `json:"error,omitempty"`
	ErrorCode    int         `json:"error_code,omitempty"`
	Details      interface{} `json:"details,omitempty"`
}

// UrbanAirshipClient struct
type UAClient struct {
	ApiKey        string
	MasterKey     string
	Authorization string
	ChannelType   string
	Message       UAMessage
}

// NewUAClient generates the value of the Authorization key
func NewUAClient(apiKey string, masterKey string, channelType string) *UAClient {
	ua := new(UAClient)
	ua.ApiKey = apiKey
	ua.MasterKey = masterKey
	generateAuth := apiKey + ":" + masterKey

	ua.Authorization = base64.StdEncoding.EncodeToString([]byte(generateAuth))

	return ua
}

// NewUATagsMessage sets the targeted tagged devices
func (this *UAClient) NewUATagsMessage(authorizationKey string, tag string, deviceTypes []string, notification Notification) *UAClient {

	this.NewSendTagMessage(authorizationKey, tag, deviceTypes, notification)

	return this
}

// NewUANamedUserMessage sets the targeted nameuser
func (this *UAClient) NewUANamedUserMessage(authorizationKey string, namedUser string, deviceTypes []string, notification Notification) *UAClient {

	this.NewSendnamedUserMessage(authorizationKey, namedUser, deviceTypes, notification)

	return this
}

// NewUAChannelIdMessage sets the targeted to channelid
func (this *UAClient) NewUAChannelIdMessage(authorizationKey string, channelId string, channelType string, deviceTypes []string, notification Notification) *UAClient {

	this.NewSendChannelIdMessage(authorizationKey, channelId, channelType, deviceTypes, notification)

	return this
}

// NewSendTagMessage sets the targeted tag and the data payload
func (this *UAClient) NewSendTagMessage(authorizationKey string, tag string, deviceTypes []string, notification Notification) *UAClient {

	this.Authorization = authorizationKey
	this.Message.Audience.Tag = tag
	this.Message.DeviceTypes = deviceTypes
	this.Message.Notification = notification

	return this
}

// NewSendnamedUserMessage sets the targeted nameduser and the data payload
func (this *UAClient) NewSendnamedUserMessage(authorizationKey string, namedUser string, deviceTypes []string, notification Notification) *UAClient {

	this.Authorization = authorizationKey
	this.Message.Audience.NamedUser = namedUser
	this.Message.DeviceTypes = deviceTypes
	this.Message.Notification = notification

	return this
}

// NewSendChannelIdMessage sets the targeted channelid and the data payload
func (this *UAClient) NewSendChannelIdMessage(authorizationKey string, channelId string, channelType string, deviceTypes []string, notification Notification) *UAClient {

	this.Authorization = authorizationKey
	if channelType == "android" {
		this.Message.Audience.AndroidChannelId = channelId
	} else if channelType == "ios" {
		this.Message.Audience.IOSChannelId = channelId
	}

	this.Message.DeviceTypes = deviceTypes
	this.Message.Notification = notification

	return this
}

// ToJsonByte converts uaMessage to a json byte
func (this *UAMessage) ToJsonByte() ([]byte, error) {

	return json.Marshal(this)

}

// ParseStatusBody parse UA response body
func (this *UAResponseStatus) ParseStatusBody(body []byte) error {

	if err := json.Unmarshal([]byte(body), &this); err != nil {
		return err
	}
	return nil

}

// sendOnce send a single request to ua
func (this *UAClient) SendOnce() (*UAResponseStatus, error) {

	uaResponseStatus := new(UAResponseStatus)

	jsonByte, err := this.Message.ToJsonByte()
	if err != nil {
		return uaResponseStatus, err
	}

	request, err := http.NewRequest("POST", UA_SERVER_URL, bytes.NewBuffer(jsonByte))
	request.Header.Set("Authorization", "Basic "+this.Authorization)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return uaResponseStatus, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return uaResponseStatus, err
	}

	err = uaResponseStatus.ParseStatusBody(body)
	if err != nil {
		return uaResponseStatus, err
	}
	uaResponseStatus.Ok = true

	return uaResponseStatus, nil
}

// Send to ua
func (this *UAClient) Send() (*UAResponseStatus, error) {
	return this.SendOnce()

}
