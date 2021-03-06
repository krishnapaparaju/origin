// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/openshift/client-go/template/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BrokerTemplateInstances returns a BrokerTemplateInstanceInformer.
	BrokerTemplateInstances() BrokerTemplateInstanceInformer
	// Templates returns a TemplateInformer.
	Templates() TemplateInformer
	// TemplateInstances returns a TemplateInstanceInformer.
	TemplateInstances() TemplateInstanceInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// BrokerTemplateInstances returns a BrokerTemplateInstanceInformer.
func (v *version) BrokerTemplateInstances() BrokerTemplateInstanceInformer {
	return &brokerTemplateInstanceInformer{factory: v.SharedInformerFactory}
}

// Templates returns a TemplateInformer.
func (v *version) Templates() TemplateInformer {
	return &templateInformer{factory: v.SharedInformerFactory}
}

// TemplateInstances returns a TemplateInstanceInformer.
func (v *version) TemplateInstances() TemplateInstanceInformer {
	return &templateInstanceInformer{factory: v.SharedInformerFactory}
}
