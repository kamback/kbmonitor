package nginxLog

import (
	"adumonitor/config"
	"github.com/influxdata/influxdb/client/v2"
	"time"
)

var ifdb config.InfluxdbConf

func init() {
	ifdb = config.Config.Ifdb
	ifdb.Table = "nginx_log"

}

func Insert(tags map[string]string, fields map[string]interface{}, t time.Time) {
	cl, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     ifdb.Addr,
		Username: ifdb.Username,
		Password: ifdb.Password,
	})
	if err != nil {
		panic(err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  ifdb.Database,
		Precision: ifdb.Precision,
	})
	if err != nil {
		panic(err)
	}

	/*tags := map[string]string{"Path": v.Path, "Method": v.Method, "Scheme": v.Scheme, "Status": v.Status}
	fields := map[string]interface{}{
		"UpstreamTime": v.UpstreamTime,
		"RequestTime":  v.RequestTime,
		"BytesSent":    v.BytesSent,
	}*/

	pt, err := client.NewPoint(ifdb.Table, tags, fields, t)
	if err != nil {
		panic(err)
	}
	bp.AddPoint(pt)

	if err := cl.Write(bp); err != nil {
		panic(err)
	}

	return
}
