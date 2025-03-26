echo "Sending request from allowed domain"
resp1=$(curl -s -i -X OPTIONS http://localhost:9000/cart \
     -H "Origin: http://allowedDomain.com" \
     -H "Access-Control-Request-Method: GET" | head -n 1)
echo $resp1
echo
echo "Sending request from NOT allowed domain"
resp2=$(curl -s -i -X OPTIONS http://localhost:9000/cart \
    -H "Origin: http://notallowedDomain.com" \
    -H "Access-Control-Request-Method: GET" | head -n 1)
echo $resp2

