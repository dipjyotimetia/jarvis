### Example Usage
```md
```sh
go build -o ./dist -ldflags="-X 'github.com/dipjyotimetia/jarvis/cmd/cmd.version=0.0.2'" ./...
.\dist\jarvis.exe generate-scenarios --path="specs/openapi/v3.0/mini_blog.yaml"
```
**Test Scenario 1: Retrieve All Blog Posts**

* Precondition:
 The API server is up and running.
* Action: Send a GET request to the "/posts" endpoint.
* Expected Result: The server responds with
 a status code of 200 and a list of all blog posts in JSON format.

**Test Scenario 2: Create a New Blog Post**

* Precondition: The API server is up and running.
* Action: Send a POST request to the "/posts" endpoint with a valid JSON payload
 representing a new blog post.
* Expected Result: The server responds with a status code of 201 and the newly created blog post in JSON format.

**Test Scenario 3: Retrieve a Specific Blog Post**

* Precondition: The API server is up and running.
* Action: Send a GET request to the "/posts/{postId}" endpoint with a valid postId.
* Expected Result: The server responds with a status code of 200 and the details of the requested blog post in JSON format.

**Test Scenario 4: Update a Blog Post**

* Precondition: The API
 server is up and running.
* Action: Send a PATCH request to the "/posts/{postId}" endpoint with a valid postId and a JSON payload containing the updated fields.
* Expected Result: The server responds with a status code of 200 and the updated blog post in JSON format.

**Test Scenario 5: Delete a Blog Post**

* Precondition: The API server is up and running.
* Action: Send a DELETE request to the "/posts/{postId}" endpoint with a valid postId.
* Expected Result: The server responds with a status code of 204 and no content in the response body.

**Test Scenario 6: Retrieve Comments for a Blog Post**

* Precondition: The API server is up and running.
* Action: Send a GET request to the "/posts/{postId}/comments" endpoint with a valid postId.
* Expected Result: The server responds with a status code of 200 and a list of comments for the specified blog post in JSON format.

**Test Scenario 7: Add a New Comment**

* Precondition: The API server is up and running.
* Action: Send a POST request to the "/posts/{postId}/comments" endpoint
 with a valid postId and a JSON payload representing a new comment.
* Expected Result: The server responds with a status code of 201 and the newly created comment in JSON format.

**Test Scenario 8: Retrieve User Profile**

* Precondition: The API server is up and running.
* Action: Send a GET request to the "/users/{userId}" endpoint with a valid userId.
* Expected Result: The server responds with a status code of 200 and the user profile details in JSON format.

**Test Scenario 9: Handle Invalid Input**

* Precondition: The API server is up and running.
* Action: Send a request to an endpoint with invalid input (e.g., an invalid postId or userId).
* Expected Result: The server responds with a status code of 400 (Bad Request) and an error message in the response body.

**Test Scenario 10: Handle Non-existent Resources**

* Precondition: The API server is up and running.
* Action: Send a request to an endpoint with a non-existent resource (e.g., a postId or userId that does not exist).
* Expected Result: The server responds with a status code of 404 (Not Found) and an error message in the response body.
```


```md
```sh
 jarvis generate-test --path="specs/proto" --output="output"
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