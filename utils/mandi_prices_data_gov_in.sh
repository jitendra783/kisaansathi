#!/bin/bash

curl -X 'GET' \
  'https://api.data.gov.in/resource/35985678-0d79-46b4-9ed6-6f13308a1d24?api-key=579b464db66ec23bdd000001185c55b29a984645468d0f573fee9052&format=csv&limit=1000' \
  -H 'accept: text/csv' \
  -o mandi_prices.csv
# Check if the file was downloaded successfully
if [ $? -eq 0 ]; then
  echo "Mandi prices data downloaded successfully."
else
  echo "Failed to download mandi prices data."
  exit 1
fi  
