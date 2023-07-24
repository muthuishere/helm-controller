/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	v1 "github.com/k3s-io/helm-controller/pkg/apis/helm.cattle.io/v1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	HelmChart() HelmChartController
	HelmChartConfig() HelmChartConfigController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) HelmChart() HelmChartController {
	return &HelmChartGenericController{
		generic.NewController[*v1.HelmChart, *v1.HelmChartList](schema.GroupVersionKind{Group: "helm.cattle.io", Version: "v1", Kind: "HelmChart"}, "helmcharts", true, v.controllerFactory),
	}
}

func (v *version) HelmChartConfig() HelmChartConfigController {
	return &HelmChartConfigGenericController{
		generic.NewController[*v1.HelmChartConfig, *v1.HelmChartConfigList](schema.GroupVersionKind{Group: "helm.cattle.io", Version: "v1", Kind: "HelmChartConfig"}, "helmchartconfigs", true, v.controllerFactory),
	}
}
