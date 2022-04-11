#!/bin/bash
source /etc/profile
javaDir=/data/work/javaDev/topo
pushd $javaDir
mvn clean install
popd
cp -f ${javaDir}/topo-server/target/app.jar $1