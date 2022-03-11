## mate
### start
```bash
docker run --rm -e REMOTE_MODE=false ttbb/pulsar:mate
```
### start daemon
```bash
docker run -d --rm -e REMOTE_MODE=false ttbb/pulsar:mate
```
### start port expose
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 -p 6650:6650 -p 8080:8080 --rm -e REMOTE_MODE=false ttbb/pulsar:mate
```
### start with index metadata interceptor
```bash
docker run -p 6650:6650 -p 8080:800 --rm -e PULSAR_BROKER_ENTRY_METADATA_INTERCEPTORS=org.apache.pulsar.common.intercept.AppendIndexMetadataInterceptor -e PULSAR_EXPOSING_BROKER_ENTRY_METADATA_TO_CLIENT_ENABLED=true -e REMOTE_MODE=false ttbb/pulsar:mate
```
### start port expose, tls
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 -p 6650:6650 -p 6651:6651 -p 8080:8080 --rm -e REMOTE_MODE=false -e PULSAR_TLS_ENABLE=true ttbb/pulsar:mate
```
### start daemon port expose
```bash
docker run -p 6650:6650 -p 2181:2181 -p 3181:3181 -p 8080:8080 -p 4181:4181 -d --rm -e REMOTE_MODE=false ttbb/pulsar:mate
```
### start daemon 4G
```bash
docker run -p 6650:6650 -p 2181:2181 -p 3181:3181 -p 8080:8080 -p 4181:4181 -m 4G -d --rm -e REMOTE_MODE=false ttbb/pulsar:mate
```
## cluster
### run cluster
```bash
docker run -d --rm -e CLUSTER_ENABLE=true -e REMOTE_MODE=false -e ZK_ADDR=localhost:2181 ttbb/pulsar:mate
```