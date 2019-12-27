IMAGE_NAME=registry.yw.zj.chinamobile.com/hubx/portal:latest
docker build -t $IMAGE_NAME .
docker push $IMAGE_NAME
