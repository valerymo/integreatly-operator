package marin3r

import (
	"fmt"
	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/apis/v1alpha1"
	l "github.com/integr8ly/integreatly-operator/pkg/resources/logger"

	"github.com/integr8ly/integreatly-operator/pkg/products/marin3r/config"
	"github.com/integr8ly/integreatly-operator/pkg/resources"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const rejectedRequestsAlertExpr = "abs(clamp_min(increase(limited_calls[1m]) - %f, 0) / (sum(increase(authorized_calls[1m])) + sum(increase(limited_calls[1m]))) - (increase(limited_calls[1m]) / (sum(increase(authorized_calls[1m])) + sum(increase(limited_calls[1m]))))) > 0.3"

func (r *Reconciler) newRejectedRequestsAlertsReconciler(logger l.Logger, installType string) (resources.AlertReconciler, error) {
	installationName := resources.InstallationNames[installType]

	namespace := r.Config.GetNamespace()
	alertName := "rejected-requests"

	if integreatlyv1alpha1.IsRHOAM(integreatlyv1alpha1.InstallationType(installType)) {
		observabilityConfig, err := r.ConfigManager.ReadObservability()
		if err != nil {
			logger.Warning("failed to get observability config")
			return nil, err
		}

		namespace = observabilityConfig.GetNamespace()
		alertName = "marin3r-rejected-requests"
	}

	limitPerMinute, err := config.ConvertRate(
		r.RateLimitConfig.Unit,
		config.Minute,
		int(r.RateLimitConfig.RequestsPerUnit),
	)
	if err != nil {
		return nil, err
	}

	return &resources.AlertReconcilerImpl{
		ProductName:  "3Scale",
		Installation: r.installation,
		Log:          logger,
		Alerts: []resources.AlertConfiguration{
			{
				AlertName: alertName,
				GroupName: "rejected-requests.rules",
				Namespace: namespace,
				Rules: []monitoringv1.Rule{
					{
						Alert: "RHOAMApiUsageRejectedRequestsMismatch",
						Annotations: map[string]string{
							"message": "The volume of rejected requests doesn't match the expected volume given the incoming requests and the configuration",
						},
						Expr:   intstr.FromString(fmt.Sprintf(rejectedRequestsAlertExpr, limitPerMinute)),
						Labels: map[string]string{"severity": "info", "product": installationName},
						For:    "30s",
					},
				},
			},
		},
	}, nil
}
