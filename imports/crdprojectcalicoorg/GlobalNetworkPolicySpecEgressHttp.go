package crdprojectcalicoorg


// HTTP contains match criteria that apply to HTTP requests.
type GlobalNetworkPolicySpecEgressHttp struct {
	// Methods is an optional field that restricts the rule to apply only to HTTP requests that use one of the listed HTTP Methods (e.g. GET, PUT, etc.) Multiple methods are OR'd together.
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// Paths is an optional field that restricts the rule to apply to HTTP requests that use one of the listed HTTP Paths.
	//
	// Multiple paths are OR'd together. e.g: - exact: /foo - prefix: /bar NOTE: Each entry may ONLY specify either a `exact` or a `prefix` match. The validator will check for it.
	Paths *[]*GlobalNetworkPolicySpecEgressHttpPaths `field:"optional" json:"paths" yaml:"paths"`
}

