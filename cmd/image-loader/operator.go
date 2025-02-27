/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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
	"fmt"
	"io/ioutil"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/controller/operator/defaults"
	"k8c.io/kubermatic/v2/pkg/semver"
	kubermaticversion "k8c.io/kubermatic/v2/pkg/version"
)

func loadKubermaticConfiguration(log *zap.SugaredLogger, filename string) (*kubermaticv1.KubermaticConfiguration, error) {
	log.Infow("Loading KubermaticConfiguration", "file", filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	config := &kubermaticv1.KubermaticConfiguration{}
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, fmt.Errorf("failed to parse file as YAML: %w", err)
	}

	defaulted, err := defaults.DefaultConfiguration(config, log)
	if err != nil {
		return nil, fmt.Errorf("failed to process: %w", err)
	}

	return defaulted, nil
}

func getVersionsFromKubermaticConfiguration(config *kubermaticv1.KubermaticConfiguration) []*kubermaticversion.Version {
	versions := []*kubermaticversion.Version{}

	assembleVersions := func(kind string, configuredVersions []semver.Semver) {
		for i := range configuredVersions {
			versions = append(versions, &kubermaticversion.Version{
				Version: configuredVersions[i].Semver(),
				Type:    kind,
			})
		}
	}

	assembleVersions("kubernetes", config.Spec.Versions.Versions)
	return versions
}
