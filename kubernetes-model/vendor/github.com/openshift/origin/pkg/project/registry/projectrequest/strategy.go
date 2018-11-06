/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package projectrequest

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/kubernetes/pkg/api/legacyscheme"

	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	projectvalidation "github.com/openshift/origin/pkg/project/apis/project/validation"
)

// strategy implements behavior for OAuthClient objects
type strategy struct {
	runtime.ObjectTyper
}

var Strategy = strategy{legacyscheme.Scheme}

func (strategy) PrepareForUpdate(ctx apirequest.Context, obj, old runtime.Object) {}

// NamespaceScoped is false for projectrequest objects
func (strategy) NamespaceScoped() bool {
	return false
}

func (strategy) GenerateName(base string) string {
	return base
}

func (strategy) PrepareForCreate(ctx apirequest.Context, obj runtime.Object) {
}

// Validate validates a new client
func (strategy) Validate(ctx apirequest.Context, obj runtime.Object) field.ErrorList {
	projectrequest := obj.(*projectapi.ProjectRequest)
	return projectvalidation.ValidateProjectRequest(projectrequest)
}

// ValidateUpdate validates a client update
func (strategy) ValidateUpdate(ctx apirequest.Context, obj runtime.Object, old runtime.Object) field.ErrorList {
	return nil
}

// Canonicalize normalizes the object after validation.
func (strategy) Canonicalize(obj runtime.Object) {
}

// AllowCreateOnUpdate is false for OAuth objects
func (strategy) AllowCreateOnUpdate() bool {
	return false
}