/*
 * Cluster Agent API
 *
 * Cluster Agent API
 *
 * API version: 1.0.0
 * Contact: you@your-company.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Pod struct {

	PodName string `json:"podName"`

	Services []string `json:"services,omitempty"`
}