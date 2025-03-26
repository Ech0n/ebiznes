NGROK_CONFIG_PATH="$HOME/.config/ngrok/ngrok.yml"

IMAGE_NAME="scala" 
IMAGE_TAG="latest"

if [ ! -z "$1" ]; then
    IMAGE_NAME="$1"
fi

if [ ! -z "$2" ]; then
    IMAGE_TAG="$2"
fi

if docker images --format '{{.Repository}}:{{.Tag}}' | grep -q "^${IMAGE_NAME}:${IMAGE_TAG}$"; then
    echo "Image $IMAGE_NAME:${IMAGE_TAG} already exists. Skip docker build."
else
    echo "Image $IMAGE_NAME:${IMAGE_TAG} does not exist."
    docker buildx build -t $IMAGE_NAME:${IMAGE_TAG} .
fi

CONTAINER_ID=$(docker run --rm -d -p 9000:9000 $IMAGE_NAME:${IMAGE_TAG})
cleanup() {
    echo "Stopping container..."
    docker stop $CONTAINER_ID
    echo "Container stopped."
}
trap cleanup EXIT 

ngrok http 9000