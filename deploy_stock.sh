#!/bin/bash

## 判断当前路径下(/root/upload)是否存在需要部署的zip
deploy_service()
{
    cd /root/upload
	
    container=$1
    zipFolder=$2
    sourceFolder=$3
    zipFileName=$4
    dockerRun=$5
	
	cp ${zipFileName} /root/backup
	mkdir -p ${zipFolder}
    # stop container
    containerID=`docker ps -a | grep ${container} | awk '{print $1}'`
    if [ -n ${containerID} ]; then
        echo "stopping container ${containerID}"
        docker stop ${containerID}
        echo "removing container"
        docker rm ${containerID}
    fi
    sleep 3
    # rm image
    imageID=`docker images | grep stocks | awk '{print $3}'`
    if [ -n ${imageID} ]; then
        echo "removing image ${imageID}"
        docker rmi ${imageID}
    fi
    # build 
    mv ${zipFileName} ${zipFolder}
    cd ${zipFolder}
	echo "cd ${zipFolder}"
    7za x ${zipFileName} -aoa
    echo "cd ${sourceFolder}"
    cd ${sourceFolder}
    sh buildfile.sh
    docker build . -t ${container}
    sh ${dockerRun}
    docker ps -a | grep ${container}
}

deploy_web()
{
    cd /root/upload
    zipFolder=$1
    sourceFolder=$2
    zipFileName=$3
    nginxName=$4
	
	cp ${zipFileName} /root/backup
	mkdir -p ${zipFolder}
    echo "unziping "${zipFileName}
    mv ${zipFileName} ${zipFolder}
    cd ${zipFolder}
    7za x ${zipFileName} -aoa
	chmod -R 777 ${zipFolder}
    echo "removing dist"
    #rm -rf ${sourceFolder}/dist/*
    cd ${sourceFolder}
	#npm install
    #npm run build
    echo "starting nginx"
    docker restart ${nginxName}
    docker ps -a | grep ${nginxName}
}

deploy()
{
  fileName=$1
  echo "#########################"
  echo "deploying $1"
  if [ "$fileName" == 'stocks.7z' ]; then
    PROJECT='stocks'
    ZIP_FOLDER='/mnt/data01'
    SOURCE_FOLDER='/mnt/data01/stocks'
    ZIP_FILE_NAME='stocks.7z'
    DOCKER_RUN='docker-run.sh'
    deploy_service "$PROJECT" "$ZIP_FOLDER" "$SOURCE_FOLDER" "$ZIP_FILE_NAME" "$DOCKER_RUN"
  fi   
  if [ "$fileName" == 'web.7z' ]; then
    ZIP_FOLDER='/mnt/data01/web'
    SOURCE_FOLDER='/mnt/data01/web/web'
    ZIP_FILE_NAME='web.7z'
    NGINX_NAME='web'
    deploy_web "$ZIP_FOLDER" "$SOURCE_FOLDER" "$ZIP_FILE_NAME" "$NGINX_NAME"
  fi 
}

rm -rf *.7z
rz -bye

for line in $(ls *.7z)
do
 echo "-----------------"
 echo "deploying $line"
 deploy "$line"
done

