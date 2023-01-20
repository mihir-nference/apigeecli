package config

import "fmt"

const (
	OpenAPIRequestValidationPolicy       = "OpenAPI-Spec-Validation-1"
	OpenAPIResponseValidationPolicy      = "OpenAPI-Spec-Response-Validation-1"
	AssignPrivacyPreservedHeaderPolicy   = "Privacy-Preserved-Rule"
	OASOrPrivacyPreservedDataFaultPolicy = "RF-OAS-Or-Privacy-Preserved-Data"
	JSCErrorHandlePolicy                 = "JS-OAS-Or-Privacy-Preserved-Error"
)

const (
	OASResponseErrorCodeVariable   = "oas_resp_error_code"
	OASResponseFaultStringVariable = "oas_resp_fault_string"
	JSResourceFileName             = "oas_response_fault.js"
	OASFileName                    = "openapi3.json"
)

func GetOASFaultNotation(property string) string {
	return fmt.Sprintf("OASValidation.%s.fault.%s", OpenAPIResponseValidationPolicy, property)
}

func GetOASFailCondition() string {
	return fmt.Sprintf("OASValidation.%s.failed = true", OpenAPIResponseValidationPolicy)
}
