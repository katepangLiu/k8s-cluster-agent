/*
 * Cluster Agent API
 *
 * Cluster Agent API
 *
 * API version: 1.0.0
 * Contact: katepangliu@gmail.com
 */
package swagger

import "k8s.io/kube-openapi/pkg/util/sets"

type NamespacesPodsStringsSet struct {
	Namespaces map[string]PodsStringsSet `json:"namespaces"`
}

type PodsStringsSet struct {
	Pods map[string]sets.String `json:"pods"`
}
