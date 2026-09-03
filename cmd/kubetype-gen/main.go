// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	goflag "flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	gengo "k8s.io/gengo/v2"
	"k8s.io/klog/v2"

	"istio.io/tools/cmd/kubetype-gen/generators"
	"istio.io/tools/cmd/kubetype-gen/scanner"
)

func main() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	s := &scanner.Scanner{}
	s.AddFlags(pflag.CommandLine)

	var inputDirs string
	pflag.StringVarP(&inputDirs, "input-dirs", "i", "", "Comma-separated list of import paths to scan")
	pflag.Parse()

	var inputs []string
	for _, dir := range strings.Split(inputDirs, ",") {
		if dir = strings.TrimSpace(dir); dir != "" {
			inputs = append(inputs, dir)
		}
	}

	if err := gengo.Execute(
		generators.NameSystems("", nil),
		generators.DefaultNameSystem(),
		s.Targets,
		gengo.StdBuildTag,
		inputs,
	); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	klog.V(2).Info("Completed successfully.")
}
