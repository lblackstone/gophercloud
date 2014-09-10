package gophercloud

import "errors"

var (
	// ErrServiceNotFound is returned when no service matches the EndpointOpts.
	ErrServiceNotFound = errors.New("No suitable service could be found in the service catalog.")

	// ErrEndpointNotFound is returned when no available endpoints match the provided EndpointOpts.
	ErrEndpointNotFound = errors.New("No suitable endpoint could be found in the service catalog.")
)

// Availability describes the accessibility of a specific service endpoint.
type Availability string

const (
	// AvailabilityAdmin makes an endpoint only available to administrators.
	AvailabilityAdmin Availability = "admin"

	// AvailabilityPublic makes an endpoint available to everyone.
	AvailabilityPublic Availability = "public"

	// AvailabilityInternal makes an endpoint only available within the cluster.
	AvailabilityInternal Availability = "internal"
)

// EndpointOpts contains options for finding an endpoint for an Openstack client.
type EndpointOpts struct {

	// Type is the service type for the client (e.g., "compute", "object-store").
	// Type is a required field.
	Type string

	// Name is the service name for the client (e.g., "nova").
	// Name is not a required field, but it is used if present.
	// Services can have the same Type but a different Name, which is one example of when both Type and Name are needed.
	Name string

	// Region is the region in which the service resides.
	// Region must be specified for services that span multiple regions.
	Region string

	// Availability is the visibility of the endpoint to be returned: AvailabilityPublic, AvailabilityInternal, or AvailabilityAdmin.
	// Availability is not required, and defaults to InterfacePublic.
	// Not all providers or services offer all Availability options.
	Availability Availability
}

// EndpointLocator is a function that describes how to locate a single endpoint from a service catalog for a specific ProviderClient.
// It should be set during ProviderClient authentication and used to discover related ServiceClients.
type EndpointLocator func(EndpointOpts) (string, error)
