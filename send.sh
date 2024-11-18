#!/bin/bash

# Set your bot token and chat ID
BOT_TOKEN="6909524608:AAGzJ5ScTz96pjBxetUZdiJoZ6jMYN1xJz0"  # Replace with your bot token
CHAT_ID="5854985079"      # Replace with your chat ID


#for ((i = 0; i <= 500; i++)) 
#do
MESSAGE="test limit"
# Telegram API URL
URL="https://api.telegram.org/bot$BOT_TOKEN/sendMessage"

# Sending the message via Telegram Bot API using curl
curl -s -X POST $URL -d chat_id=$CHAT_ID -d text="$MESSAGE" -o /dev/null -w ""
#done
