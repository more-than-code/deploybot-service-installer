package task

import (
	"encoding/json"

	types "deploybot-service-launcher/deploybot-types"
	"deploybot-service-launcher/util"

	"github.com/kelseyhightower/envconfig"
)

type RunnerConfig struct {
	ProjectsPath string `envconfig:"PROJECTS_PATH"`
	DockerHost   string `envconfig:"DOCKER_HOST"`
}

type Runner struct {
	cfg     RunnerConfig
	cHelper *util.ContainerHelper
}

func NewRunner() *Runner {
	var cfg RunnerConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	return &Runner{cfg: cfg, cHelper: util.NewContainerHelper(cfg.DockerHost)}
}

func (r *Runner) DoTask(t types.Task, arguments []string) error {

	var c types.DeployConfig

	bs, err := json.Marshal(t.Config)

	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, &c)

	if err != nil {
		return err
	}

	go func() {
		r.cHelper.StartContainer(&c)
	}()

	return nil
}
