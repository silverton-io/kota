Kota is an Okta System Log Collection and Distribution system.

**It interfaces the four ways of collecting System Logs from your Okta Tenant:**

- [AWS EventBridge Log Streaming](https://help.okta.com/en-us/content/topics/reports/log-streaming/add-aws-eb-log-stream.htm)
- [Splunk Cloud Log Streaming](https://help.okta.com/en-us/content/topics/reports/log-streaming/add-splunk-log-stream.htm)
- [Event Hooks](https://developer.okta.com/docs/concepts/event-hooks/)
- [System Log API](https://developer.okta.com/docs/reference/api/system-log/)


**with downstream systems your company uses:**

- [Datadog](https://www.datadoghq.com/)
- [Elasticsearch](https://www.elastic.co/)
- [Splunk (Self-Hosted)](https://www.splunk.com/)
- [Grafana](https://grafana.com/)
- [InfluxDB](https://www.influxdata.com/)
- [NewRelic](https://newrelic.com/)
- [Segment](https://segment.com/)
- [Mixpanel](https://mixpanel.com/)
- [Sumo Logic](https://www.sumologic.com/)


**Kota includes support for object stores so you can durably persist more than 90 days of your Okta System Log data to cheap storage:**

- [Amazon S3](https://aws.amazon.com/s3/)
- [Google Cloud Storage](https://cloud.google.com/storage)
- [Azure Blob Storage](https://azure.microsoft.com/en-us/products/storage/blobs)
- [MinIO](https://min.io/)˛


It will eventually run an [embedded OLAP database](https://duckdb.org/) for rolling up statistics, aggregating events, filtering, security detections, and more.
