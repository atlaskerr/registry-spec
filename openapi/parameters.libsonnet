// Copyright © 2019 Titan Distribution Authors
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

local namespace = {
  name: 'namespace',
  'in': 'path',
  required: true,
  schema: { type: 'string' },
};

local project = {
  name: 'project',
  'in': 'path',
  required: true,
  schema: { type: 'string' },
};

local repo = {
  name: 'repo',
  'in': 'path',
  required: true,
  schema: { type: 'string' },
};

local package = {
  name: 'package',
  'in': 'path',
  required: true,
  schema: { type: 'string' },
};

local label = {
  name: 'label',
  'in': 'query',
  schema: { type: 'string' },
};

local perPage = {
  name: 'perPage',
  'in': 'query',
  schema: { type: 'string' },
};

{
  namespace:: namespace,
  project:: project,
  repo:: repo,
  label:: label,
  perPage:: perPage,
}
