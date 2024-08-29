// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package streams

import (
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows"
)

func (cfg TriggerConfig) New(w *workflows.WorkflowSpecFactory) workflows.CapDefinition[[]Feed] {
	ref := "trigger"
	def := workflows.StepDefinition{
		ID: "streams-trigger@1.0.0", Ref: ref,
		Inputs: workflows.StepInputs{},
		Config: map[string]any{
			"feedIds":        cfg.FeedIds,
			"maxFrequencyMs": cfg.MaxFrequencyMs,
		},
		CapabilityType: capabilities.CapabilityTypeTrigger,
	}

	step := workflows.Step[[]Feed]{Definition: def}
	return step.AddTo(w)
}

type FeedCap interface {
	workflows.CapDefinition[Feed]
	BenchmarkPrice() workflows.CapDefinition[string]
	FeedId() FeedIdCap
	FullReport() workflows.CapDefinition[string]
	ObservationTimestamp() workflows.CapDefinition[int]
	ReportContext() workflows.CapDefinition[string]
	Signatures() workflows.CapDefinition[[]string]
	private()
}

// FeedCapFromStep should only be called from generated code to assure type safety
func FeedCapFromStep(w *workflows.WorkflowSpecFactory, step workflows.Step[Feed]) FeedCap {
	raw := step.AddTo(w)
	return &feed{CapDefinition: raw}
}

type feed struct {
	workflows.CapDefinition[Feed]
}

func (*feed) private() {}
func (c *feed) BenchmarkPrice() workflows.CapDefinition[string] {
	return workflows.AccessField[Feed, string](c.CapDefinition, "BenchmarkPrice")
}
func (c *feed) FeedId() FeedIdCap {
	return FeedIdCap(workflows.AccessField[Feed, FeedId](c.CapDefinition, "FeedId"))
}
func (c *feed) FullReport() workflows.CapDefinition[string] {
	return workflows.AccessField[Feed, string](c.CapDefinition, "FullReport")
}
func (c *feed) ObservationTimestamp() workflows.CapDefinition[int] {
	return workflows.AccessField[Feed, int](c.CapDefinition, "ObservationTimestamp")
}
func (c *feed) ReportContext() workflows.CapDefinition[string] {
	return workflows.AccessField[Feed, string](c.CapDefinition, "ReportContext")
}
func (c *feed) Signatures() workflows.CapDefinition[[]string] {
	return workflows.AccessField[Feed, []string](c.CapDefinition, "Signatures")
}

func NewFeedFromFields(
	benchmarkPrice workflows.CapDefinition[string],
	feedId FeedIdCap,
	fullReport workflows.CapDefinition[string],
	observationTimestamp workflows.CapDefinition[int],
	reportContext workflows.CapDefinition[string],
	signatures workflows.CapDefinition[[]string]) FeedCap {
	return &simpleFeed{
		CapDefinition: workflows.ComponentCapDefinition[Feed]{
			"benchmarkPrice":       benchmarkPrice.Ref(),
			"feedId":               feedId.Ref(),
			"fullReport":           fullReport.Ref(),
			"observationTimestamp": observationTimestamp.Ref(),
			"reportContext":        reportContext.Ref(),
			"signatures":           signatures.Ref(),
		},
		benchmarkPrice:       benchmarkPrice,
		feedId:               feedId,
		fullReport:           fullReport,
		observationTimestamp: observationTimestamp,
		reportContext:        reportContext,
		signatures:           signatures,
	}
}

type simpleFeed struct {
	workflows.CapDefinition[Feed]
	benchmarkPrice       workflows.CapDefinition[string]
	feedId               FeedIdCap
	fullReport           workflows.CapDefinition[string]
	observationTimestamp workflows.CapDefinition[int]
	reportContext        workflows.CapDefinition[string]
	signatures           workflows.CapDefinition[[]string]
}

func (c *simpleFeed) BenchmarkPrice() workflows.CapDefinition[string] {
	return c.benchmarkPrice
}
func (c *simpleFeed) FeedId() FeedIdCap {
	return c.feedId
}
func (c *simpleFeed) FullReport() workflows.CapDefinition[string] {
	return c.fullReport
}
func (c *simpleFeed) ObservationTimestamp() workflows.CapDefinition[int] {
	return c.observationTimestamp
}
func (c *simpleFeed) ReportContext() workflows.CapDefinition[string] {
	return c.reportContext
}
func (c *simpleFeed) Signatures() workflows.CapDefinition[[]string] {
	return c.signatures
}

func (c *simpleFeed) private() {}

type FeedIdCap workflows.CapDefinition[FeedId]