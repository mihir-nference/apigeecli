// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/srinandan/apigeecli/apiclient"
	"github.com/srinandan/apigeecli/client/orgs"
)

func main() {

	//set client options
	apiclient.NewApigeeClient(apiclient.ApigeeClientOptions{
		Org:            "apigee-org-name",
		ServiceAccount: "/Users/srinandans/local_workspace/srinandans-hybrid-orgadmin.json", //"path-to-service-account.json",
		SkipLogInfo:    true,                                                                //skip printing client logs
	})

	//invoke list of orgs
	respBody, err := orgs.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(respBody))
}
