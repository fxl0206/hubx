IMAGE_NAME=iseex.picp.io:30500/hubx/hubxs:latest
docker build -t $IMAGE_NAME .
docker push $IMAGE_NAME
