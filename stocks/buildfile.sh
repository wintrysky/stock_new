#/bin/bash

# clean up
rm -rf `pwd`/bin

mkdir -p `pwd`/bin

chmod +x `pwd`/build.sh

docker run  \
    -w /data/src \
    -v `pwd`:/data/src \
    -v `pwd`/bin:/data/bin/ \
    golang:1.16 /bin/bash -c '/data/src/build.sh'