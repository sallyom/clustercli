// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	configv1 "github.com/openshift/api/config/v1"
	versioned "github.com/openshift/client-go/config/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/config/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/config/listers/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterCLIInformer provides access to a shared informer and lister for
// ClusterCLIs.
type ClusterCLIInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterCLILister
}

type clusterCLIInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterCLIInformer constructs a new informer for ClusterCLI type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterCLIInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterCLIInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterCLIInformer constructs a new informer for ClusterCLI type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterCLIInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().ClusterCLIs().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().ClusterCLIs().Watch(context.TODO(), options)
			},
		},
		&configv1.ClusterCLI{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterCLIInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterCLIInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterCLIInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1.ClusterCLI{}, f.defaultInformer)
}

func (f *clusterCLIInformer) Lister() v1.ClusterCLILister {
	return v1.NewClusterCLILister(f.Informer().GetIndexer())
}
