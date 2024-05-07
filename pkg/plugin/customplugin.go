package plugin

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// Name is the name of the plugin used in the plugin registry and configurations.

const Name = "custom-prescore-plugin"

type custom_prescore_plugin struct{}

var _ framework.PreScorePlugin = &custom_prescore_plugin{}

// New initializes a new plugin and returns it.
func New(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &custom_prescore_plugin{}, nil
}

// Name returns name of the plugin.
func (pl *custom_prescore_plugin) Name() string {
	return Name
}

func (pl *custom_prescore_plugin) PreScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodes []*v1.Node) *framework.Status {
	log.Println(nodes)
	return framework.NewStatus(framework.Success, "Node: "+pod.Name)
}
