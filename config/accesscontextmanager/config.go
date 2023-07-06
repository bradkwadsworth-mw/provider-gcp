package accesscontextmanager

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_access_context_manager_access_policy", func(r *config.Resource) {})
	p.AddResourceConfigurator("google_access_context_manager_access_level", func(r *config.Resource) {
		r.References["parent"] = config.Reference{
			Type: "AccessPolicy",
		}
	})
	p.AddResourceConfigurator("google_access_context_manager_access_level_condition", func(r *config.Resource) {
		r.References["access_level"] = config.Reference{
			Type: "AccessLevel",
		}
	})
	p.AddResourceConfigurator("google_access_context_manager_gcp_user_access_binding", func(r *config.Resource) {
		r.References["access_levels"] = config.Reference{
			Type: "AccessLevel",
		}
	})
	p.AddResourceConfigurator("google_access_context_manager_access_policy_iam_member", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "AccessPolicy",
		}
	})
	p.AddResourceConfigurator("google_access_context_manager_service_perimeter", func(r *config.Resource) {
		r.References["parent"] = config.Reference{
			Type: "AccessPolicy",
		}
		r.References["status.access_levels"] = config.Reference{
			Type: "AccessLevel",
		}
		r.References["spec.access_levels"] = config.Reference{
			Type: "AccessLevel",
		}
	})
	p.AddResourceConfigurator("google_access_context_manager_service_perimeter_resource", func(r *config.Resource) {
		r.References["perimeter_name"] = config.Reference{
			Type: "ServicePerimeter",
		}
		if s, ok := r.TerraformResource.Schema["resource"]; ok {
			s.Optional = false
			s.Computed = false
		}
	})
}