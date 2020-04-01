# Azure Event Hubs Service

[Azure Event Hubs](https://azure.microsoft.com/en-us/services/event-hubs/) is a hyper-scale telemetry ingestion service that collects, transforms, and stores millions of events. As a distributed streaming platform, it gives you low latency and configurable time retention, which enables you to ingress massive amounts of telemetry into the cloud and read the data from multiple applications using publish-subscribe semantics.

## Plans

| Plan| Ingress events                | Capture  | Apache Kafka| Extended Retention*|
|--------|----------------------------|----------|-------------|--------------------|
|basic   | $0.028 per million events  |          |             |                    | 
|standard| $0.028 per million events  |$0.10/hour| yes         | $0.12/GB/month     |

## Configuration Options

The following options can be configured across all supported platforms. Notes below document any platform specific information for mapping that might be required.

| Option Name     | Value Type   | Values                      | Default |
|-----------------|--------------|-----------------------------|---------|
| instance_name   | string       | `^[a-z][a-z0-9-]+$`         |vsb-eventhubs-*instance_id* |

#### Azure specific config parameters

| Option Name     | Value Type   | Values                      | Default                |
|-----------------|--------------|-----------------------------|------------------------|
| location        | string       | Azure location of Eventhubs instance | westus |
| resource_group  | string       | Resource group for Eventhubs instance | rg-*instance_name* |
| sku             | string       | basic or standard              |                     |
| auto_inflate_enabled | boolean        |                         |  false              |
| partition_count      | integer        |                         |   2                 |
| message_retention    | integer        |                         |   1                 |

## Binding Credentials

The binding credentials for MySQL have the following shape:

```
 credentials": {
     "event_hub_connection_string": "Endpoint=sb://",
     "event_hub_name": "vsb-eventhubs-",
     "eventhub_name": "vsb-eventhubs-",
     "eventhub_rg_name": "rg-vsb-eventhubs-",
     "namespace_connection_string": "Endpoint=sb://vsb-eventhubs-",
     "namespace_name": "vsb-eventhubs-",
     "shared_access_key_name": "",
     "shared_access_key_value": ""
    }
```




