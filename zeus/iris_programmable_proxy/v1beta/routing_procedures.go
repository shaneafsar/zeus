package iris_programmable_proxy_v1_beta

import (
	"time"

	"github.com/phf/go-queue/queue"
	iris_operators "github.com/zeus-fyi/zeus/pkg/iris/operators"
)

const (
	RequestHeaderRoutingProcedureHeader = "X-Routing-Procedure"

	ResponseHeaderProcedureLatency = "X-Procedure-Latency-Milliseconds"
)

// pre-canned routing procedures for QuickNode marketplace users

const (
	MaxBlockAggReduce = "max-block-agg-reduce"
)

type IrisRoutingProcedure struct {
	Name string `json:"name"`

	OrderedSteps *queue.Queue `json:"steps"`
}

type BroadcastInstructions struct {
	RoutingPath  string        `json:"routingPath"`
	RestType     string        `json:"restType"`
	Payload      any           `json:"payload,omitempty"`
	MaxDuration  time.Duration `json:"maxRuntime,omitempty"`
	MaxTries     int           `json:"maxTries,omitempty"`
	RoutingTable string        `json:"routingTable"`
	FanInRules   *FanInRules   `json:"fanInRules,omitempty"`
}

type FanInRules struct {
	Rule BroadcastRules `json:"rule"`
}

type BroadcastRules string

const (
	FanInRuleFirstValidResponse = "returnOnFirstSuccess"
	FanInRuleReturnAllResponses = "returnAllSuccessful"
)

// ReturnFirstResultOnSuccess returns the first result from the fan-in that is not an error
func (b BroadcastRules) ReturnFirstResultOnSuccess() string {
	return FanInRuleFirstValidResponse
}

// ReturnResultsOnSuccess returns all results from the fan-in that are not errors that complete before any timeouts occur
// this is the default behavior
func (b BroadcastRules) ReturnResultsOnSuccess() string {
	return FanInRuleReturnAllResponses
}

type IrisRoutingProcedureStep struct {
	BroadcastInstructions BroadcastInstructions                   `json:"broadcastInstructions,omitempty"`
	TransformSlice        []iris_operators.IrisRoutingResponseETL `json:"transformSlice,omitempty"`
	AggregateMap          map[string]iris_operators.Aggregation   `json:"aggregateMap,omitempty"`
}

func (r *IrisRoutingProcedureStep) Aggregate() error {
	if len(r.AggregateMap) == 0 {
		return nil
	}
	for _, v := range r.TransformSlice {
		agg, ok := r.AggregateMap[v.ExtractionKey]
		if !ok {
			continue
		}
		err := agg.AggregateOn(v.Value, v)
		if err != nil {
			return err
		}
		r.AggregateMap[v.ExtractionKey] = agg
	}
	return nil
}
