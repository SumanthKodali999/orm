/*
Copyright 2022.

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

package kubernetes

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	iLog = ctrl.Log.WithName("informer")
)

type InformerFactory struct {
	dynamicinformer.DynamicSharedInformerFactory
}

func (i *InformerFactory) WatchResourceWithGVK(gvk schema.GroupVersionKind, handler cache.ResourceEventHandler) error {

	gvr := r.FindGVRfromGVK(gvk)
	if gvr == nil {
		return errors.New("Source resource " + gvk.String() + "is not installed")
	}

	i.ForResource(*gvr).Informer().AddEventHandler(handler)

	i.Start(r.ctx.Done())
	return nil
}