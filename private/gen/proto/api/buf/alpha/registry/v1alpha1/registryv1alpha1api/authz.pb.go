// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-api. DO NOT EDIT.

package registryv1alpha1api

import (
	context "context"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
)

// AuthzService supplies authorization helpers.
type AuthzService interface {
	// UserCanCreateOrganizationRepository returns whether the user is authorized
	// to create repositories in an organization.
	UserCanCreateOrganizationRepository(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanSeeRepositorySettings returns whether the user is authorized
	// to see repository settings.
	UserCanSeeRepositorySettings(ctx context.Context, repositoryId string) (authorized bool, err error)
	// UserCanSeeOrganizationSettings returns whether the user is authorized
	// to see organization settings.
	UserCanSeeOrganizationSettings(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanUpdateOrganizationSettings returns whether the user is authorized
	// to update organization settings.
	UserCanUpdateOrganizationSettings(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanReadPlugin returns whether the user has read access to the specified plugin.
	UserCanReadPlugin(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
	// UserCanCreatePluginVersion returns whether the user is authorized
	// to create a plugin version under the specified plugin.
	UserCanCreatePluginVersion(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
	// UserCanCreateTemplateVersion returns whether the user is authorized
	// to create a template version under the specified template.
	UserCanCreateTemplateVersion(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
	// UserCanCreateOrganizationPlugin returns whether the user is authorized to create
	// a plugin in an organization.
	UserCanCreateOrganizationPlugin(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanCreateOrganizationPlugin returns whether the user is authorized to create
	// a template in an organization.
	UserCanCreateOrganizationTemplate(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanSeePluginSettings returns whether the user is authorized
	// to see plugin settings.
	UserCanSeePluginSettings(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
	// UserCanSeeTemplateSettings returns whether the user is authorized
	// to see template settings.
	UserCanSeeTemplateSettings(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
	// UserCanAddOrganizationMember returns whether the user is authorized to add
	// any members to the organization and the list of roles they can add.
	UserCanAddOrganizationMember(ctx context.Context, organizationId string) (authorizedRoles []v1alpha1.OrganizationRole, err error)
	// UserCanUpdateOrganizationMember returns whether the user is authorized to update
	// any members' membership information in the organization and the list of roles they can update.
	UserCanUpdateOrganizationMember(ctx context.Context, organizationId string) (authorizedRoles []v1alpha1.OrganizationRole, err error)
	// UserCanRemoveOrganizationMember returns whether the user is authorized to remove
	// any members from the organization and the list of roles they can remove.
	UserCanRemoveOrganizationMember(ctx context.Context, organizationId string) (authorizedRoles []v1alpha1.OrganizationRole, err error)
	// UserCanDeleteOrganization returns whether the user is authorized
	// to delete an organization.
	UserCanDeleteOrganization(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanDeleteRepository returns whether the user is authorized
	// to delete a repository.
	UserCanDeleteRepository(ctx context.Context, repositoryId string) (authorized bool, err error)
	// UserCanDeleteTemplate returns whether the user is authorized
	// to delete a template.
	UserCanDeleteTemplate(ctx context.Context, templateId string) (authorized bool, err error)
	// UserCanDeletePlugin returns whether the user is authorized
	// to delete a plugin.
	UserCanDeletePlugin(ctx context.Context, pluginId string) (authorized bool, err error)
	// UserCanDeleteUser returns whether the user is authorized
	// to delete a user.
	UserCanDeleteUser(ctx context.Context) (authorized bool, err error)
	// UserCanSeeServerAdminPanel returns whether the user is authorized
	// to see server admin panel.
	UserCanSeeServerAdminPanel(ctx context.Context) (authorized bool, err error)
	// UserCanManageRepositoryContributors returns whether the user is authorized to manage
	// any contributors to the repository and the list of roles they can manage.
	UserCanManageRepositoryContributors(ctx context.Context, repositoryId string) (authorizedRoles []v1alpha1.RepositoryRole, err error)
	// UserCanManageRepositoryContributor returns whether the user is authorized to manage
	// a specific user to the repository and the list of roles they can manage.
	UserCanManageRepositoryContributor(
		ctx context.Context,
		repositoryId string,
		userId string,
	) (authorizedRoles []v1alpha1.RepositoryRole, err error)
	// UserCanManagePluginContributors returns whether the user is authorized to manage
	// any contributors to the plugin and the list of roles they can manage.
	UserCanManagePluginContributors(ctx context.Context, pluginId string) (authorizedRoles []v1alpha1.PluginRole, err error)
	// UserCanManagePluginContributor returns whether the user is authorized to manage
	// a specific user to the plugin and the list of roles they can manage.
	UserCanManagePluginContributor(
		ctx context.Context,
		pluginId string,
		userId string,
	) (authorizedRoles []v1alpha1.PluginRole, err error)
	// UserCanManageTemplateContributors returns whether the user is authorized to manage
	// any contributors to the template and the list of roles they can manage.
	UserCanManageTemplateContributors(ctx context.Context, templateId string) (authorizedRoles []v1alpha1.TemplateRole, err error)
	// UserCanManageTemplateContributor returns whether the user is authorized to manage
	// a specific user to the template and the list of roles they can manage.
	UserCanManageTemplateContributor(
		ctx context.Context,
		templateId string,
		userId string,
	) (authorizedRoles []v1alpha1.TemplateRole, err error)
}
