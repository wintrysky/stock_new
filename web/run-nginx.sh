#/bin/bash
docker run --name web -p 8030:80 \
-v /etc/localtime:/etc/localtime:ro \
-v /mnt/data01/web/web/dist:/usr/share/nginx/html \
-v /mnt/data01/web/web/docker/nginx.conf:/etc/nginx/conf.d/default.conf  \
-d nginx