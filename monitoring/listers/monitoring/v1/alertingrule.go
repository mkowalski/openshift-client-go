// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AlertingRuleLister helps list AlertingRules.
// All objects returned here must be treated as read-only.
type AlertingRuleLister interface {
	// List lists all AlertingRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AlertingRule, err error)
	// AlertingRules returns an object that can list and get AlertingRules.
	AlertingRules(namespace string) AlertingRuleNamespaceLister
	AlertingRuleListerExpansion
}

// alertingRuleLister implements the AlertingRuleLister interface.
type alertingRuleLister struct {
	indexer cache.Indexer
}

// NewAlertingRuleLister returns a new AlertingRuleLister.
func NewAlertingRuleLister(indexer cache.Indexer) AlertingRuleLister {
	return &alertingRuleLister{indexer: indexer}
}

// List lists all AlertingRules in the indexer.
func (s *alertingRuleLister) List(selector labels.Selector) (ret []*v1.AlertingRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AlertingRule))
	})
	return ret, err
}

// AlertingRules returns an object that can list and get AlertingRules.
func (s *alertingRuleLister) AlertingRules(namespace string) AlertingRuleNamespaceLister {
	return alertingRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlertingRuleNamespaceLister helps list and get AlertingRules.
// All objects returned here must be treated as read-only.
type AlertingRuleNamespaceLister interface {
	// List lists all AlertingRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AlertingRule, err error)
	// Get retrieves the AlertingRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AlertingRule, error)
	AlertingRuleNamespaceListerExpansion
}

// alertingRuleNamespaceLister implements the AlertingRuleNamespaceLister
// interface.
type alertingRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AlertingRules in the indexer for a given namespace.
func (s alertingRuleNamespaceLister) List(selector labels.Selector) (ret []*v1.AlertingRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AlertingRule))
	})
	return ret, err
}

// Get retrieves the AlertingRule from the indexer for a given namespace and name.
func (s alertingRuleNamespaceLister) Get(name string) (*v1.AlertingRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("alertingrule"), name)
	}
	return obj.(*v1.AlertingRule), nil
}
