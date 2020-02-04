// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterCLILister helps list ClusterCLIs.
// All objects returned here must be treated as read-only.
type ClusterCLILister interface {
	// List lists all ClusterCLIs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterCLI, err error)
	// Get retrieves the ClusterCLI from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterCLI, error)
	ClusterCLIListerExpansion
}

// clusterCLILister implements the ClusterCLILister interface.
type clusterCLILister struct {
	indexer cache.Indexer
}

// NewClusterCLILister returns a new ClusterCLILister.
func NewClusterCLILister(indexer cache.Indexer) ClusterCLILister {
	return &clusterCLILister{indexer: indexer}
}

// List lists all ClusterCLIs in the indexer.
func (s *clusterCLILister) List(selector labels.Selector) (ret []*v1.ClusterCLI, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterCLI))
	})
	return ret, err
}

// Get retrieves the ClusterCLI from the index for a given name.
func (s *clusterCLILister) Get(name string) (*v1.ClusterCLI, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clustercli"), name)
	}
	return obj.(*v1.ClusterCLI), nil
}
