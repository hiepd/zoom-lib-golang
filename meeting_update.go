package zoom

import "fmt"

// CreateMeetingOptions are the options to create a meeting with
type UpdateMeetingStatusOptions struct {
	MeetingID string `json:"meeting_id"`
	Action    string `json:"action"`
}

// CreateMeetingPath - v2 create a meeting for a user
const UpdateMeetingStatusPath = "/meetings/%s/status"

// CreateMeeting calls POST /users/{userId}/meetings
func UpdateMeetingStatus(opts UpdateMeetingStatusOptions) (Meeting, error) {
	return defaultClient.UpdateMeetingStatus(opts)
}

// CreateMeeting calls POST /users/{userId}/meetings
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingstatus
func (c *Client) UpdateMeetingStatus(opts UpdateMeetingStatusOptions) (Meeting, error) {
	var ret = Meeting{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Put,
		Path:           fmt.Sprintf(UpdateMeetingStatusPath, opts.MeetingID),
		DataParameters: &opts,
		Ret:            &ret,
	})
}
