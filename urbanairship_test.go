package urbanairship

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendNotificationByTag(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(pushHandle))
	defer srv.Close()

	c := NewUAClient("appKey", "masterKey", "channelType")

	c.Message.Audience.Tag = "tag"
	c.Message.DeviceTypes = []string{"android"}
	c.Message.Notification.Alert = "Test push notification by tag"

	reponse, err := c.Send()
	if err != nil {
		t.Error("Response Error : ", err)
	}
	if reponse == nil {
		t.Error("Res is nil")
	}
}

func TestSendNotificationByNamedUser(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(pushHandle))
	defer srv.Close()

	c := NewUAClient("appKey", "masterKey", "channelType")

	c.Message.Audience.NamedUser = "named-user"
	c.Message.DeviceTypes = []string{"android"}
	c.Message.Notification.Alert = "Test push notification by named user"

	reponse, err := c.Send()
	if err != nil {
		t.Error("Response Error : ", err)
	}
	if reponse == nil {
		t.Error("Res is nil")
	}
}

func TestSendNotificationBy(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(pushHandle))
	defer srv.Close()

	c := NewUAClient("appKey", "masterKey", "channelType")

	c.Message.Audience.AndroidChannelId = "Android-channel-id"
	c.Message.DeviceTypes = []string{"android"}
	c.Message.Notification.Alert = "Test push notification by channel id"

	reponse, err := c.Send()
	if err != nil {
		t.Error("Response Error : ", err)
	}
	if reponse == nil {
		t.Error("Res is nil")
	}
}

func pushHandle(w http.ResponseWriter, r *http.Request) {
	result := `{"ok": true,"operation_id": "73c2a3bf-8efb-4837-8718-4e4d9f639898","push_ids": ["062d57da-7d35-49c7-9ead-b22533c93904"]}`
	fmt.Fprintln(w, result)
}
