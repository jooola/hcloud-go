package schema

import "time"

// PrimaryIP defines a Primary IP.
type PrimaryIP struct {
	ID           int64               `json:"id"`
	IP           string              `json:"ip"`
	Labels       map[string]string   `json:"labels"`
	Name         string              `json:"name"`
	Type         string              `json:"type"`
	Protection   PrimaryIPProtection `json:"protection"`
	DNSPtr       []PrimaryIPDNSPTR   `json:"dns_ptr"`
	AssigneeID   *int64              `json:"assignee_id"`
	AssigneeType string              `json:"assignee_type"`
	AutoDelete   bool                `json:"auto_delete"`
	Blocked      bool                `json:"blocked"`
	Created      time.Time           `json:"created"`
	Datacenter   Datacenter          `json:"datacenter"`
}

// PrimaryIPProtection represents the protection level of a Primary IP.
type PrimaryIPProtection struct {
	Delete bool `json:"delete"`
}

// PrimaryIPDNSPTR contains reverse DNS information for a
// IPv4 or IPv6 Primary IP.
type PrimaryIPDNSPTR struct {
	DNSPtr string `json:"dns_ptr"`
	IP     string `json:"ip"`
}

// PrimaryIPCreateOpts defines the request to
// create a Primary IP.
type PrimaryIPCreateRequest struct {
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	AssigneeType string            `json:"assignee_type"`
	AssigneeID   *int64            `json:"assignee_id,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	AutoDelete   *bool             `json:"auto_delete,omitempty"`
	Datacenter   string            `json:"datacenter,omitempty"`
}

// PrimaryIPCreateResponse defines the schema of the response
// when creating a Primary IP.
type PrimaryIPCreateResponse struct {
	PrimaryIP PrimaryIP `json:"primary_ip"`
	Action    *Action   `json:"action"`
}

// PrimaryIPGetResponse defines the response when retrieving a single Primary IP.
type PrimaryIPGetResponse struct {
	PrimaryIP PrimaryIP `json:"primary_ip"`
}

// PrimaryIPListResponse defines the response when listing Primary IPs.
type PrimaryIPListResponse struct {
	PrimaryIPs []PrimaryIP `json:"primary_ips"`
}

// PrimaryIPUpdateOpts defines the request to
// update a Primary IP.
type PrimaryIPUpdateRequest struct {
	Name       string            `json:"name,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
	AutoDelete *bool             `json:"auto_delete,omitempty"`
}

// PrimaryIPUpdateResponse defines the response
// when updating a Primary IP.
type PrimaryIPUpdateResponse struct {
	PrimaryIP PrimaryIP `json:"primary_ip"`
}

// PrimaryIPActionChangeDNSPtrRequest defines the schema for the request to
// change a Primary IP's reverse DNS pointer.
type PrimaryIPActionChangeDNSPtrRequest struct {
	IP     string  `json:"ip"`
	DNSPtr *string `json:"dns_ptr"`
}

// PrimaryIPActionChangeDNSPtrResponse defines the response when setting a reverse DNS
// pointer for a IP address.
type PrimaryIPActionChangeDNSPtrResponse struct {
	Action Action `json:"action"`
}

// PrimaryIPActionAssignRequest defines the request to
// assign a Primary IP to an assignee (usually a server).
type PrimaryIPActionAssignRequest struct {
	AssigneeID   int64  `json:"assignee_id"`
	AssigneeType string `json:"assignee_type"`
}

// PrimaryIPActionAssignResponse defines the response when assigning a Primary IP to a
// assignee.
type PrimaryIPActionAssignResponse struct {
	Action Action `json:"action"`
}

// PrimaryIPActionUnassignResponse defines the response to unassign a Primary IP.
type PrimaryIPActionUnassignResponse struct {
	Action Action `json:"action"`
}

// PrimaryIPActionChangeProtectionRequest defines the request to
// change protection configuration of a Primary IP.
type PrimaryIPActionChangeProtectionRequest struct {
	Delete bool `json:"delete"`
}

// PrimaryIPActionChangeProtectionResponse defines the response when changing the
// protection of a Primary IP.
type PrimaryIPActionChangeProtectionResponse struct {
	Action Action `json:"action"`
}
