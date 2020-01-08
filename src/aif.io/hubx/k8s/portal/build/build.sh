IMAGE_NAME=doc.hubx.site/hubx/portal:latest
docker build -t $IMAGE_NAME .
docker push $IMAGE_NAME
