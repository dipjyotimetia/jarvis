### Example Usage
```md
```sh
go build -o ./dist -ldflags="-X 'github.com/dipjyotimetia/jarvis/cmd/cmd.version=0.0.2'" ./...
.\dist\jarvis.exe generate-scenarios --path="specs/openapi/v3.0"
```

**Positive Test Cases:**

* **List all pets (GET /pets):
**
    * Send a GET request to `/pets` without any query parameters.
    * Expected response: Status code 200, a list
 of pets in JSON format.
* **Create a pet (POST /pets):**
    * Send a POST request to `/pets` with a valid pet object in the request body.
    * Expected response: Status code 201, no response body.
* **Get pet by ID (
GET /pets/{petId}):**
    * Send a GET request to `/pets/{petId}` with a valid pet ID.
    * Expected response: Status code 200, the pet object in JSON format.

**Negative Test Cases:**

* **List all pets with invalid limit (GET /pets):**
    * Send a GET request to `/pets` with an invalid limit value (e.g., -10).
    * Expected response: Status code 400 (Bad Request), an error message in JSON format.
* **Create a pet with invalid data (POST
 /pets):**
    * Send a POST request to `/pets` with an invalid pet object (e.g., missing required properties).
    * Expected response: Status code 400 (Bad Request), an error message in JSON format.
* **Get pet by invalid ID (GET /pets/{petId}):**
    * Send a GET request to `/pets/{petId}` with an invalid pet ID (e.g., "abc").
    * Expected response: Status code 404 (Not Found), an error message in JSON format.

**Edge Cases:**

* **List all pets with maximum limit (GET /pets):**
    * Send a GET request to `/pets` with the maximum allowed limit value (100).
    * Expected response: Status code 200, a list of 100 pets in JSON format.
* **Create a pet with empty name (POST /pets):**
    * Send a POST request to `/pets` with a pet object with an empty name.
    * Expected response: Status code 201, no response body.
* **Get pet with a non-existent ID (GET /pets/{petId
}):**
    * Send a GET request to `/pets/{petId}` with a non-existent pet ID.
    * Expected response: Status code 404 (Not Found), an error message in JSON format.
```


```md
```sh
 .\dist\jarvis.exe generate-test --path="specs/proto" --output="output"
```
```go
import (
	"context"
	"fmt"

	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go
-resty/resty/v2"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"

	notifypb "github.com/username
/project/notify/notifypb"
)

func TestCreateNotification(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/v1/notifications", r.URL.Path)

		var req notifypb.CreateNotificationRequest
		assert.NoError(t, jsonpb.Unmarshal(r.Body, &req))

		resp := &notifypb
.CreateNotificationResponse{
			Notification: &notifypb.Notification{
				Id:           1,
				Title:        "Test Notification",
				Message:      "This is a test notification.",
				RecipientId:  "user-1",
				Timestamp:    "2023-03-08 15:06:30",
				Status:       "sent",
			},
		}
		jsonpb.Marshal(w, resp)
	}))
	defer srv.Close()

	client := resty.New()
	client.SetBaseURL(srv.URL)

	now := time.Now()
	timestamp, err := ptypes.TimestampProto(now)
	assert.NoError(t, err)

	req := &notifypb.CreateNotificationRequest{
		Notification: &notifypb.Notification{
			Title:        "Test Notification",
			Message:      "This is a test notification.",
			RecipientId:  "user-1",
			Timestamp:    timestamp.String(),
			Status:       "sent",
		},
	}
	resp, err := client.
R().
		SetBody(req).
		Post("/v1/notifications")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var respProto notifypb.CreateNotificationResponse
	assert.NoError(t, jsonpb.Unmarshal(resp.Body(), &respProto))
	assert.Equal(t, int32(1), respProto.Notification.Id)
	assert.Equal(t, "Test Notification", respProto.Notification.Title)
	assert.Equal(t, "This is a test notification.", respProto.Notification.Message)
	assert.Equal(t, "user-1", respProto.Notification.RecipientId)
	assert.Equal(t, timestamp.String(), respProto.Notification.Timestamp)
	assert.Equal(t, "sent", respProto.Notification.Status)
}

func TestGetNotifications(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/notifications/user-1", r.URL.Path)

		resp := &notifypb.GetNotificationsResponse{
			Notifications: &notifypb.NotificationList{
				Notifications: []*notifypb.Notification{
					{
						Id:           1,
						Title:        "Test Notification 1",
						Message:      "This is a test notification 1.",
						RecipientId:  "user-1",
						Timestamp:    "2023-03-08 15:06:30",
						Status:       "sent",
					},
					{
						Id:           2,
						Title:        "Test Notification 2",
						Message:      "This is a test notification 2.",
						RecipientId:  "user-1",
						Timestamp:    "2023-03-08 15:07:00",
						Status:       "delivered",
					},
				},
			},
		}
		jsonpb.Marshal(w, resp)
	}))
	defer srv.
Close()

	client := resty.New()
	client.SetBaseURL(srv.URL)

	resp, err := client.R().
		Get("/v1/notifications/user-1")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var respProto notifypb.GetNotificationsResponse
	assert.NoError(t, jsonpb.Unmarshal(resp.Body(), &respProto))
	assert.Equal(t, 2, len(respProto.Notifications.Notifications))
	for _, n := range respProto.Notifications.Notifications {
		assert.Contains(t, []string{"sent", "delivered"}, n.Status)
	}
}
```

```