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

const HeaderCrossOriginEmbedderPolicy = "Cross-Origin-Embedder-Policy"

type CrossOriginEmbedderPolicyValue string

const (
	CrossOriginEmbedderPolicyRequireCorp    CrossOriginEmbedderPolicyValue = "require-corp"
	CrossOriginEmbedderPolicyCredentialless CrossOriginEmbedderPolicyValue = "credentialless"
)

type CrossOriginEmbedderPolicy struct {
	opt *Option
}

func NewCrossOriginEmbedderPolicy(funs ...OptionFunc) *CrossOriginEmbedderPolicy {
	f := &CrossOriginEmbedderPolicy{
		opt: &Option{
			enable: true,

			crossOriginEmbedderPolicy: "",
		},
	}

	for _, fun := range funs {
		fun(f.opt)
	}

	return f
}

var _ http.Handler = (*CrossOriginEmbedderPolicy)(nil)

func (h *CrossOriginEmbedderPolicy) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	if h.opt.enable && h.opt.crossOriginEmbedderPolicy != "" {
		rw.Header().Set(HeaderCrossOriginEmbedderPolicy, string(h.opt.crossOriginEmbedderPolicy))
	}
}
