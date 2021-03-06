package shared

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	"github.com/openshift/origin/pkg/client"
	ocache "github.com/openshift/origin/pkg/client/cache"
)

type ClusterPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	// still use an indexer, no telling what someone will want to index on someday
	Indexer() cache.Indexer
	Lister() client.SyncedClusterPoliciesListerInterface
}

type clusterPolicyInformer struct {
	*sharedInformerFactory
}

func (f *clusterPolicyInformer) Informer() cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerObj := &authorizationapi.ClusterPolicy{}
	informerType := reflect.TypeOf(informerObj)
	informer, exists := f.coreInformers[informerType]
	if exists {
		return informer
	}

	lw := f.customListerWatchers.GetListerWatcher(authorizationapi.Resource("clusterpolicies"))
	if lw == nil {
		lw = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return f.originClient.ClusterPolicies().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return f.originClient.ClusterPolicies().Watch(options)
			},
		}
	}

	informer = cache.NewSharedIndexInformer(
		lw,
		informerObj,
		f.defaultResync,
		cache.Indexers{},
	)
	f.coreInformers[informerType] = informer

	return informer
}

func (f *clusterPolicyInformer) Indexer() cache.Indexer {
	informer := f.Informer()
	return informer.GetIndexer()
}

func (f *clusterPolicyInformer) Lister() client.SyncedClusterPoliciesListerInterface {
	return &ocache.InformerToClusterPolicyLister{SharedIndexInformer: f.Informer()}
}

type ClusterPolicyBindingInformer interface {
	Informer() cache.SharedIndexInformer
	// still use an indexer, no telling what someone will want to index on someday
	Indexer() cache.Indexer
	Lister() client.SyncedClusterPolicyBindingsListerInterface
}

type clusterPolicyBindingInformer struct {
	*sharedInformerFactory
}

func (f *clusterPolicyBindingInformer) Informer() cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerObj := &authorizationapi.ClusterPolicyBinding{}
	informerType := reflect.TypeOf(informerObj)
	informer, exists := f.coreInformers[informerType]
	if exists {
		return informer
	}

	lw := f.customListerWatchers.GetListerWatcher(authorizationapi.Resource("clusterpolicybindings"))
	if lw == nil {
		lw = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return f.originClient.ClusterPolicyBindings().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return f.originClient.ClusterPolicyBindings().Watch(options)
			},
		}
	}

	informer = cache.NewSharedIndexInformer(
		lw,
		informerObj,
		f.defaultResync,
		cache.Indexers{},
	)
	f.coreInformers[informerType] = informer

	return informer
}

func (f *clusterPolicyBindingInformer) Indexer() cache.Indexer {
	informer := f.Informer()
	return informer.GetIndexer()
}

func (f *clusterPolicyBindingInformer) Lister() client.SyncedClusterPolicyBindingsListerInterface {
	return &ocache.InformerToClusterPolicyBindingLister{SharedIndexInformer: f.Informer()}
}

type PolicyInformer interface {
	Informer() cache.SharedIndexInformer
	// still use an indexer, no telling what someone will want to index on someday
	Indexer() cache.Indexer
	Lister() client.SyncedPoliciesListerNamespacer
}

type policyInformer struct {
	*sharedInformerFactory
}

func (f *policyInformer) Informer() cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerObj := &authorizationapi.Policy{}
	informerType := reflect.TypeOf(informerObj)
	informer, exists := f.coreInformers[informerType]
	if exists {
		return informer
	}

	lw := f.customListerWatchers.GetListerWatcher(authorizationapi.Resource("policies"))
	if lw == nil {
		lw = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return f.originClient.Policies(metav1.NamespaceAll).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return f.originClient.Policies(metav1.NamespaceAll).Watch(options)
			},
		}
	}

	informer = cache.NewSharedIndexInformer(
		lw,
		informerObj,
		f.defaultResync,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	f.coreInformers[informerType] = informer

	return informer
}

func (f *policyInformer) Indexer() cache.Indexer {
	informer := f.Informer()
	return informer.GetIndexer()
}

func (f *policyInformer) Lister() client.SyncedPoliciesListerNamespacer {
	return &ocache.InformerToPolicyNamespacer{SharedIndexInformer: f.Informer()}
}

type PolicyBindingInformer interface {
	Informer() cache.SharedIndexInformer
	// still use an indexer, no telling what someone will want to index on someday
	Indexer() cache.Indexer
	Lister() client.SyncedPolicyBindingsListerNamespacer
}

type policyBindingInformer struct {
	*sharedInformerFactory
}

func (f *policyBindingInformer) Informer() cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerObj := &authorizationapi.PolicyBinding{}
	informerType := reflect.TypeOf(informerObj)
	informer, exists := f.coreInformers[informerType]
	if exists {
		return informer
	}

	lw := f.customListerWatchers.GetListerWatcher(authorizationapi.Resource("policybindings"))
	if lw == nil {
		lw = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return f.originClient.PolicyBindings(metav1.NamespaceAll).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return f.originClient.PolicyBindings(metav1.NamespaceAll).Watch(options)
			},
		}
	}

	informer = cache.NewSharedIndexInformer(
		lw,
		informerObj,
		f.defaultResync,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	f.coreInformers[informerType] = informer

	return informer
}

func (f *policyBindingInformer) Indexer() cache.Indexer {
	informer := f.Informer()
	return informer.GetIndexer()
}

func (f *policyBindingInformer) Lister() client.SyncedPolicyBindingsListerNamespacer {
	return &ocache.InformerToPolicyBindingNamespacer{SharedIndexInformer: f.Informer()}
}
