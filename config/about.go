// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package config

import "fmt"

// Name of the CLI
func (c *Config) Name() string {
	return "Apache CloudStack 🐵 CloudMonkey"
}

// Version CLI
func (c *Config) Version() string {
	return "6.4.0-rc"
}

// PrintHeader prints startup message in CLI mode
func (c *Config) PrintHeader() {
	fmt.Println(c.Name(), c.Version())
	fmt.Println("Report issues: https://github.com/apache/cloudstack-cloudmonkey/issues")
	fmt.Println()
}
