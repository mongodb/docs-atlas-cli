// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clusters

import (
	"context"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type WatchOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	name  string
	store store.AtlasClusterDescriber
}

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *WatchOpts) watcher() (bool, error) {
	result, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return false, err
	}
	return result.StateName == "IDLE", nil
}

func (opts *WatchOpts) Run() error {
	if err := opts.Watch(opts.watcher); err != nil {
		return err
	}

	return opts.Print(nil)
}

// mongocli atlas cluster(s) watch <clusterName> [--projectId projectId].
func WatchBuilder() *cobra.Command {
	opts := &WatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch <clusterName>",
		Short: "Watch for a cluster to be available.",
		Long: `This command checks the cluster's status periodically until it reaches an IDLE state. 
Once the cluster reaches the expected state, the command prints "Cluster available."
If you run the command in the terminal, it blocks the terminal session until the resource state changes to IDLE.
You can interrupt the command's polling at any time with CTRL-C.`,
		Example: `$ mongocli atlas cluster watch clusterNameSample`,
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"args":            "clusterName",
			"clusterNameDesc": "Name of the cluster to watch.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), "\nCluster available.\n"),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
