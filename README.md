# go-urbanairship : Urbanairship Library for Go

Urbanairship Push Notification Library using golang ( Go )

###### Features

* Send push notification by tag
* Send push notification by named user
* Send push notification by channel id
    - Install 1st Flight app in your device (android/ios)
	- Get channel id from app 

## Usage

```
go get github.com/heaptracetechnology/go-urbanairship
```

#### Urbanairship tutorials docs
```
https://docs.urbanairship.com/tutorials/getting-started/1st-flight-app/
```

#### Urbanairship create project

Create project in urbanairship:

1. Login with https://go.urbanairship.com
2. Select New project
3. Create project with all details


#### Get App key and Master Secret key

App key and Master Secret key can be found in:

1. Urbanairship project settings
2. APIs and Integrations
3. then copy the app key and click on textbox below master secret key to revel master key

# Examples

### Send push notification by Tag

```go

package main

import (
	"fmt"
	"github.com/heaptracetechnology/go-urbanairship"
)

func main() {

	appKey := "APP_KEY"
	masterKey := "MASTER_KEY"

	var message urbanairship.UAMessage
	var audiance urbanairship.Audiance
	var notification urbanairship.Notification

	audiance.Tag = "tag-name" // from app setting
	notification.Alert = "Push notification with tag name" // alert message
	channelType := ""

	message.Audience = audiance
	message.Notification = notification
	message.DeviceTypes = []string{"android"}

	client := *urbanairship.NewUAClient(appKey, masterKey, channelType)
	client.Message = message

	response, err := client.Send()
	if err == nil {
		fmt.Println(response)
	} else {
		fmt.Println(err)
	}

}


```


### Send push notification by Named User

```go

package main

import (
	"fmt"
	"github.com/heaptracetechnology/go-urbanairship"
)

func main() {

	appKey := "APP_KEY"
	masterKey := "MASTER_KEY"

	var message urbanairship.UAMessage
	var audiance urbanairship.Audiance
	var notification urbanairship.Notification

	audiance.NamedUser = "named_user" // from app setting
	notification.Alert = "Push notification with named user" // alert message
	channelType := ""

	message.Audience = audiance
	message.Notification = notification
	message.DeviceTypes = []string{"android"}

	client := *urbanairship.NewUAClient(appKey, masterKey, channelType)
	client.Message = message

	response, err := client.Send()
	if err == nil {
		fmt.Println(response)
	} else {
		fmt.Println(err)
	}

}


```

### Send push notification by Channel Id

```go

package main

import (
	"fmt"
	"github.com/heaptracetechnology/go-urbanairship"
)

func main() {

	appKey := "APP_KEY"
	masterKey := "MASTER_KEY"

	var message urbanairship.UAMessage
	var audiance urbanairship.Audiance
	var notification urbanairship.Notification

	audiance.AndroidChannelId = "Channel_Id" // from app setting
	notification.Alert = "Push notification by channel id" // alert message
	channelType := "android" // required field ios/android 

	message.Audience = audiance
	message.Notification = notification
	message.DeviceTypes = []string{"android"}

	client := *urbanairship.NewUAClient(appKey, masterKey, channelType)
	client.Message = message
	response, err := client.Send()
	if err == nil {
		fmt.Println(response)
	} else {
		fmt.Println(err)
	}

}

```