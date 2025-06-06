// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ProjectsGetter has a method to return a ProjectInterface.
// A group's client should implement this interface.
type ProjectsGetter interface {
	Projects() ProjectInterface
}

// ProjectInterface has methods to work with Project resources.
type ProjectInterface interface {
	Create(ctx context.Context, project *managementv1.Project, opts metav1.CreateOptions) (*managementv1.Project, error)
	Update(ctx context.Context, project *managementv1.Project, opts metav1.UpdateOptions) (*managementv1.Project, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, project *managementv1.Project, opts metav1.UpdateOptions) (*managementv1.Project, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*managementv1.Project, error)
	List(ctx context.Context, opts metav1.ListOptions) (*managementv1.ProjectList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *managementv1.Project, err error)
	ListMembers(ctx context.Context, projectName string, options metav1.GetOptions) (*managementv1.ProjectMembers, error)
	ListTemplates(ctx context.Context, projectName string, options metav1.GetOptions) (*managementv1.ProjectTemplates, error)
	ListRunners(ctx context.Context, projectName string, options metav1.GetOptions) (*managementv1.ProjectRunners, error)
	ListClusters(ctx context.Context, projectName string, options metav1.GetOptions) (*managementv1.ProjectClusters, error)
	MigrateVirtualClusterInstance(ctx context.Context, projectName string, projectMigrateVirtualClusterInstance *managementv1.ProjectMigrateVirtualClusterInstance, opts metav1.CreateOptions) (*managementv1.ProjectMigrateVirtualClusterInstance, error)
	ImportSpace(ctx context.Context, projectName string, projectImportSpace *managementv1.ProjectImportSpace, opts metav1.CreateOptions) (*managementv1.ProjectImportSpace, error)
	MigrateSpaceInstance(ctx context.Context, projectName string, projectMigrateSpaceInstance *managementv1.ProjectMigrateSpaceInstance, opts metav1.CreateOptions) (*managementv1.ProjectMigrateSpaceInstance, error)

	ProjectExpansion
}

// projects implements ProjectInterface
type projects struct {
	*gentype.ClientWithList[*managementv1.Project, *managementv1.ProjectList]
}

// newProjects returns a Projects
func newProjects(c *ManagementV1Client) *projects {
	return &projects{
		gentype.NewClientWithList[*managementv1.Project, *managementv1.ProjectList](
			"projects",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *managementv1.Project { return &managementv1.Project{} },
			func() *managementv1.ProjectList { return &managementv1.ProjectList{} },
		),
	}
}

// ListMembers takes name of the project, and returns the corresponding managementv1.ProjectMembers object, and an error if there is any.
func (c *projects) ListMembers(ctx context.Context, projectName string, options metav1.GetOptions) (result *managementv1.ProjectMembers, err error) {
	result = &managementv1.ProjectMembers{}
	err = c.GetClient().Get().
		Resource("projects").
		Name(projectName).
		SubResource("members").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// ListTemplates takes name of the project, and returns the corresponding managementv1.ProjectTemplates object, and an error if there is any.
func (c *projects) ListTemplates(ctx context.Context, projectName string, options metav1.GetOptions) (result *managementv1.ProjectTemplates, err error) {
	result = &managementv1.ProjectTemplates{}
	err = c.GetClient().Get().
		Resource("projects").
		Name(projectName).
		SubResource("templates").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// ListRunners takes name of the project, and returns the corresponding managementv1.ProjectRunners object, and an error if there is any.
func (c *projects) ListRunners(ctx context.Context, projectName string, options metav1.GetOptions) (result *managementv1.ProjectRunners, err error) {
	result = &managementv1.ProjectRunners{}
	err = c.GetClient().Get().
		Resource("projects").
		Name(projectName).
		SubResource("runners").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// ListClusters takes name of the project, and returns the corresponding managementv1.ProjectClusters object, and an error if there is any.
func (c *projects) ListClusters(ctx context.Context, projectName string, options metav1.GetOptions) (result *managementv1.ProjectClusters, err error) {
	result = &managementv1.ProjectClusters{}
	err = c.GetClient().Get().
		Resource("projects").
		Name(projectName).
		SubResource("clusters").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// MigrateVirtualClusterInstance takes the representation of a projectMigrateVirtualClusterInstance and creates it.  Returns the server's representation of the projectMigrateVirtualClusterInstance, and an error, if there is any.
func (c *projects) MigrateVirtualClusterInstance(ctx context.Context, projectName string, projectMigrateVirtualClusterInstance *managementv1.ProjectMigrateVirtualClusterInstance, opts metav1.CreateOptions) (result *managementv1.ProjectMigrateVirtualClusterInstance, err error) {
	result = &managementv1.ProjectMigrateVirtualClusterInstance{}
	err = c.GetClient().Post().
		Resource("projects").
		Name(projectName).
		SubResource("migratevirtualclusterinstance").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectMigrateVirtualClusterInstance).
		Do(ctx).
		Into(result)
	return
}

// ImportSpace takes the representation of a projectImportSpace and creates it.  Returns the server's representation of the projectImportSpace, and an error, if there is any.
func (c *projects) ImportSpace(ctx context.Context, projectName string, projectImportSpace *managementv1.ProjectImportSpace, opts metav1.CreateOptions) (result *managementv1.ProjectImportSpace, err error) {
	result = &managementv1.ProjectImportSpace{}
	err = c.GetClient().Post().
		Resource("projects").
		Name(projectName).
		SubResource("importspace").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectImportSpace).
		Do(ctx).
		Into(result)
	return
}

// MigrateSpaceInstance takes the representation of a projectMigrateSpaceInstance and creates it.  Returns the server's representation of the projectMigrateSpaceInstance, and an error, if there is any.
func (c *projects) MigrateSpaceInstance(ctx context.Context, projectName string, projectMigrateSpaceInstance *managementv1.ProjectMigrateSpaceInstance, opts metav1.CreateOptions) (result *managementv1.ProjectMigrateSpaceInstance, err error) {
	result = &managementv1.ProjectMigrateSpaceInstance{}
	err = c.GetClient().Post().
		Resource("projects").
		Name(projectName).
		SubResource("migratespaceinstance").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectMigrateSpaceInstance).
		Do(ctx).
		Into(result)
	return
}
