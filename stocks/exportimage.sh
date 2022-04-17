#/bin/bash

IMAGE_NAME=stocks
IMAGE_TAG=stocks
LOG_FILG=../build.log
OUTPUT_FOLDER=../out

# check if EXE file exist
if [ ! -f `pwd`"/bin/stocks" ]
then
	echo "file stocks does not exist" >> ${LOG_FILG}
	exit 0
fi

# check if Dockerfile exist
if [ ! -f `pwd`"/Dockerfile" ]
then
	echo "Dockerfile does not exist" >> ${LOG_FILG}
	exit 0
fi

echo ${OUTPUT_FOLDER}
echo "image name is ${IMAGE_NAME}" >> ${LOG_FILG}
echo "IMAGE_TAG is ${IMAGE_TAG}" >> ${LOG_FILG}

FILE_NAME=${IMAGE_TAG}.tar
DOCKER_FULL_NAME=${IMAGE_NAME}":"${IMAGE_TAG}

echo "FILE_NAME is ${FILE_NAME}" >> ${LOG_FILG}
echo "DOCKER_FULL_NAME is ${DOCKER_FULL_NAME}" >> ${LOG_FILG}

echo "building Dockerfile..." >> ${LOG_FILG}
# docker build --rm --build-arg exename=${IMAGE_NAME} -t ${IMAGE_NAME}:${IMAGE_TAG} .
docker build --rm -t ${IMAGE_NAME}:${IMAGE_TAG} .

echo "saving docker file to ${OUTPUT_FOLDER}/${FILE_NAME}" >> ${LOG_FILG}
docker save -o ${OUTPUT_FOLDER}/$FILE_NAME $DOCKER_FULL_NAME

echo "DONE" >> ${LOG_FILG}