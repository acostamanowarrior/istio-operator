package apis

import v1 "github.com/acostamanowarrior/istio-operator/pkg/apis/external/jaeger/v1"

func init() {
	AddToSchemes = append(AddToSchemes,
		v1.SchemeBuilder.AddToScheme,
	)
}
