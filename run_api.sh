
set -e

docker build -t cyse6225_assignment_2 .
docker run --rm -p 8000:8000 -e SLACK_WEBHOOK_URL="${SLACK_WEBHOOK_URL:-}" cyse6225_assignment_2