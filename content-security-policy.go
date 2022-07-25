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

import "net/http"

const (
	HeaderContentSecurityPolicy           = "Content-Security-Policy"
	HeaderContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"
)

const DefaultContentSecurityPolicyDirectives = `default-src 'self'; object-src 'none'; child-src https:;`

type ContentSecurityPolicy struct {
	opt *Option

	headerKey string
}

func NewContentSecurityPolicy(funs ...OptionFunc) *ContentSecurityPolicy {
	f := &ContentSecurityPolicy{
		opt: &Option{
			enable: true,

			contentSecurityPolicyReportOnly: false,
			contentSecurityPolicyDirectives: DefaultContentSecurityPolicyDirectives,
		},
	}

	for _, fun := range funs {
		fun(f.opt)
	}

	f.headerKey = HeaderContentSecurityPolicy
	if f.opt.contentSecurityPolicyReportOnly {
		f.headerKey = HeaderContentSecurityPolicyReportOnly
	}

	return f
}

var _ http.Handler = (*ContentSecurityPolicy)(nil)

func (h *ContentSecurityPolicy) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	if h.opt.enable {
		rw.Header().Set(h.headerKey, h.opt.contentSecurityPolicyDirectives)
	}
}
