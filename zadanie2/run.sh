DOCKER_IMAGE="scala:latest" 
NGROK_CONFIG_PATH="$HOME/.config/ngrok/ngrok.yml"

if [[ ! -f "$NGROK_CONFIG_PATH" ]]; then
    echo "Error: ngrok configuration file not found at $NGROK_CONFIG_PATH"
    exit 1
fi

docker run --rm -p 9000:9000 -v ./play-scala:/app -v $NGROK_CONFIG_PATH:/root/.config/ngrok/ngrok.yml $1 $DOCKER_IMAGE
