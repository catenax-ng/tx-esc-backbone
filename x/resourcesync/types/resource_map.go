package types

func NewResourceMap(resource Resource) ResourceMap {
	return ResourceMap{
		Resource:  resource,
		AuditLogs: nil,
	}
}
