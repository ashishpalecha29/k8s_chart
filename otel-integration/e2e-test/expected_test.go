package e2e

var expectedSchemaURL = map[string]bool{
	"https://opentelemetry.io/schemas/1.6.1": false,
	"https://opentelemetry.io/schemas/1.9.0": false,
}

const expectedScopeVersion = "0.111.0"

var expectedScopeNames = map[string]bool{
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/networkscraper":    false,
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper":        false,
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper": false,
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/memoryscraper":     false,
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/loadscraper":       false,
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper":       false,
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper":    false,
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver":                                   false,
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver":                                     false,
}

var expectedResourceAttributesKubeletstatreceiver = map[string]string{
	"k8s.pod.uid":              "",
	"k8s.pod.name":             "",
	"k8s.namespace.name":       "",
	"k8s.container.name":       "",
	"k8s.daemonset.name":       "",
	"k8s.deployment.name":      "",
	"cx.otel_integration.name": "coralogix-integration-helm",
	"k8s.cluster.name":         "otel-integration-agent-e2e",
	"k8s.node.name":            "otel-integration-agent-e2e-control-plane",
	"host.name":                "",
	"host.id":                  "",
	"os.type":                  "linux",
	"cloud.provider":           "azure",
	"cloud.platform":           "",
	"cloud.region":             "",
	"cloud.account.id":         "",
	"azure.vm.name":            "",
	"azure.vm.size":            "",
	"azure.vm.scaleset.name":   "",
	"azure.resourcegroup.name": "",
}

var expectedResourceAttributesHostmetricsreceiver = map[string]string{
	"cx.otel_integration.name": "coralogix-integration-helm",
	"k8s.cluster.name":         "otel-integration-agent-e2e",
	"k8s.node.name":            "otel-integration-agent-e2e-control-plane",
	"host.name":                "",
	"host.id":                  "",
	"os.type":                  "linux",
	"cloud.provider":           "azure",
	"cloud.platform":           "azure_vm",
	"cloud.region":             "",
	"cloud.account.id":         "",
	"azure.vm.name":            "",
	"azure.vm.size":            "",
	"azure.vm.scaleset.name":   "",
	"azure.resourcegroup.name": "",
	"process.pid":              "",
	"process.parent_pid":       "",
	"process.executable.name":  "",
	"process.executable.path":  "",
	"process.command":          "",
	"process.command_line":     "",
	"process.owner":            "",
}

var expectedResourceAttributesPrometheusreceiver = map[string]string{
	"service.name":             "opentelemetry-collector",
	"net.host.name":            "",
	"service.instance.id":      "",
	"net.host.port":            "",
	"http.scheme":              "http",
	"server.address":           "",
	"server.port":              "",
	"url.scheme":               "",
	"service_version":          expectedScopeVersion,
	"cx.otel_integration.name": "coralogix-integration-helm",
	"k8s.cluster.name":         "otel-integration-agent-e2e",
	"k8s.node.name":            "otel-integration-agent-e2e-control-plane",
	"host.name":                "",
	"host.id":                  "",
	"k8s.pod.ip":               "",
	"k8s.pod.name":             "",
	"k8s.deployment.name":      "",
	"k8s.namespace.name":       "",
	"k8s.daemonset.name":       "",
	"os.type":                  "linux",
	"cloud.provider":           "azure",
	"cloud.platform":           "",
	"cloud.region":             "",
	"cloud.account.id":         "",
	"azure.vm.name":            "",
	"azure.vm.size":            "",
	"azure.vm.scaleset.name":   "",
	"azure.resourcegroup.name": "",
	"k8s_node_name":            "",
}

