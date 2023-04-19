// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/openshift/client-go/monitoring/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AlertRelabelConfigs returns a AlertRelabelConfigInformer.
	AlertRelabelConfigs() AlertRelabelConfigInformer
	// AlertingRules returns a AlertingRuleInformer.
	AlertingRules() AlertingRuleInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AlertRelabelConfigs returns a AlertRelabelConfigInformer.
func (v *version) AlertRelabelConfigs() AlertRelabelConfigInformer {
	return &alertRelabelConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AlertingRules returns a AlertingRuleInformer.
func (v *version) AlertingRules() AlertingRuleInformer {
	return &alertingRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
