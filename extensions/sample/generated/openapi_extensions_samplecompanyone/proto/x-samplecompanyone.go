// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED.

package samplecompanyone

import (
	"fmt"
	"github.com/googleapis/gnostic/compiler"
	"strings"
)

func Version() string {
	return "samplecompanyone"
}

func NewSampleCompanyOneBook(in interface{}, context *compiler.Context) (*SampleCompanyOneBook, error) {
	errors := make([]error, 0)
	x := &SampleCompanyOneBook{}
	m, ok := compiler.UnpackMap(in)
	if !ok {
		message := fmt.Sprintf("has unexpected value: %+v", in)
		errors = append(errors, compiler.NewError(context, message))
	} else {
		requiredKeys := []string{"code", "message"}
		missingKeys := compiler.MissingKeysInMap(m, requiredKeys)
		if len(missingKeys) > 0 {
			message := fmt.Sprintf("is missing required %s: %+v", compiler.PluralProperties(len(missingKeys)), strings.Join(missingKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		allowedKeys := []string{"code", "message"}
		allowedPatterns := []string{}
		invalidKeys := compiler.InvalidKeysInMap(m, allowedKeys, allowedPatterns)
		if len(invalidKeys) > 0 {
			message := fmt.Sprintf("has invalid %s: %+v", compiler.PluralProperties(len(invalidKeys)), strings.Join(invalidKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		// int64 code = 1;
		v1 := compiler.MapValueForKey(m, "code")
		if v1 != nil {
			t, ok := v1.(int)
			if ok {
				x.Code = int64(t)
			} else {
				message := fmt.Sprintf("has unexpected value for code: %+v", v1)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// int64 message = 2;
		v2 := compiler.MapValueForKey(m, "message")
		if v2 != nil {
			t, ok := v2.(int)
			if ok {
				x.Message = int64(t)
			} else {
				message := fmt.Sprintf("has unexpected value for message: %+v", v2)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
	}
	return x, compiler.NewErrorGroupOrNil(errors)
}

func NewSampleCompanyOneShelve(in interface{}, context *compiler.Context) (*SampleCompanyOneShelve, error) {
	errors := make([]error, 0)
	x := &SampleCompanyOneShelve{}
	m, ok := compiler.UnpackMap(in)
	if !ok {
		message := fmt.Sprintf("has unexpected value: %+v", in)
		errors = append(errors, compiler.NewError(context, message))
	} else {
		requiredKeys := []string{"bar", "foo1"}
		missingKeys := compiler.MissingKeysInMap(m, requiredKeys)
		if len(missingKeys) > 0 {
			message := fmt.Sprintf("is missing required %s: %+v", compiler.PluralProperties(len(missingKeys)), strings.Join(missingKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		allowedKeys := []string{"bar", "foo1"}
		allowedPatterns := []string{}
		invalidKeys := compiler.InvalidKeysInMap(m, allowedKeys, allowedPatterns)
		if len(invalidKeys) > 0 {
			message := fmt.Sprintf("has invalid %s: %+v", compiler.PluralProperties(len(invalidKeys)), strings.Join(invalidKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		// int64 foo1 = 1;
		v1 := compiler.MapValueForKey(m, "foo1")
		if v1 != nil {
			t, ok := v1.(int)
			if ok {
				x.Foo1 = int64(t)
			} else {
				message := fmt.Sprintf("has unexpected value for foo1: %+v", v1)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// int64 bar = 2;
		v2 := compiler.MapValueForKey(m, "bar")
		if v2 != nil {
			t, ok := v2.(int)
			if ok {
				x.Bar = int64(t)
			} else {
				message := fmt.Sprintf("has unexpected value for bar: %+v", v2)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
	}
	return x, compiler.NewErrorGroupOrNil(errors)
}

func (m *SampleCompanyOneBook) ResolveReferences(root string) (interface{}, error) {
	errors := make([]error, 0)
	return nil, compiler.NewErrorGroupOrNil(errors)
}

func (m *SampleCompanyOneShelve) ResolveReferences(root string) (interface{}, error) {
	errors := make([]error, 0)
	return nil, compiler.NewErrorGroupOrNil(errors)
}
