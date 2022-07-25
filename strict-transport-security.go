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
	"time"
)

const HeaderStrictTransportSecurity = "Strict-Transport-Security"

type StrictTransportSecurity struct {
	opt *Option

	value string
}

func NewStrictTransportSecurity(funs ...OptionFunc) *StrictTransportSecurity {
	f := &StrictTransportSecurity{
		opt: &Option{
			enable: true,

			strictTransportSecurityMaxAge:            180 * 24 * time.Hour,
			strictTransportSecurityIncludeSubDomains: false,
			strictTransportSecurityPreload:           false,
		},
	}

	for _, fun := range funs {
		fun(f.opt)
	}

	maxAge := int(f.opt.strictTransportSecurityMaxAge.Seconds())
	if maxAge > 0 {
		f.value = "max-age=" + strconv.Itoa(maxAge)

		if f.opt.strictTransportSecurityIncludeSubDomains {
			f.value += "; includeSubDomains"
		}

		if f.opt.strictTransportSecurityPreload {
			f.value += "; preload"
		}
	}

	return f
}

var _ http.Handler = (*StrictTransportSecurity)(nil)

func (h *StrictTransportSecurity) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	if h.opt.enable {
		rw.Header().Set(HeaderStrictTransportSecurity, h.value)
	}
}
