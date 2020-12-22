package zoom

import "fmt"

// CreateMeetingOptions are the options to create a meeting with
type UpdateMeetingStatusOptions struct {
	MeetingID string `json:"-"`
	Action    string `json:"action,omitempty"`
}

// CreateMeetingPath - v2 create a meeting for a user
const UpdateMeetingStatusPath = "/meetings/%s/status"

// CreateMeeting calls POST /users/{userId}/meetings
func UpdateMeetingStatus(opts UpdateMeetingStatusOptions) error {
	return defaultClient.UpdateMeetingStatus(opts)
}

// CreateMeeting calls POST /users/{userId}/meetings
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingstatus
func (c *Client) UpdateMeetingStatus(opts UpdateMeetingStatusOptions) error {
	return c.requestV2(requestV2Opts{
		Method:         Put,
		Path:           fmt.Sprintf(UpdateMeetingStatusPath, opts.MeetingID),
		DataParameters: &opts,
		HeadResponse:   true,
	})
}
