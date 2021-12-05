package memberroll

import (
	"context"

	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"

	"github.com/acostamanowarrior/istio-operator/pkg/controller/common"
)

var ctx = common.NewContextWithLog(context.Background(), logf.Log)
