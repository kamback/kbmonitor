package logs

import (
	"adumonitor/model"
)

func Write(writeCh chan *Message) {
	for v := range writeCh {
		tags := map[string]string{"Path": v.Path, "Method": v.Method, "Scheme": v.Scheme, "Status": v.Status}
		fields := map[string]interface{}{
			"UpstreamTime": v.UpstreamTime,
			"RequestTime":  v.RequestTime,
			"BytesSent":    v.BytesSent,
		}
		nginxLog.Insert(tags, fields, v.TimeLocal)
	}
}
