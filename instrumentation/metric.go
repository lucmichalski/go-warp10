package instrumentation

import b "github.com/miton18/go-warp10/base"

// Metric is a thing you can monitor
type Metric interface {
	Name() string
	Help() string
	Get() b.GTSList
	Reset()
}

// Metrics is a collection of Metric
type Metrics []Metric

// Allow to sort metrics by name
func (m Metrics) Len() int {
	return len(m)
}
func (m Metrics) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m Metrics) Less(i, j int) bool {
	if m[i] == nil {
		return false
	}
	if m[j] == nil {
		return true
	}
	return (m[i]).Name() > (m[j]).Name()
}

func (m Metrics) Get() b.GTSList {
	res := b.GTSList{}
	for _, metric := range m {
		res = append(res, metric.Get()...)
	}
	return res
}
