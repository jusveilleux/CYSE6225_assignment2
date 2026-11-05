set -e

docker build -f Dockerfile.test -t cyse6225_tests .
docker run --rm cyse6225_tests