//go:build generate
// +build generate

// Copyright 2021 Upbound Inc
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

// Generate OpenAPI server stubs
//go:generate go run -tags generate github.com/deepmap/oapi-codegen/cmd/oapi-codegen -old-config-style -generate types,chi-server,spec -o health/server.gen.go -package api health.yaml

//go:generate go run -tags generate github.com/deepmap/oapi-codegen/cmd/oapi-codegen -old-config-style -generate types,chi-server,spec -o demo/server.gen.go -package api demo.yaml

package api

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen" //nolint:typecheck
)
