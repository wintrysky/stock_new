#!/bin/sh

docker run -itd \
  -v /etc/timezone:/etc/timezone:ro \
  -v /etc/localtime:/etc/localtime:ro \
  -v `pwd`/config:/data/app/config \
  -v `pwd`/logs:/data/app/logs \
  -p 8020:8020 \
  --env LOG_LEVEL=info \
  --name stocks \
  --network host \
  --restart=always \
  stocks
