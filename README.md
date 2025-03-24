# Go Sample Binary for AWS ECS

간단한 Go 웹 서버 애플리케이션으로, AWS ECS와 ECR을 연습하기 위한 샘플 프로젝트입니다.

## 기능

- `/health` 엔드포인트에 접속하면 `{"message": "SUCCESS"}` JSON 응답을 반환합니다.
- 8081 포트에서 서비스됩니다.

## 특징

- 경량화된 Docker 이미지 (scratch 베이스 이미지 사용)
- 정적 링킹된 Go 바이너리로 최소 크기 제공
- 멀티 스테이지 빌드로 최적화된 이미지 생성

## 로컬에서 실행하기

```bash
# Go 애플리케이션 직접 실행
go run main.go

# 또는 빌드 후 실행
go build -o server .
./server
```

## Docker 빌드 및 실행

```bash
# Docker 이미지 빌드
docker build -t go-sample-binary .

# Docker 컨테이너 실행
docker run -p 8081:8081 go-sample-binary
```

## AWS ECR에 이미지 푸시하기

```bash
# AWS ECR 로그인 (AWS CLI가 설치되어 있어야 합니다)
aws ecr get-login-password --region <your-region> | docker login --username AWS --password-stdin <your-account-id>.dkr.ecr.<your-region>.amazonaws.com

# 이미지 태그 지정
docker tag go-sample-binary:latest <your-account-id>.dkr.ecr.<your-region>.amazonaws.com/go-sample-binary:latest

# ECR에 이미지 푸시
docker push <your-account-id>.dkr.ecr.<your-region>.amazonaws.com/go-sample-binary:latest
```

## 테스트하기

서버가 실행된 후, 다음 명령으로 테스트할 수 있습니다:

```bash
curl http://localhost:8081/health
```

예상 응답:

```json
{ "message": "SUCCESS" }
```
