// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package toolsversionchecker

import (
	"time"

	"github.com/juju/errors"
	"gopkg.in/juju/names.v2"
	"gopkg.in/juju/worker.v1"
	"gopkg.in/juju/worker.v1/dependency"

	"github.com/juju/juju/agent"
	apiagent "github.com/juju/juju/api/agent"
	"github.com/juju/juju/api/agenttools"
	"github.com/juju/juju/api/base"
	"github.com/juju/juju/cmd/jujud/agent/engine"
	"github.com/juju/juju/state/multiwatcher"
)

// ManifoldConfig defines the names of the manifolds on which a Manifold will depend.
type ManifoldConfig engine.AgentAPIManifoldConfig

// Manifold returns a dependency manifold that runs a toolsversionchecker worker,
// using the api connection resource named in the supplied config.
func Manifold(config ManifoldConfig) dependency.Manifold {
	typedConfig := engine.AgentAPIManifoldConfig(config)
	return engine.AgentAPIManifold(typedConfig, newWorker)
}

func newWorker(a agent.Agent, apiCaller base.APICaller) (worker.Worker, error) {
	st, err := apiagent.NewState(apiCaller)
	if err != nil {
		return nil, errors.Trace(err)
	}
	isMM, err := isModelManager(a, st)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if !isMM {
		return nil, dependency.ErrMissing
	}

	// 4 times a day seems a decent enough amount of checks.
	checkerParams := VersionCheckerParams{
		CheckInterval: time.Hour * 6,
	}
	return New(agenttools.NewFacade(apiCaller), &checkerParams), nil
}

func isModelManager(a agent.Agent, st *apiagent.State) (bool, error) {
	cfg := a.CurrentConfig()

	// Grab the tag and ensure that it's for a machine.
	tag, ok := cfg.Tag().(names.MachineTag)
	if !ok {
		return false, errors.New("this manifold may only be used inside a machine agent")
	}

	entity, err := st.Entity(tag)
	if err != nil {
		return false, err
	}

	for _, job := range entity.Jobs() {
		if job == multiwatcher.JobManageModel {
			return true, nil
		}
	}

	return false, nil
}
