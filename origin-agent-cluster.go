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

const HeaderOriginAgentCluster = "Origin-Agent-Cluster"

type OriginAgentCluster struct {
	opt *Option
}

func NewOriginAgentCluster(funs ...OptionFunc) *OriginAgentCluster {
	f := &OriginAgentCluster{
		opt: &Option{
			enable: true,
		},
	}

	for _, fun := range funs {
		fun(f.opt)
	}

	return f
}

var _ http.Handler = (*OriginAgentCluster)(nil)

func (h *OriginAgentCluster) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	if h.opt.enable {
		rw.Header().Set(HeaderOriginAgentCluster, "?1")
	}
}
