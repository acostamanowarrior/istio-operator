package apis

import "github.com/acostamanowarrior/istio-operator/pkg/apis/external/kiali/v1alpha1"

func init() {
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
	)
}
