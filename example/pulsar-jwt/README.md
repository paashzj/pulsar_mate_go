## interactive
### command
```bash
# create a secret key
bin/pulsar tokens create-secret-key --output /opt/sh/pulsar/secret.key
# generate token
bin/pulsar tokens create --secret-key file:///opt/sh/pulsar/secret.key --subject root
```
### run
```bash
docker-compose run --rm pulsar bash
```