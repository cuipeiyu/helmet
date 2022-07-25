// Copyright 2022 cuipeiyu (xiaocui1023@vip.qq.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helmet

import (
	"net/http"
	"strconv"
)

const HeaderExpectCT = "Expect-CT"

type ExpectCT struct {
	opt *Option

	value string
}

func NewExpectCT(funs ...OptionFunc) *ExpectCT {
	f := &ExpectCT{
		opt: &Option{
			enable:            true,
			expectCTMaxAge:    0,
			expectCTEnforce:   false,
			expectCTReportUri: "",
		},
	}

	for _, fun := range funs {
		fun(f.opt)
	}

	maxAge := int(f.opt.expectCTMaxAge.Seconds())
	if maxAge >= 0 {
		f.value = "max-age=" + strconv.Itoa(maxAge)

		if f.opt.expectCTEnforce {
			f.value += "; enforce"
		}

		if f.opt.expectCTReportUri != "" {
			f.value += `; report-uri="` + f.opt.expectCTReportUri + `"`
		}
	}

	return f
}

var _ http.Handler = (*ExpectCT)(nil)

func (h *ExpectCT) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	if h.opt.enable {
		rw.Header().Set(HeaderExpectCT, h.value)
	}
}
