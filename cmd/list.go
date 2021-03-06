package cmd

import (
	"os"
	"strings"

	"github.com/cnrancher/autok3s/pkg/cluster"
	"github.com/cnrancher/autok3s/pkg/common"
	"github.com/cnrancher/autok3s/pkg/providers"
	"github.com/cnrancher/autok3s/pkg/providers/alibaba"
	"github.com/cnrancher/autok3s/pkg/providers/native"
	"github.com/cnrancher/autok3s/pkg/providers/tencent"
	"github.com/cnrancher/autok3s/pkg/types"
	typesAli "github.com/cnrancher/autok3s/pkg/types/alibaba"
	typesNative "github.com/cnrancher/autok3s/pkg/types/native"
	typesTencent "github.com/cnrancher/autok3s/pkg/types/tencent"
	"github.com/cnrancher/autok3s/pkg/utils"

	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List K3s clusters",
		Example: `  autok3s list`,
	}
)

func ListCommand() *cobra.Command {
	listCmd.Run = func(cmd *cobra.Command, args []string) {
		listCluster()
	}
	return listCmd
}

func listCluster() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeaderLine(false)
	table.SetColumnSeparator("")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Name", "Region", "Provider", "Status", "Masters", "Workers"})

	v := common.CfgPath
	if v == "" {
		logrus.Fatalln("state path is empty")
	}

	clusters, err := utils.ReadYaml(v, common.StateFile)
	if err != nil {
		logrus.Fatalf("read state file error, msg: %v\n", err)
	}

	result, err := cluster.ConvertToClusters(clusters)
	if err != nil {
		logrus.Fatalf("failed to unmarshal state file, msg: %v\n", err)
	}

	var (
		p         providers.Provider
		filters   []*types.Cluster
		removeCtx []string
	)

	// filter useless clusters & contexts.
	for _, r := range result {
		switch r.Provider {
		case "alibaba":
			region := r.Name[strings.LastIndex(r.Name, ".")+1:]

			b, err := yaml.Marshal(r.Options)
			if err != nil {
				logrus.Debugf("failed to convert cluster %s options\n", r.Name)
				removeCtx = append(removeCtx, r.Name)
				continue
			}

			option := &typesAli.Options{}
			if err := yaml.Unmarshal(b, option); err != nil {
				removeCtx = append(removeCtx, r.Name)
				logrus.Debugf("failed to convert cluster %s options\n", r.Name)
				continue
			}
			option.Region = region

			p = &alibaba.Alibaba{
				Metadata: r.Metadata,
				Options:  *option,
			}
		case "tencent":
			region := r.Name[strings.LastIndex(r.Name, ".")+1:]

			b, err := yaml.Marshal(r.Options)
			if err != nil {
				logrus.Debugf("failed to convert cluster %s options\n", r.Name)
				removeCtx = append(removeCtx, r.Name)
				continue
			}

			option := &typesTencent.Options{}
			if err := yaml.Unmarshal(b, option); err != nil {
				removeCtx = append(removeCtx, r.Name)
				logrus.Debugf("failed to convert cluster %s options\n", r.Name)
				continue
			}
			option.Region = region

			p = &tencent.Tencent{
				Metadata: r.Metadata,
				Options:  *option,
			}
		case "native":
			b, err := yaml.Marshal(r.Options)
			if err != nil {
				logrus.Debugf("failed to convert cluster %s options\n", r.Name)
				removeCtx = append(removeCtx, r.Name)
				continue
			}
			option := &typesNative.Options{}
			if err := yaml.Unmarshal(b, option); err != nil {
				removeCtx = append(removeCtx, r.Name)
				logrus.Debugf("failed to convert cluster %s options\n", r.Name)
				continue
			}
			p = &native.Native{
				Metadata: r.Metadata,
				Options:  *option,
				Status:   r.Status,
			}
		}
		if p == nil {
			continue
		}
		isExist, ids, err := p.IsClusterExist()
		if err != nil {
			logrus.Fatalln(err)
		}

		if isExist && len(ids) > 0 {
			filters = append(filters, &types.Cluster{
				Metadata: r.Metadata,
				Options:  r.Options,
				Status:   r.Status,
			})
		} else {
			removeCtx = append(removeCtx, r.Name)
		}
	}

	// remove useless clusters from .state.
	if err := cluster.FilterState(filters); err != nil {
		logrus.Fatalf("failed to remove useless clusters\n")
	}

	// remove useless contexts from kubeCfg.
	for _, r := range removeCtx {
		if err := cluster.OverwriteCfg(r); err != nil {
			logrus.Fatalf("failed to remove useless contexts\n")
		}
	}

	var (
		name   string
		region string
	)
	for _, f := range filters {
		if f.Provider != "native" && strings.Contains(f.Name, ".") {
			name = f.Name[:strings.LastIndex(f.Name, ".")]
			region = f.Name[strings.LastIndex(f.Name, ".")+1:]
		} else {
			name = f.Name
			region = "-"
		}
		table.Append([]string{
			name,
			region,
			f.Provider,
			strings.ToLower(f.Status.Status),
			f.Master,
			f.Worker,
		})
	}

	table.Render()
}
