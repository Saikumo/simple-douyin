package common

const (
	MAX_USERNAME_LENGTH    int = 32
	MIN_PASSWORD_LENGTH    int = 6
	MAX_PASSWORD_LENGTH    int = 32
	MAX_VIDEO_TITLE_LENGTH int = 50

	SUCCESS_STATUS_CODE int = 0
	FAILURE_STATUS_CODE int = 1
)

var (
	SUPPORTED_VIDEO_TYPE_SET = map[string]struct{}{
		".mp4": {},
	}
)
