
set -e

docker build -t cyse6225_assignment_2 .
docker run --rm -p 8000:8000 