package fuse

import (
	"fmt"

	"github.com/integr8ly/integreatly-operator/pkg/resources"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *Reconciler) newAlertsReconciler(installType string) resources.AlertReconciler {
	installationName := resources.InstallationNames[installType]

	return &resources.AlertReconcilerImpl{
		ProductName:  "Fuse",
		Installation: r.installation,
		Log:          r.log,
		Alerts: []resources.AlertConfiguration{
			{
				AlertName: "ksm-endpoint-alerts",
				Namespace: r.Config.GetNamespace(),
				GroupName: "fuse-online-endpoint.rules",
				Rules: []monitoringv1.Rule{
					{
						Alert: "RHMIFuseOnlineBrokerAmqTcpServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='broker-amq-tcp', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "RHMIFuseOnlineSyndesisMetaServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='syndesis-meta', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "RHMIFuseOnlineSyndesisOauthproxyServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='syndesis-oauthproxy', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "RHMIFuseOnlineSyndesisPrometheusServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='syndesis-prometheus', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "RHMIFuseOnlineSyndesisServerServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='syndesis-server', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "RHMIFuseOnlineSyndesisUiServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No endpoints available for the {{  $labels.endpoint  }} service in the %s namespace", r.Config.GetNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='syndesis-ui', namespace='%s'} < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
				},
			},

			{
				AlertName: "ksm-endpoint-alerts",
				Namespace: r.Config.GetOperatorNamespace(),
				GroupName: "fuse-online-operator-endpoint.rules",
				Rules: []monitoringv1.Rule{
					{
						Alert: "RHMIFuseOnlineOperatorRhmiRegistryCsServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetOperatorNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='rhmi-registry-cs', namespace='%s'} < 1", r.Config.GetOperatorNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "warning", "product": installationName},
					},
					{
						Alert: "RHMIFuseOnlineOperatorSyndesisOperatorMetricsServiceEndpointDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlEndpointAvailableAlert,
							"message": fmt.Sprintf("No {{  $labels.endpoint  }} endpoints in namespace %s. Expected at least 1.", r.Config.GetOperatorNamespace()),
						},
						Expr:   intstr.FromString(fmt.Sprintf("kube_endpoint_address_available{endpoint='syndesis-operator-metrics', namespace='%s'} < 1", r.Config.GetOperatorNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
				},
			},

			{
				AlertName: "ksm-fuse-online-alerts",
				Namespace: r.Config.GetNamespace(),
				GroupName: "general.rules",
				Rules: []monitoringv1.Rule{
					{
						Alert: "FuseOnlineSyndesisServerInstanceDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlAlertsAndTroubleshooting,
							"message": "Fuse Online Syndesis Server instance {{ $labels.pod }} in namespace {{ $labels.namespace }} is down.",
						},
						Expr:   intstr.FromString(fmt.Sprintf("(1 - absent(kube_pod_status_ready{condition='true',namespace='%[1]v'} * on (pod, namespace) kube_pod_labels{label_deploymentconfig='syndesis-server'})) or sum(kube_pod_status_ready{condition='true',namespace='%[1]v'} * on (pod, namespace) kube_pod_labels{label_deploymentconfig='syndesis-server'}) < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
					{
						Alert: "FuseOnlineSyndesisUIInstanceDown",
						Annotations: map[string]string{
							"sop_url": resources.SopUrlAlertsAndTroubleshooting,
							"message": " Fuse Online Syndesis UI instance {{ $labels.pod }} in namespace {{ $labels.namespace }} is down.",
						},
						Expr:   intstr.FromString(fmt.Sprintf("(1 - absent(kube_pod_status_ready{condition='true',namespace='%[1]v'} * on (pod, namespace) kube_pod_labels{label_deploymentconfig='syndesis-ui'})) or sum(kube_pod_status_ready{condition='true',namespace='%[1]v'} * on (pod, namespace) kube_pod_labels{label_deploymentconfig='syndesis-ui'}) < 1", r.Config.GetNamespace())),
						For:    "5m",
						Labels: map[string]string{"severity": "critical", "product": installationName},
					},
				},
			},
		},
	}
}
