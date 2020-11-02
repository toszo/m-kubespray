package inventory

type Node struct {
	Name        string
	IpAddress   string
	AnsibleHost string
}

type InventoryWrapper struct {
	Inventory Inventory `yaml:"all"`
}

type Inventory struct {
	Vars     InventoryVars     `yaml:"vars"`
	Hosts    map[string]Node   `yaml:"hosts"`
	Children InventoryChildren `yaml:"children"`
}

type InventoryVars struct {
	User        string `yaml:"ansible_user"`
	KeyFilePath string `yaml:"ansible_ssh_private_key_file"`
}

type InventoryChildren struct {
	KubeMaster HostsList    `yaml:"kube-master"`
	KubeNode   HostsList    `yaml:"kube-node"`
	Etcd       HostsList    `yaml:"etcd"`
	K8sCluster ChildrenList `yaml:"k8s-cluster"`
	//TODO another cluster components
}

type ChildrenList struct {
	Hosts []string `yaml:"children"`
}

type HostsList struct {
	Hosts []string `yaml:"hosts"`
}
