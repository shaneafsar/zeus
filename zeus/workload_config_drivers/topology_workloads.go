package zeus_topology_config_drivers

import "github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"

type TopologyConfigDriver struct {
	*IngressDriver
	*StatefulSetDriver
	*ServiceDriver
	*DeploymentDriver
	*ServiceMonitorDriver
	*ConfigMapDriver
}

func (t *TopologyConfigDriver) SetCustomConfig(inf *topology_workloads.TopologyBaseInfraWorkload) {
	if inf.Ingress != nil && t.IngressDriver != nil {
		t.SetIngressConfigs(inf.Ingress)
	}
	if inf.StatefulSet != nil && t.StatefulSetDriver != nil {
		t.SetStatefulSetConfigs(inf.StatefulSet)
	}
	if inf.Deployment != nil && t.DeploymentDriver != nil {
		t.SetDeploymentConfigs(inf.Deployment)
	}
	if inf.Service != nil && t.ServiceDriver != nil {
		t.SetServiceConfigs(inf.Service)
	}
	if inf.ServiceMonitor != nil && t.ServiceMonitorDriver != nil {
		t.SetServiceMonitorConfigs(inf.ServiceMonitor)
	}
	if inf.ConfigMap != nil && t.ConfigMapDriver != nil {
		t.SetConfigMaps(inf.ConfigMap)
	}
}
