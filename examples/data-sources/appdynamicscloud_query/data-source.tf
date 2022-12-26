data "appdynamicscloud_query" "example" {
  query = "fetch id: id, name: attributes(service.name), cpm: metrics(apm:response_time) {source, timestamp, min, max} from entities(apm:service)[attributes(service.namespace) = 'Levitate'].out.to(apm:service_instance) since -3h"
}