package amqonline

import (
	"fmt"

	l "github.com/integr8ly/integreatly-operator/pkg/resources/logger"

	"github.com/integr8ly/integreatly-operator/pkg/resources"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *Reconciler) newAlertsReconciler(logger l.Logger, installType string) resources.AlertReconciler {
	installationName := resources.InstallationNames[installType]
	keycloakServicePortCount := 2

	return &resources.AlertReconcilerImpl{
		ProductName:  "AMQ Online",
		Installation: r.inst,
		Log:          logger,
		Alerts: []resources.AlertConfiguration{
			{
				AlertName: "rhmi-amq-online-slo",
				GroupName: "amqonline.rules",
				Namespace: r.Config.GetNamespace(),
				Rules: []monitoringv1.Rule{
					{
						Alert: "AMQOnlineConsoleAvailable",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlAlertsAndTroubleshooting,
							"message": fmt.Sprintf("AMQ-SLO-1.1: AMQ Online console is not available in namespace '%s'", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("absent(kube_endpoint_address_available{endpoint='console',namespace='%s'}==2)", r.Config.GetNamespace())),
						For:    "300s",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "AMQOnlineKeycloakAvailable",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlAlertsAndTroubleshooting,
							"message": fmt.Sprintf("AMQ-SLO-1.4: Keycloak is not available in namespace %s", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("absent(kube_endpoint_address_available{endpoint='standard-authservice',namespace='%s'}==%v)", r.Config.GetNamespace(), keycloakServicePortCount)),
						For:    "300s",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "AMQOnlineOperatorAvailable",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlAlertsAndTroubleshooting,
							"message": fmt.Sprintf("AMQ-SLO-1.5: amq-online(enmasse) operator is not available in namespace %s", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("absent(kube_pod_status_ready{condition='true',namespace='%s',pod=~'enmasse-operator-.*'}==1)", r.Config.GetNamespace())),
						For:    "300s",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
				},
			},

			{
				AlertName: "ksm-endpoint-alerts",
				GroupName: "amqonline-endpoint.rules",
				Namespace: r.Config.GetNamespace(),
				Rules: []monitoringv1.Rule{
					{
						Alert: "RHMIAMQOnlineNoneAuthServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='none-authservice', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "RHMIAMQOnlineAddressSpaceControllerServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='address-space-controller', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "RHMIAMQOnlineConsoleServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='console', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "RHMIAMQOnlineRegistryCsServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='rhmi-registry-cs', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "warning", "product": installationName},
					},
					{
						Alert: "RHMIAMQOnlineStandardAuthServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='standard-authservice', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "RHMIAMQOnlineEnmasseOperatorMetricsServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='enmasse-operator-metrics', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
				},
			},

			{
				AlertName: "ksm-amqonline-alerts",
				GroupName: "general.rules",
				Namespace: r.Config.GetNamespace(),
				Rules: []monitoringv1.Rule{
					{
						Alert: "AMQOnlinePodCount",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlAlertsAndTroubleshooting,
							"message": "Pod count for namespace {{  $labels.namespace  }} is {{  $value }}. Expected at least 2 pods.",
						},
						Expr:   intstr.FromString(fmt.Sprintf("(1-absent(kube_pod_status_ready{condition='true', namespace='%[1]v'})) or sum(kube_pod_status_ready{condition='true', namespace='%[1]v'}) < 2", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "AMQOnlineContainerHighMemory",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlAlertsAndTroubleshooting,
							"message": "The {{  $labels.container  }} Container in the {{ $labels.pod  }} Pod has been using {{  $value }}% of available memory for longer than 15 minutes.",
						},
						Expr:   intstr.FromString(fmt.Sprintf("sum by(container, pod) (container_memory_usage_bytes{container!='',namespace='%[1]v'}) / sum by(container, pod) (kube_pod_container_resource_requests{namespace='%[1]v', resource='memory'}) * 100 > 90", r.Config.GetNamespace())),
						For:    "15m",
						Labels: map[string]string{"severity": "warning", "product": installationName},
					},
				},
			},
		},
	}
}
