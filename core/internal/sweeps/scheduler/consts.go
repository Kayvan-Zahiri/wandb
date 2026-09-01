package scheduler

const (
	// runsPageSize is how many runs one poll page requests.
	runsPageSize = 200

	// warmStartPageSize bounds the batch of prior runs the optimizer
	// ingests at a time.
	warmStartPageSize = 100

	// historySampleCount is how many rows each run's metric history is
	// sampled down to.
	historySampleCount = 20
)
