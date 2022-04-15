#!/bin/bash
image=harbor.dev.21vianet.com/cmdb/cmdb_ui:latest
pushd /data/work/go/src/configcenter/deploy_ui/DockerUI || exit
docker build -f Dockerfile -t ${image} /data/work/go/src/configcenter/indc-front/
popd || exit
docker push ${image}