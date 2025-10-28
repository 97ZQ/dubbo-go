/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package global

import (
	"dubbo.apache.org/dubbo-go/v3/common"
)

type RouterConfig struct {
	Scope      string   `validate:"required" yaml:"scope" json:"scope,omitempty" property:"scope"`
	Key        string   `validate:"required" yaml:"key" json:"key,omitempty" property:"key"`
	Force      *bool    `default:"false" yaml:"force" json:"force,omitempty" property:"force"`
	Runtime    *bool    `default:"false" yaml:"runtime" json:"runtime,omitempty" property:"runtime"`
	Enabled    *bool    `default:"true" yaml:"enabled" json:"enabled,omitempty" property:"enabled"`
	Valid      *bool    `default:"true" yaml:"valid" json:"valid,omitempty" property:"valid"`
	Priority   int      `default:"0" yaml:"priority" json:"priority,omitempty" property:"priority"`
	Conditions []string `yaml:"conditions" json:"conditions,omitempty" property:"conditions"`
	Tags       []Tag    `yaml:"tags" json:"tags,omitempty" property:"tags"`
	ScriptType string   `yaml:"type" json:"type,omitempty" property:"type"`
	Script     string   `yaml:"script" json:"script,omitempty" property:"script"`
}

type Tag struct {
	Name      string               `yaml:"name" json:"name,omitempty" property:"name"`
	Match     []*common.ParamMatch `yaml:"match" json:"match,omitempty" property:"match"`
	Addresses []string             `yaml:"addresses" json:"addresses,omitempty" property:"addresses"`
}

func DefaultRouterConfig() *RouterConfig {
	return &RouterConfig{
		Conditions: make([]string, 0),
		Tags:       make([]Tag, 0),
	}
}

func (c *RouterConfig) Clone() *RouterConfig {
	if c == nil {
		return nil
	}

	var newForce *bool
	if c.Force != nil {
		newForce = new(bool)
		*newForce = *c.Force
	}

	var newRuntime *bool
	if c.Runtime != nil {
		newRuntime = new(bool)
		*newRuntime = *c.Runtime
	}

	var newEnabled *bool
	if c.Enabled != nil {
		newEnabled = new(bool)
		*newEnabled = *c.Enabled
	}

	var newValid *bool
	if c.Valid != nil {
		newValid = new(bool)
		*newValid = *c.Valid
	}

	newConditions := make([]string, len(c.Conditions))
	copy(newConditions, c.Conditions)

	newTags := make([]Tag, len(c.Tags))
	copy(newTags, c.Tags)

	return &RouterConfig{
		Scope:      c.Scope,
		Key:        c.Key,
		Force:      newForce,
		Runtime:    newRuntime,
		Enabled:    newEnabled,
		Valid:      newValid,
		Priority:   c.Priority,
		Conditions: newConditions,
		Tags:       newTags,
		ScriptType: c.ScriptType,
		Script:     c.Script,
	}
}
