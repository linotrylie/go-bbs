package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
// DefaultBuckets prometheus buckets in seconds.
// DefaultBuckets = []float64{0.3, 1.2, 5.0}
)

const (
	reqsName              = "http_requests_total"
	latencyName           = "http_request_duration_seconds"
	ormName               = "orm_requests_total"
	ormlatencyName        = "orm_duration_seconds"
	httpClientReqsName    = "http_client_requests_total"
	httpClientLatencyName = "http_client_duration_seconds"

	kafkaProducerReqsName    = "kafka_producer_requests_total"
	kafkaProducerLatencyName = "kafka_producer_duration_seconds"
)

// Prometheus is a handler that exposes prometheus metrics for the number of requests,
// the latency and the response size, partitioned by status code, method and HTTP path.
//
// Usage: pass its `ServeHTTP` to a route or globally.
type Prometheus struct {
	reqs    *prometheus.CounterVec
	latency *prometheus.HistogramVec
	listen  string

	ormReqs    *prometheus.CounterVec
	ormLatency *prometheus.HistogramVec
	counters   []*prometheus.CounterVec
	histograms []*prometheus.HistogramVec
}

// newPrometheus returns a new prometheus middleware.
//
// If buckets are empty then `DefaultBuckets` are set.
func NewPrometheus() *Prometheus {
	p := &Prometheus{}
	return p
}
func RegisterPrometheus(p *Prometheus, name, listen string) {
	p.listen = listen
	p.reqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        reqsName,
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			ConstLabels: prometheus.Labels{"service": name},
		},
		[]string{"http_code", "code", "method", "path"},
	)
	prometheus.MustRegister(p.reqs)
	p.latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        latencyName,
		Help:        "How long it took to process the request, partitioned by status code, method and HTTP path.",
		ConstLabels: prometheus.Labels{"service": name},
	},
		[]string{"http_code", "code", "method", "path"},
	)
	prometheus.MustRegister(p.latency)

	p.ormReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        ormName,
			Help:        "",
			ConstLabels: prometheus.Labels{"service": name},
		},
		[]string{"result", "model", "method"},
	)
	prometheus.MustRegister(p.ormReqs)

	p.ormLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        ormlatencyName,
		Help:        "",
		ConstLabels: prometheus.Labels{"service": name},
	},
		[]string{"model", "method", "result"},
	)
	prometheus.MustRegister(p.ormLatency)

	for i := 0; i < len(p.counters); i++ {
		prometheus.MustRegister(p.counters[i])
	}
	for i := 0; i < len(p.histograms); i++ {
		prometheus.MustRegister(p.histograms[i])
	}
	return
}

// NewPrometheusHandle .
func NewPrometheusHandle(p *Prometheus) gin.HandlerFunc {
	//go func() {
	//	if strings.Index(p.listen, ":") == 0 {
	//		global.LOG.Info(fmt.Sprintf("[go-bbs] Now prometheus listening on: http://0.0.0.0%s\n", p.listen))
	//	} else {
	//		global.LOG.Info(fmt.Sprintf("[go-bbs] Now prometheus listening on: http://%s\n", p.listen))
	//	}
	//	http.ListenAndServe(p.listen, nil)
	//}()
	return func(ctx *gin.Context) {
		start := time.Now()
		r := ctx.Request
		statusCode := strconv.Itoa(ctx.Writer.Status())
		code := "0"
		path := ctx.FullPath()
		p.reqs.WithLabelValues(statusCode, code, r.Method, path).
			Inc()
		p.latency.WithLabelValues(statusCode, code, r.Method, path).
			Observe(float64(time.Since(start).Nanoseconds()) / 1000000000)
		ctx.Next()
	}
}

func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

// OrmWithLabelValues .
func (p *Prometheus) OrmWithLabelValues(model, method string, e error, starTime time.Time) {
	if p.listen == "" {
		return
	}

	result := "ok"
	if e != nil && !strings.Contains(fmt.Sprint(e), "record not found") {
		result = "error"
	}

	p.ormReqs.WithLabelValues(model, method, result).Inc()
	p.ormLatency.WithLabelValues(model, method, result).Observe(float64(time.Since(starTime).Nanoseconds()) / 1000000000)
}

// RegisterCounter .
func (p *Prometheus) RegisterCounter(conter *prometheus.CounterVec) {
	p.counters = append(p.counters, conter)
}

// RegisterHistogram .
func (p *Prometheus) RegisterHistogram(histogram *prometheus.HistogramVec) {
	p.histograms = append(p.histograms, histogram)
}
