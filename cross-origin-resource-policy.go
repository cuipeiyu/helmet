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

const HeaderCrossOriginResourcePolicy = "Cross-Origin-Resource-Policy"

type CrossOriginResourcePolicyValue string

const (
	CrossOriginResourcePolicySameOrigin  CrossOriginResourcePolicyValue = "same-origin"
	CrossOriginResourcePolicySameSite    CrossOriginResourcePolicyValue = "same-site"
	CrossOriginResourcePolicyCrossOrigin CrossOriginResourcePolicyValue = "cross-origin"
)

type CrossOriginResourcePolicy struct {
	opt *Option
}

func NewCrossOriginResourcePolicy(funs ...OptionFunc) *CrossOriginResourcePolicy {
	f := &CrossOriginResourcePolicy{
		opt: &Option{
			enable: true,

			crossOriginResourcePolicy: "",
		},
	}

	for _, fun := range funs {
		fun(f.opt)
	}

	return f
}

var _ http.Handler = (*CrossOriginResourcePolicy)(nil)

func (h *CrossOriginResourcePolicy) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	if h.opt.enable && h.opt.crossOriginResourcePolicy != "" {
		rw.Header().Set(HeaderCrossOriginResourcePolicy, string(h.opt.crossOriginResourcePolicy))
	}
}
