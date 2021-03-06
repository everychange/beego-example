// Copyright 2020 beego-dev
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"time"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	// use SetTimeout(connectTimeout, readWriteTimeout)
	resp, err := httplib.Get("http://beego.me/").SetTimeout(100*time.Second, 30*time.Second).Response()
	if err != nil {
		logs.Error(err)
	}

	logs.Info(resp)
}
