// # Copyright Meshery Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package system

import (
	// "fmt"

	// "github.com/layer5io/meshery/mesheryctl/internal/cli/root/config"
	// "github.com/layer5io/meshery/mesheryctl/pkg/utils"
	// "github.com/pkg/errors"

	// log "github.com/sirupsen/logrus"

	"fmt"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var (
	system string
	credential string
	sheetId string
	modelsOutput string
	imgsOutput string
)

// PublishCmd represents the publish command to publish Meshery Models to Websites, Remote Provider, Meshery
var PublishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish Meshery Models to Websites, Remote Provider, Meshery",
	Long:  `Publishes metadata about Meshery Models to Websites, Remote Provider, Meshery by reading from a Google Spreadsheet.`,
	Example: `
// Publish To Meshery
mesheryctl registry publish --system=meshery --credential=$GoogleCredential --sheetId=$GoogleSheetID --models-output=../../server/meshmodel

// Publish To Remote Provider
mesheryctl registry publish --system=remote-provider --credential=$GoogleCredential --sheetId=$GoogleSheetID --models-output=<remote-provider>/meshmodels/models --imgs-output=<remote-provider>/ui/public/img/meshmodels

// Publish To Website
mesheryctl registry publish --system=website --credential=$GoogleCredential --sheetId=$GoogleSheetID --models-path=<repo>/integrations --imgs-output=<remote-provider>/ui/public/img/meshmodels
	`,
	// PreRunE: func(cmd *cobra.Command, args []string) error {
	// 	//Check prerequisite
	// },
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("args:", args)
		return nil
	},
}

func init() {
	PublishCmd.Flags().StringVarP(&system, "system", "sys", "", "System to publish to. Available options: meshery, remote-provider, website")
	PublishCmd.Flags().StringVarP(&credential, "credential", "cred", "", "Google Credential File")
	PublishCmd.Flags().StringVarP(&sheetId, "sheetId", "", "", "Google Sheet ID to read from")
	PublishCmd.Flags().StringVarP(&modelsOutput, "models-output", "", "", "Path to Meshery Models output")
	PublishCmd.Flags().StringVarP(&imgsOutput, "imgs-output", "g", "", "Path to Meshery Models Images output")
}