var expectedMetrics map[string]bool = map[string]bool{
	"system.network.connections":                     false,
	"system.network.dropped":                         false,
	"system.network.errors":                          false,
	"system.network.io":                              false,
	"system.network.packets":                         false,
	"system.cpu.time":                                false,
	"system.filesystem.inodes.usage":                 false,
	"system.filesystem.usage":                        false,
	"system.memory.usage":                            false,
	"system.memory.utilization":                      false,
	"system.cpu.load_average.15m":                    false,
	"system.cpu.load_average.1m":                     false,
	"system.cpu.load_average.5m":                     false,
	"system.cpu.utilization":                         false,
	"system.disk.io":                                 false,
	"system.disk.io_time":                            false,
	"system.disk.merged":                             false,
	"system.disk.operation_time":                     false,
	"system.disk.operations":                         false,
	"system.disk.pending_operations":                 false,
	"system.disk.weighted_io_time":                   false,
	"k8s.node.cpu.time":                              false,
	"k8s.node.cpu.utilization":                       false,
	"k8s.node.filesystem.available":                  false,
	"k8s.node.filesystem.capacity":                   false,
	"k8s.node.filesystem.usage":                      false,
	"k8s.node.memory.available":                      false,
	"k8s.node.memory.major_page_faults":              false,
	"k8s.node.memory.page_faults":                    false,
	"k8s.node.memory.rss":                            false,
	"k8s.node.memory.usage":                          false,
	"k8s.node.memory.working_set":                    false,
	"k8s.node.network.errors":                        false,
	"k8s.node.network.io":                            false,
	"k8s.pod.cpu.time":                               false,
	"k8s.pod.cpu.utilization":                        false,
	"k8s.pod.filesystem.available":                   false,
	"k8s.pod.filesystem.capacity":                    false,
	"k8s.pod.filesystem.usage":                       false,
	"k8s.pod.memory.available":                       false,
	"k8s.pod.memory.major_page_faults":               false,
	"k8s.pod.memory.page_faults":                     false,
	"k8s.pod.memory.rss":                             false,
	"k8s.pod.memory.usage":                           false,
	"k8s.pod.memory.working_set":                     false,
	"k8s.pod.network.errors":                         false,
	"k8s.pod.network.io":                             false,
	"container.cpu.time":                             false,
	"container.cpu.utilization":                      false,
	"container.filesystem.available":                 false,
	"container.filesystem.capacity":                  false,
	"container.filesystem.usage":                     false,
	"container.memory.available":                     false,
	"container.memory.major_page_faults":             false,
	"container.memory.page_faults":                   false,
	"container.memory.rss":                           false,
	"container.memory.usage":                         false,
	"container.memory.working_set":                   false,
	"otelcol_process_memory_rss":                     false,
	"otelcol_receiver_accepted_metric_points":        false,
	"otelcol_exporter_queue_capacity":                false,
	"otelcol_otelsvc_k8s_ip_lookup_miss":             false,
	"otelcol_process_runtime_total_alloc_bytes":      false,
	"otelcol_otelsvc_k8s_pod_added":                  false,
	"otelcol_process_runtime_total_sys_memory_bytes": false,
	"otelcol_process_uptime":                         false,
	"otelcol_processor_accepted_metric_points":       false,
	"otelcol_processor_batch_metadata_cardinality":   false,
	"otelcol_receiver_refused_log_records":           false,
	"otelcol_receiver_refused_metric_points":         false,
	"otelcol_exporter_queue_size":                    false,
	"otelcol_exporter_send_failed_metric_points":     false,
	"otelcol_exporter_sent_metric_points":            false,
	"otelcol_exporter_sent_log_records":              false,
	"otelcol_otelsvc_k8s_pod_table_size":             false,
	"otelcol_otelsvc_k8s_pod_updated":                false,
	"otelcol_scraper_errored_metric_points":          false,
	"otelcol_process_cpu_seconds":                    false,
	"otelcol_process_runtime_heap_alloc_bytes":       false,
	"otelcol_scraper_scraped_metric_points":          false,
	"otelcol_receiver_accepted_log_records":          false,
	"otelcol_processor_batch_timeout_trigger_send":   false,
	"otelcol_exporter_send_failed_log_records":       false,
	"otelcol_processor_batch_batch_send_size":        false,
	"otelcol_fileconsumer_open_files":                false,
	"otelcol_fileconsumer_reading_files":             false,
	"otelcol_processor_outgoing_items":               false,
	"otelcol_processor_incoming_items":               false,
	"scrape_samples_scraped":                         false,
	"scrape_samples_post_metric_relabeling":          false,
	"scrape_series_added":                            false,
	"scrape_duration_seconds":                        false,
	"up":                                             false,
	"process.cpu.time":                               false,
	"process.cpu.utilization":                        false,
	"process.disk.io":                                false,
	"process.memory.usage":                           false,
	"process.memory.virtual":                         false,
	"process.memory.utilization":                     false,
	"process.threads":                                false,
}