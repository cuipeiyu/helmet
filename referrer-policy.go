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
	"strings"
)

const HeaderReferrerPolicy = "Referrer-Policy"

type ReferrerPolicyValue string

const (
	ReferrerPolicyNoReferrer                  ReferrerPolicyValue = "no-referrer"
	ReferrerPolicyNoReferrerWhenDowngrade     ReferrerPolicyValue = "no-referrer-when-downgrade"
	ReferrerPolicyOrigin                      ReferrerPolicyValue = "origin"
	ReferrerPolicyOriginWhenCrossOrigin       ReferrerPolicyValue = "origin-when-cross-origin"
	ReferrerPolicySmaeOrigin                  ReferrerPolicyValue = "same-origin"
	ReferrerPolicyStrictOrigin                ReferrerPolicyValue = "strict-origin"
	ReferrerPolicyStrictOriginWhenCrossOrigin ReferrerPolicyValue = "strict-origin-when-cross-origin"
	ReferrerPolicyUnsafeURL                   ReferrerPolicyValue = "unsafe-url"
)

type ReferrerPolicy struct {
	opt *Option

	directives string
}

func NewReferrerPolicy(funs ...OptionFunc) *ReferrerPolicy {
	f := &ReferrerPolicy{
		opt: &Option{
			enable:                   true,
			referrerPolicyDirectives: []ReferrerPolicyValue{},
		},
	}

	for _, fun := range funs {
		fun(f.opt)
	}

	for i, v := range f.opt.referrerPolicyDirectives {
		f.directives += string(v)
		if i < len(f.opt.referrerPolicyDirectives)-1 {
			f.directives += ", "
		}
	}

	f.directives = strings.TrimSpace(f.directives)

	return f
}

var _ http.Handler = (*ReferrerPolicy)(nil)

func (h *ReferrerPolicy) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	if h.opt.enable {
		if h.directives != "" {
			rw.Header().Set(HeaderReferrerPolicy, h.directives)
		}
	}
}
