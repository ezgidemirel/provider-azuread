package administrativeunits

import "github.com/upbound/upjet/pkg/config"

const group = "administrativeunits"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_administrative_unit", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_administrative_unit_member", func(r *config.Resource) {
		r.Kind = "Member"
		r.References["administrative_unit_object_id"] = config.Reference{
			Type: "Unit",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
