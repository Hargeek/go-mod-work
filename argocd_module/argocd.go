package argocd_module

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	argoappv1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"go-mod-work/dao"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	GitRepoURL    = "https://github.com/AliyunContainerService/gitops-demo.git"
	ClusterServer = "https://172.16.xx.xx:6443"
	ClusterName   = "cc0e3a41fe49c4a7bbcd5a1f0xxxxx-test"
)

type ArgoCDClient struct {
	projectClient project.ProjectServiceClient
	clusterClient cluster.ClusterServiceClient
	appClient     application.ApplicationServiceClient
}

type Connection struct {
	Address string
	Token   string
}

// Init ArgoCD client
func Init() {
	_, _ = dao.ArgoCD.GetArgoCDList()

}

func NewArgoCDClient(conn *Connection) (*ArgoCDClient, error) {
	argoCDApiClient, err := apiclient.NewClient(&apiclient.ClientOptions{
		ServerAddr: fmt.Sprintf(conn.Address),
		Insecure:   true,
		AuthToken:  conn.Token,
	})
	if err != nil {
		return nil, err
	}

	_, appClient, err := argoCDApiClient.NewApplicationClient()
	if err != nil {
		return nil, err
	}

	_, projectClient, err := argoCDApiClient.NewProjectClient()
	if err != nil {
		return nil, err
	}

	_, clusterClient, err := argoCDApiClient.NewClusterClient()
	if err != nil {
		return nil, err
	}

	return &ArgoCDClient{
		projectClient: projectClient,
		clusterClient: clusterClient,
		appClient:     appClient,
	}, nil
}

func (c *ArgoCDClient) CreateApplication(upsert, validate bool) error {
	app := &argoappv1.Application{
		ObjectMeta: v1.ObjectMeta{
			Name:      "demo-argocd-app-manual",
			Namespace: "argocd",
		},
		Spec: argoappv1.ApplicationSpec{
			Project: "default",
			Source: &argoappv1.ApplicationSource{
				RepoURL:        GitRepoURL,
				Path:           "manifests/helm/echo-server",
				TargetRevision: "HEAD",
			},
			Destination: argoappv1.ApplicationDestination{
				Server:    ClusterServer,
				Name:      ClusterName,
				Namespace: "echo-server-demo2",
			},
			SyncPolicy: &argoappv1.SyncPolicy{
				SyncOptions: argoappv1.SyncOptions{
					"CreateNamespace=true",
				},
			},
		},
	}

	if _, err := c.appClient.Create(context.Background(), &application.ApplicationCreateRequest{
		Application: app,
		Upsert:      &upsert,
		Validate:    &validate,
	}); err != nil {
		return err
	}

	return nil
}
