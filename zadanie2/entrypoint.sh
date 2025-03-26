#!/bin/bash

NGROK_PATH="/root/.config/ngrok/ngrok.yml"

/build/my-play-scala-app-1.0/bin/my-play-scala-app -Dhttp.port=9000

# export NGROK_CONFIG=$NGROK_PATH
# ngrok http 9000 --config="/root/.config/ngrok/ngrok.yml"

# ngrok http 9000 --config=$NGROK_PATH