package k8s

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodFinding struct {
	Namespace string
	Name      string
	Reason    string
	Message   string
	Restarts  int32
}

func ScanPods(client *kubernetes.Clientset, namespace string) ([]PodFinding, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := metav1.ListOptions{}
	pods, err := client.CoreV1().Pods(namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}

	findings := []PodFinding{}

	for _, pod := range pods.Items {
		for _, cs := range pod.Status.ContainerStatuses {
			if cs.State.Waiting != nil {
				reason := cs.State.Waiting.Reason
				if reason == "CrashLoopBackOff" || reason == "ImagePullBackOff" {
					findings = append(findings, PodFinding{
						Namespace: pod.Namespace,
						Name:      pod.Name,
						Reason:    reason,
						Message:   cs.State.Waiting.Message,
						Restarts:  cs.RestartCount,
					})
				}
			}
		}
	}

	return findings, nil
}
