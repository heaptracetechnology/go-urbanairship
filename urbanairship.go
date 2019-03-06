package urbanairship

import (
	"encoding/base64"
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
