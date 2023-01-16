/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"k8s.io/klog/v2"

	nfdtopologygarbagecollector "sigs.k8s.io/node-feature-discovery/pkg/nfd-topology-gc"
	"sigs.k8s.io/node-feature-discovery/pkg/version"
)

const (
	// ProgramName is the canonical name of this program
	ProgramName = "nfd-topology-gc"
)

func main() {
	flags := flag.NewFlagSet(ProgramName, flag.ExitOnError)

	printVersion := flags.Bool("version", false, "Print version and exit.")

	args := parseArgs(flags, os.Args[1:]...)

	if *printVersion {
		fmt.Println(ProgramName, version.Get())
		os.Exit(0)
	}

	// Assert that the version is known
	if version.Undefined() {
		klog.Warningf("version not set! Set -ldflags \"-X sigs.k8s.io/node-feature-discovery/pkg/version.version=`git describe --tags --dirty --always`\" during build or run.")
	}

	// Get new TopologyGC instance
	gc, err := nfdtopologygarbagecollector.New(args)
	if err != nil {
		klog.Exit(err)
	}

	if err = gc.Run(); err != nil {
		klog.Exit(err)
	}
}

func parseArgs(flags *flag.FlagSet, osArgs ...string) *nfdtopologygarbagecollector.Args {
	args := initFlags(flags)

	_ = flags.Parse(osArgs)
	if len(flags.Args()) > 0 {
		fmt.Fprintf(flags.Output(), "unknown command line argument: %s\n", flags.Args()[0])
		flags.Usage()
		os.Exit(2)
	}

	return args
}

func initFlags(flagset *flag.FlagSet) *nfdtopologygarbagecollector.Args {
	args := &nfdtopologygarbagecollector.Args{}

	flagset.DurationVar(&args.GCPeriod, "gc-interval", time.Duration(1)*time.Hour,
		"Interval between which Garbage Collector will try to cleanup any missed but already obsolete NodeResourceTopology. [Default: 1h]")
	flagset.StringVar(&args.Kubeconfig, "kubeconfig", "",
		"Kubeconfig to use")

	klog.InitFlags(flagset)

	return args
}
