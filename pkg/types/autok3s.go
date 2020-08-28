package types

type AutoK3s struct {
	Clusters []Cluster `json:"clusters" yaml:"clusters"`
}

type Cluster struct {
	Metadata `json:",inline" mapstructure:",squash"`
	Options  interface{} `json:"options,omitempty"`

	Status `json:"status" yaml:"status"`
}

type Metadata struct {
	Name            string `json:"name" yaml:"name"`
	Provider        string `json:"provider" yaml:"provider"`
	Master          string `json:"master" yaml:"master"`
	Worker          string `json:"worker" yaml:"worker"`
	Token           string `json:"token,omitempty" yaml:"token,omitempty"`
	UI              string `json:"ui,omitempty" yaml:"ui,omitempty"`
	Repo            string `json:"repo,omitempty" yaml:"repo,omitempty"`
	ClusterCIDR     string `json:"cluster-cidr,omitempty" yaml:"cluster-cidr,omitempty"`
	MasterExtraArgs string `json:"master-extra-args,omitempty" yaml:"master-extra-args,omitempty"`
	WorkerExtraArgs string `json:"worker-extra-args,omitempty" yaml:"worker-extra-args,omitempty"`
}

type Status struct {
	MasterNodes []Node `json:"master-nodes,omitempty"`
	WorkerNodes []Node `json:"worker-nodes,omitempty"`
}

type Node struct {
	SSH `json:",inline"`

	Master            bool     `json:"master,omitempty" yaml:"master,omitempty"`
	InstanceID        string   `json:"instance-id,omitempty" yaml:"instance-id,omitempty"`
	InstanceStatus    string   `json:"instance-status,omitempty" yaml:"instance-status,omitempty"`
	PublicIPAddress   []string `json:"public-ip-address,omitempty" yaml:"public-ip-address,omitempty"`
	InternalIPAddress []string `json:"internal-ip-address,omitempty" yaml:"internal-ip-address,omitempty"`
}

type SSH struct {
	Port   string `json:"ssh-port,omitempty" yaml:"ssh-port,omitempty"`
	User   string `json:"user,omitempty" yaml:"user,omitempty"`
	SSHKey string `json:"ssh-key,omitempty" yaml:"ssh-key,omitempty"`
}

type Flag struct {
	Name      string
	P         *string
	V         string
	ShortHand string
	Usage     string
	Required  bool
}
