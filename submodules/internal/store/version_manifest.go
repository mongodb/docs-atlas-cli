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

package store

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_version_manifest.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store VersionManifestUpdater,VersionManifestGetter,VersionManifestUpdaterServiceVersionDescriber

type VersionManifestUpdater interface {
	UpdateVersionManifest(*opsmngr.VersionManifest) (*opsmngr.VersionManifest, error)
}

type VersionManifestGetter interface {
	GetVersionManifest(string) (*opsmngr.VersionManifest, error)
}

type VersionManifestUpdaterServiceVersionDescriber interface {
	VersionManifestUpdater
	ServiceVersionDescriber
}

// UpdateVersionManifest encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateVersionManifest(versionManifest *opsmngr.VersionManifest) (*opsmngr.VersionManifest, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).VersionManifest.Update(s.ctx, versionManifest)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// GetVersionManifest encapsulates the logic to manage different cloud providers.
func (s *Store) GetVersionManifest(version string) (*opsmngr.VersionManifest, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).VersionManifest.Get(s.ctx, version)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
