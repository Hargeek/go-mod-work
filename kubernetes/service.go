package kubernetes

import (
	"context"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type svc struct{}

func (s *svc) GetServices(client *kubernetes.Clientset, namespace string) (serviceList *v1.ServiceList, err error) {
	serviceList, err = client.CoreV1().Services(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, errors.New("get service list failed, " + err.Error())
	}
	return
}
