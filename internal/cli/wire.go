// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package cli

import (
	"context"

	"github.com/google/go-cloud/wire"
	"github.com/windmilleng/tilt/internal/build"
	"github.com/windmilleng/tilt/internal/engine"
	"github.com/windmilleng/tilt/internal/image"
	"github.com/windmilleng/tilt/internal/k8s"
	"github.com/windmilleng/tilt/internal/model"
	"github.com/windmilleng/tilt/internal/service"
	"github.com/windmilleng/wmclient/pkg/dirs"
)

func wireServiceCreator(ctx context.Context) (model.ServiceCreator, error) {
	wire.Build(
		service.ProvideMemoryManager,
		k8s.DetectEnv,
		image.NewImageHistory,
		build.DefaultDockerClient,
		engine.NewUpper,
		engine.NewLocalBuildAndDeployer,
		dirs.UseWindmillDir,
		provideServiceCreator)
	return nil, nil
}
