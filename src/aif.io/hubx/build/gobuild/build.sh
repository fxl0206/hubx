export IMAGE_NAME=iseex.picp.io:30500/hubx/go-build:1.0
docker build -t $IMAGE_NAME .
docker push $IMAGE_NAME
