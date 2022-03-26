package testtool

import (
	"strings"

	"github.com/google/uuid"
)

func GenReqId() string {
	return strings.ReplaceAll(strings.ToLower(uuid.New().String()), "-", "")
}

func CreateAsyncResponse() (string, []byte) {
	id := GenReqId()
	bs := []byte(`{"request_id": "` + id + `","jobs_url": "https://dpi.dns-platform.jp/v1/jobs/` + id + `"}`)
	return id, bs
}
