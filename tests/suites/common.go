package suites

import "fmt"

const (
	TestNamespace               = "default"
	DefaultContainerName        = "k8skafka"
	DefaultKafkaStatefulSetName = "kafka-kafka"
	DefaultZkStatefulSetName    = "zk-zookeeper"
	DefaultKudoKafkaInstance    = "kafka"
)

func GetBrokerPodName(id int) string {
	return fmt.Sprintf("%s-%d", DefaultKafkaStatefulSetName, id)
}
