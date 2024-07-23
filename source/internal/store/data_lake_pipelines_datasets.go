// Copyright 2023 MongoDB Inc
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

// This code was autogenerated at 2023-05-04T11:20:23+01:00. Note: Manual updates are allowed, but may be overwritten.

package store

//go:generate mockgen -destination=../mocks/mock_data_lake_pipelines_datasets.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store PipelineDatasetDeleter

type PipelineDatasetDeleter interface {
	DeletePipelineDataset(string, string, string) error
}

// DeletePipelineDataset encapsulates the logic to manage different cloud providers.
func (s *Store) DeletePipelineDataset(projectID, pipelineName, id string) error {
	_, _, err := s.clientv2.DataLakePipelinesApi.DeletePipelineRunDataset(s.ctx, projectID, pipelineName, id).Execute()
	return err
}
