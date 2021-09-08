## mate
### start 
```bash
docker run -d ttbb/pulsar:mate
```
### start port expose
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 -p 6650:6650 -p 8080:8080 ttbb/pulsar:mate
```
### start daemon
```bash
docker run -p 6650:6650 -p 2181:2181 -p 3181:3181 -p 8080:8080 -p 4181:4181 -d ttbb/pulsar:mate
```
### start daemon 4G
```bash
docker run -p 6650:6650 -p 2181:2181 -p 3181:3181 -p 8080:8080 -p 4181:4181 -m 4G -d ttbb/pulsar:mate
```
## cluster
### run cluster
```bash
docker run -e ZK_ADDR=172.16.0.56:2181,172.16.0.55:2181,172.16.0.54:2181 -d ttbb/pulsar:cluster
```