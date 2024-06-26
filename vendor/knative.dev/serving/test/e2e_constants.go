/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// All test-affecting constants should be placed in this file
// At some point it may make sense to be able to modify them
// via a configuration mechanism (see https://github.com/knative/serving/issues/6109)

package test

const (
	// Environment propagation conformance test objects

	// ConformanceConfigMap is the name of the configmap to propagate env variables from
	ConformanceConfigMap = "conformance-test-configmap"

	// ConformanceSecret is the name of the secret to propagate env variables from
	ConformanceSecret = "conformance-test-secret"

	// EnvKey is the configmap/secret key which contains test value
	EnvKey = "testKey"

	// EnvValue is the configmap/secret test value to match env variable with
	EnvValue = "testValue"

	// testAnnotation is an annotation attached to resources originating from tests.
	testAnnotation = "knative-e2e-test"

	// GatewayNamespaceOverride is the name of env variable with gateway namespace
	GatewayNamespaceOverride = "GATEWAY_NAMESPACE_OVERRIDE"
)
