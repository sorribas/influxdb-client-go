// Copyright 2020-2021 InfluxData, Inc. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package influxdb2

import (
	"fmt"
	"runtime"

	"github.com/influxdata/influxdb-client-go/v2/internal/http"
)

const (
	Version = "2.2.3"
)

func init() {
	http.UserAgent = fmt.Sprintf("influxdb-client-go/%s  (%s; %s)", Version, runtime.GOOS, runtime.GOARCH)
}
