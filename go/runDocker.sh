CONTAINER_NAME="go_my_simple_container"

if [[ "$(docker images -q $CONTAINER_NAME:latest 2> /dev/null)" == "" ]]; then
    echo "Docker image not found. Building the image..."
    docker buildx build -t $CONTAINER_NAME .
else
    echo "Docker image already exists."
fi

echo "starting docker container"
echo "this make take a while..."
docker run -it -p 8080:8080 --rm -v $(pwd)/app:/app $CONTAINER_NAME
