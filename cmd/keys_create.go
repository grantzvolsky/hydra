/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ory/hydra/cmd/cli"
)

// createCmd represents the create command
func NewKeysCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <set> <key>",
		Short: "Create a new JSON Web Key Set",
		Run:   cli.NewHandler().Keys.CreateKeys,
	}
	cmd.Flags().StringP("alg", "a", "RS256", "The algorithm to be used to generated they key. Supports: RS256, ES512, HS256, EdDSA")
	cmd.Flags().StringP("use", "u", "sig", "The intended use of this key")
	return cmd
}
