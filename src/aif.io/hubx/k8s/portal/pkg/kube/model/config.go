
package model

import (
	"github.com/gogo/protobuf/proto"
	"time"
)

const (
	// IstioAPIGroupDomain defines API group domain of all Istio configuration resources.
	// Group domain suffix to the proto schema's group to generate the full resource group.
	IstioAPIGroupDomain = ".istio.io"

	// Default API version of an Istio config proto message.
	istioAPIVersion = "v1alpha2"

	// NamespaceAll is a designated symbol for listing across all namespaces
	NamespaceAll = ""

	// IstioMeshGateway is the built in gateway for all sidecars
	IstioMeshGateway = "mesh"

	// IstioSystemNamespace is the namespace where Istio's components are deployed
	IstioSystemNamespace = "istio-system"
)


var (



	Listener = ProtoSchema{
		Type:        "listener",
		Plural:      "listeners",
		Group:       "networking",
		Version:     "v1alpha3",
		MessageName: "aif.io.hubx.k8s.portal.api.v1.Listener",
		Validate:    ValidateFunction,
		Collection:  "istio/networking/v1alpha3/listeners",
	}

	/*IstioNetworkingV1alpha3Functions = b.Register(
		"istio/networking/v1alpha3/functions",
		"type.googleapis.com/aif.io.api.networking.v1.Function")
*/
	// IstioConfigTypes lists all Istio config types with schemas and validation
	IstioConfigTypes = ConfigDescriptor{
		Listener,
	}
)

type ConfigStore interface {
	// ConfigDescriptor exposes the configuration type schema known by the config store.
	// The type schema defines the bidrectional mapping between configuration
	// types and the protobuf encoding schema.
	ConfigDescriptor() ConfigDescriptor

	// Get retrieves a configuration element by a type and a key
	Get(typ, name, namespace string) *Config

	// List returns objects by type and namespace.
	// Use "" for the namespace to list across namespaces.
	List(typ, namespace string) ([]Config, error)

	// Create adds a new configuration object to the store. If an object with the
	// same name and namespace for the type already exists, the operation fails
	// with no side effects.
	Create(config Config) (revision string, err error)

	// Update modifies an existing configuration object in the store.  Update
	// requires that the object has been created.  Resource version prevents
	// overriding a value that has been changed between prior _Get_ and _Put_
	// operation to achieve optimistic concurrency. This method returns a new
	// revision if the operation succeeds.
	Update(config Config) (newRevision string, err error)

	// Delete removes an object from the store by key
	Delete(typ, name, namespace string) error
}

type ConfigStoreCache interface {
	ConfigStore

	// RegisterEventHandler adds a handler to receive config update events for a
	// configuration type
	RegisterEventHandler(typ string, handler func(Config, Event))

	// Run until a signal is received
	Run(stop <-chan struct{})

	// HasSynced returns true after initial cache synchronization is complete
	HasSynced() bool
}

func ValidateFunction(name, namespace string, msg proto.Message) (errs error){
	return ;
}
type ConfigMeta struct {
	// Type is a short configuration name that matches the content message type
	// (e.g. "route-rule")
	Type string `json:"type,omitempty"`

	// Group is the API group of the config.
	Group string `json:"group,omitempty"`

	// Version is the API version of the Config.
	Version string `json:"version,omitempty"`

	// Name is a unique immutable identifier in a namespace
	Name string `json:"name,omitempty"`

	// Namespace defines the space for names (optional for some types),
	// applications may choose to use namespaces for a variety of purposes
	// (security domains, fault domains, organizational domains)
	Namespace string `json:"namespace,omitempty"`

	// Domain defines the suffix of the fully qualified name past the namespace.
	// Domain is not a part of the unique key unlike name and namespace.
	Domain string `json:"domain,omitempty"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects.
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	Annotations map[string]string `json:"annotations,omitempty"`

	// ResourceVersion is an opaque identifier for tracking updates to the config registry.
	// The implementation may use a change index or a commit log for the revision.
	// The config client should not make any assumptions about revisions and rely only on
	// exact equality to implement optimistic concurrency of read-write operations.
	//
	// The lifetime of an object of a particular revision depends on the underlying data store.
	// The data store may compactify old revisions in the interest of storage optimization.
	//
	// An empty revision carries a special meaning that the associated object has
	// not been stored and assigned a revision.
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// CreationTimestamp records the creation time
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}

// Config is a configuration unit consisting of the type of configuration, the
// key identifier that is unique per type, and the content represented as a
// protobuf message.
type Config struct {
	ConfigMeta

	// Spec holds the configuration object as a protobuf message
	Spec proto.Message
}

// ConfigDescriptor defines the bijection between the short type name and its
// fully qualified protobuf message name
type ConfigDescriptor []ProtoSchema

// ProtoSchema provides description of the configuration schema and its key function
// nolint: maligned
type ProtoSchema struct {
	// ClusterScoped is true for resource in cluster-level.
	ClusterScoped bool

	// Type is the config proto type.
	Type string

	// Plural is the type in plural.
	Plural string

	// Group is the config proto group.
	Group string

	// Version is the config proto version.
	Version string

	// MessageName refers to the protobuf message type name corresponding to the type
	MessageName string

	// Validate configuration as a protobuf message assuming the object is an
	// instance of the expected message type
	Validate func(name, namespace string, config proto.Message) error

	// MCP collection for this configuration resource schema
	Collection string
}

// GetByType finds a schema by type if it is available
func (descriptor ConfigDescriptor) GetByType(name string) (ProtoSchema, bool) {
	for _, schema := range descriptor {
		if schema.Type == name {
			return schema, true
		}
	}
	return ProtoSchema{}, false
}




