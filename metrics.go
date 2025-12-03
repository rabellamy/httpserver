package httpserver

import (
	"github.com/rabellamy/promstrap/strategy"
)

func newRED(namespace string) (*strategy.RED, error) {
	red, err := strategy.NewRED(strategy.REDOpts{
		Namespace: namespace,
		RequestsOpt: strategy.REDRequestsOpt{
			RequestType:   "http",
			RequestLabels: []string{"path", "verb"},
		},
		ErrorsOpt: strategy.REDErrorsOpt{
			ErrorLabels: []string{"error"},
		},
		DurationOpt: strategy.REDDurationOpt{
			DurationLabels: []string{"path"},
		},
	})

	if err != nil {
		return nil, err
	}

	return red, nil
}
