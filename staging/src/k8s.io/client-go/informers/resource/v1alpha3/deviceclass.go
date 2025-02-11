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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha3

import (
	context "context"
	time "time"

	apiresourcev1alpha3 "k8s.io/api/resource/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	resourcev1alpha3 "k8s.io/client-go/listers/resource/v1alpha3"
	cache "k8s.io/client-go/tools/cache"
)

// DeviceClassInformer provides access to a shared informer and lister for
// DeviceClasses.
type DeviceClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() resourcev1alpha3.DeviceClassLister
}

type deviceClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDeviceClassInformer constructs a new informer for DeviceClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeviceClassInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeviceClassInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDeviceClassInformer constructs a new informer for DeviceClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeviceClassInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().DeviceClasses().List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().DeviceClasses().Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().DeviceClasses().List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().DeviceClasses().Watch(ctx, options)
			},
		},
		&apiresourcev1alpha3.DeviceClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *deviceClassInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeviceClassInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *deviceClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiresourcev1alpha3.DeviceClass{}, f.defaultInformer)
}

func (f *deviceClassInformer) Lister() resourcev1alpha3.DeviceClassLister {
	return resourcev1alpha3.NewDeviceClassLister(f.Informer().GetIndexer())
}
