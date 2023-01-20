// Copyright 2021 Google LLC
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

package apis

import (
	"fmt"

	"github.com/apigee/apigeecli/apiclient"
	bundle "github.com/apigee/apigeecli/bundlegen"
	proxybundle "github.com/apigee/apigeecli/bundlegen/proxybundle"
	"github.com/apigee/apigeecli/client/apis"
	"github.com/spf13/cobra"
)

var OasCreateCmd = &cobra.Command{
	Use:     "openapi",
	Aliases: []string{"oas"},
	Short:   "Creates an API proxy from an OpenAPI Specification",
	Long:    "Creates an API proxy from an OpenAPI Specification",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if oasFile == "" && oasURI == "" {
			return fmt.Errorf("either oasfile or oasuri must be passed")
		}
		if targetUrl != "" && targetUrlRef != "" {
			return fmt.Errorf("either target-url or target-url-ref must be passed, not both")
		}
		if createOnline {
			return apiclient.SetApigeeOrg(org)
		} else if importProxy {
			return fmt.Errorf("cannot import offline proxy")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var oasDocName string

		if oasFile != "" {
			oasDocName, _, err = bundle.LoadDocumentFromFile(oasFile, deployConfigFile, validateSpec, formatValidation)
		} else {
			oasDocName, _, err = bundle.LoadDocumentFromURI(oasURI, deployConfigFile, validateSpec, formatValidation)
		}
		if err != nil {
			return err
		}

		//Generate the apiproxy struct
		err, oasContent := bundle.GenerateAPIProxyDefFromOAS(
			name,
			oasDocName,
			skipPolicy,
			addCORS,
			oasGoogleAcessTokenScopeLiteral,
			oasGoogleIdTokenAudLiteral,
			oasGoogleIdTokenAudRef,
			targetUrlRef,
			targetUrl,
		)

		if err != nil {
			return err
		}

		//Create the API proxy bundle
		err = proxybundle.GenerateAPIProxyBundleFromOAS(name,
			string(oasContent),
			skipPolicy,
			addCORS,
			oasGoogleAcessTokenScopeLiteral,
			oasGoogleIdTokenAudLiteral,
			oasGoogleIdTokenAudRef,
			targetUrlRef,
			targetUrl)

		if err != nil {
			return err
		}

		if importProxy {
			_, err = apis.CreateProxy(name, name+".zip")
		}

		return err
	},
}

var oasFile, oasURI, targetUrl, deployConfigFile string
var oasGoogleAcessTokenScopeLiteral, oasGoogleIdTokenAudLiteral, oasGoogleIdTokenAudRef string
var validateSpec, formatValidation, createOnline bool

func init() {
	OasCreateCmd.Flags().StringVarP(&name, "name", "n",
		"", "API Proxy name")
	OasCreateCmd.Flags().StringVarP(&oasFile, "oasfile", "f",
		"", "Open API 3.0 Specification file")
	OasCreateCmd.Flags().StringVarP(&deployConfigFile, "configfile", "",
		"config.json", "Deployment config file")
	OasCreateCmd.Flags().StringVarP(&oasURI, "oasuri", "u",
		"", "Open API 3.0 Specification URI location")
	OasCreateCmd.Flags().StringVarP(&oasGoogleAcessTokenScopeLiteral, "google-accesstoken-scope-literal", "",
		"", "Generate Google Access token with target endpoint and set scope")
	OasCreateCmd.Flags().StringVarP(&oasGoogleIdTokenAudLiteral, "google-idtoken-aud-literal", "",
		"", "Generate Google ID token with target endpoint and set audience")
	OasCreateCmd.Flags().StringVarP(&oasGoogleIdTokenAudRef, "google-idtoken-aud-ref", "",
		"", "Generate Google ID token token with target endpoint and set audience reference")
	OasCreateCmd.Flags().StringVarP(&targetUrlRef, "target-url-ref", "",
		"", "Set a reference variable containing the target endpoint")
	OasCreateCmd.Flags().StringVarP(&targetUrl, "target-url", "",
		"", "Set a target URL for the target endpoint")
	OasCreateCmd.Flags().BoolVarP(&createOnline, "createOnline", "",
		false, "Create API proxy revision online to the configured org and environment")
	OasCreateCmd.Flags().BoolVarP(&importProxy, "import", "",
		false, "Import API Proxy after generation from spec")
	OasCreateCmd.Flags().BoolVarP(&validateSpec, "validate", "",
		true, "Validate Spec before generating proxy")
	OasCreateCmd.Flags().BoolVarP(&skipPolicy, "skip-policy", "",
		false, "Skip adding the OAS Validate policy")
	OasCreateCmd.Flags().BoolVarP(&addCORS, "add-cors", "",
		false, "Add a CORS policy")
	OasCreateCmd.Flags().BoolVarP(&formatValidation, "formatValidation", "",
		true, "disables validation of schema type formats")

	_ = OasCreateCmd.MarkFlagRequired("name")
}
