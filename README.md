# Effective Gin

본 프로젝트는 Go 언어의 Gin 프레임워크를 사용하여 효과적인 웹 애플리케이션 구축 방법을 보여줍니다. 프로젝트 구조는 공식 "Effective Go" 문서에 제시된 지침과 모범 사례를 따릅니다.

본 프로젝트는 견고하고 유지보수가 용이한 Gin 기반 애플리케이션 개발을 위한 탄탄한 기반을 제공하는 것을 목표로 합니다. 개발 워크플로우를 개선하고 코드 품질을 보장하기 위한 도구를 포함하고 있습니다.

## 프로젝트 구조

본 프로젝트는 "Effective Go"의 원칙에 따라 관심사 분리 및 개선된 구성을 촉진하는 구조를 따릅니다. 프로젝트의 복잡성에 따라 정확한 구조는 다를 수 있지만, 일반적으로 찾을 수 있는 디렉토리는 다음과 같습니다.

* `cmd`: 주요 애플리케이션 진입점을 포함합니다.
* `internal`: 공개 API로 노출되어서는 안 되는 핵심 애플리케이션 로직을 담습니다. 이 디렉토리는 `handlers`, `errors`, `database` 등과 같은 하위 패키지로 더 나눌 수 있습니다.
* `utils`: 프로젝트에서 사용하는 유틸리티를 포함합니다.
* `docs`: 자동으로 생성된 Swagger API 문서입니다. 변경이 필요할 경우 make 명령어를 사용합니다.

## make 사용법

터미널에서 make [option] 명령어를 실행하여 원하는 작업을 수행합니다.
예: make build , make test, make clean, make check
이 Makefile은 프로젝트의 빌드 및 테스트 프로세스를 자동화하여 개발 생산성을 향상시키고 코드 품질을 유지하는 데 도움이 됩니다.
make check 를 사용하면 모든 로직을 검사하고 초기상태로 돌아갑니다. 
swag를 생성 후 fmt, lint, test를 진행하고 build를 진행해서 문제가 없다면 clean을 진행합니다.