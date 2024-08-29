// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package externalreferenceaction

import (
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	referenceaction "github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/referenceaction"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows"
)

type SomeConfig referenceaction.SomeConfig

func (cfg SomeConfig) New(w *workflows.WorkflowSpecFactory, ref string, input ActionInput) referenceaction.SomeOutputsCap {

	def := workflows.StepDefinition{
		ID: "external-reference-test-action@1.0.0", Ref: ref,
		Inputs:         input.ToSteps(),
		Config:         map[string]any{},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	step := workflows.Step[referenceaction.SomeOutputs]{Definition: def}
	return referenceaction.SomeOutputsCapFromStep(w, step)
}

type ActionInput = referenceaction.ActionInput