docker build -t go_api . && \
docker run -d -p 8080:8080 --rm --name go_api go_api