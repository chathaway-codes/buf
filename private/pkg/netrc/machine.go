// Copyright 2020-2021 Buf Technologies, Inc.
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

package netrc

type machine struct {
	name     string
	login    string
	password string
}

func newMachine(
	name string,
	login string,
	password string,
) *machine {
	return &machine{
		name:     name,
		login:    login,
		password: password,
	}
}

func (m *machine) Name() string {
	return m.name
}

func (m *machine) Login() string {
	return m.login
}

func (m *machine) Password() string {
	return m.password
}
