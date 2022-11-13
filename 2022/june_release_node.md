### 20220-06-04

- MSA, 추가 기능 확장을 고려한 DB 테이블 설계
  [box] table

### 20220-06-07

- 시퀄라이즈 bulkCreate
  - ignoreDuplicated
    - pk 중복된 값이 있으면 insert 하지않는 옵션
    - default = false
- 시퀄라이즈 include insert 이슈
  - 쿼리 한번에 자식 여러개를 insert 하는게 아니라 insert 쿼리를 자식수만큼 요청
- orm 단에서 unique 를 복수로 설정해도 실제 db 에서 적용이 안돼고 orm 단에서 적용
  - string fk 설정도 마찬가지로 orm 단에서 적용되는 형태

### 2022-06-08

- 사내 msa [ grpc || kafka ] 이슈
  - grpc
    - 레거시 레포에 gprc 를 적용해서 payment,member ...등등 필요한 함수 호출이 필요했는데 동시에 rest API 도 되어야했다.
      그럴려면 rest APi 가 되게끔 다 적용시켜줘야되며 proto 파일도 sync 를 항상 맞춰주어야했다.
      너무 헤비해지는 경향이 없지않아있어 중단
  - kafka
    - 필요한 데이터를 큐로 날리고 response로 받아와서 다음 로직을 실행해야되는데 성공했는지 확인할려면 큐날리곳에서 다시 해당 함수로 큐를 날려 콘슘으로 메세지 상태를 받아 다음 로직이
      실행 되게끔 해야했다. 이 방법 역시 헤비해져 잠시 보류중...
- 토스 slash21
  - 토스를 구성하는 서버기술
    - https://toss.im/slash-21/sessions/1-3
  - 토스 테스트 커버리지 100%
    - https://toss.im/slash-21/sessions/1-6
    - 테스트 400개 이상 초과시 1분 이 넘어갔다.
    - 컨텍스트 로딩 제거로 속도 최적화
      - 모킹 라이브러리를 이용하여 실제 애플리케이션을 호출하지않고 모킹 애플리케이션 이용하여 최적화
        - 1600개에서 1분 초과
    - 100% 유지하는 이유
      - 새로 추가한 코드가 조금이라도 커버되지 안으면 언제나 실패
    - 테스트 코드 잘못 짜져있는걸 체크하는 방법
      - 프로덕션 코드를 무작위로 조작하여 테스트가 통과하면 테스트케이스가 부족한것

### 2022-06-09

- 토스 slash21
  - 토스 뱅크 데이터 설계 사상
    - https://toss.im/slash-21/sessions/2-4
    - 데이터양이 적으면 성는상 이슈가 없기때문에 통합하는게 좋다
    - 데이터 양이 많으면 슈퍼/서브 테이블로 분리하는게 좋다
    - 고민하고 만들어라
    - 데이터 양이 많으면 슈퍼/서브 테이블로 분리하는게 좋다
  - 토스 서버 인프라 모니터링
    - https://toss.im/slash-21/sessions/1-2
    - 236개의 서버를 띄우고있다. 컨테이너 기반의 오케스트레이션으로 운영
    - AADC [ Active-Active IDC(inernet data center)] 로 구현
      - DC DR => 주 데이터센터가 있고 제해복구센터가 있는 형태 [데이터센터 : 모든 서버를 모아둔곳]
      - Active-Active 둘다 가동시켜 트래픽 분산 시키고 한쪽이 죽어도 사용자 입장에서는 알 수 가 없다.
      - Active-standby 한쪽만 가동시키고 한쪽은 중지 DC 가 죽으면 DR이 가동되는데 콜드스타트가 발생하여 잠깐 지연이있다.
    - API Gateway 를 사용시 api Gateway 에서 장애가 나면 전면 장애로 전파가 되었다.
      - 서비스 메쉬는 분산 API Gateway 로 서버사이드 로드 밸런싱이아니라 스스로 밸런싱하는 클라이언트 사이드 로드 밸런싱이다보니까
        하나의 API gateway 에서 장애가 발생해도 영향이 적다.
        또한 사이드카 패턴으로 구현을 하여 서비스의 인바운드 아웃바운드 통제권을 가져가기때문에 mtls, 서킷브레이커, 리트라이 로직을 작성할 필요가 없도록
        애플리케이션 코드에서 분리할수있음 네트워크 통신을 전체적으로 오케스트레이션 할수잇는 점이 매력이다.
        <img width="679" alt="스크린샷 2022-06-09 오후 10 02 20" src="https://user-images.githubusercontent.com/65056776/172853323-5e542ecd-b491-4f69-a618-fa88651d6f2e.png">
    - 프로 메테우스의 단점
      - 스스로 어플리케이션들의 메트릭을 스크랩해가는 방식이다 보니 죽으면 메트릭 수집이 중단 되며 누락이된다.
      - 메모리 이슈가 있다. 수집되는 메트릭 량과 저장하는 포멧인 tsdb가 로드되는 양과 비례됨
        해결
      - 프로 메테우스가 죽지않도록 메모리 이슈를 해결하는것
        1. 수집하는 메트릭 량을 줄이고 tsdb사이즈 줄이기
        2. 메트릭들의 카디널리티가 적합한지 판단하여 쿼리 카운트를 통해 필요 없는 메트릭은 없애고 카디널리티가 넓은 메트릭은 어그리게이션을 통해
           축소하여 메트릭양을 줄였다
           결론
        - 프로메테우스를 스켈아웃하여 수집하는 메트릭량을 축소
        - hashmod 라는 옵션으로 샤딩 기능 제공
          - 타겟을 분리해서 쌓을 수 있따.
          - 타겟 ip 를 프로메테우스 개수만큼 가져갈수있다. 부하가 골고루 분산됨

### 2022-06-10

- 토스 slash

### 2022-06-11

- 사내 프로젝트 dockerfile 작성
  - 고려사항
    - env 파일 구성 방법이 정해지지 않아 고민을 많이했다.
      보안과 관리가 용이한 aws 파라미터 스토어를 적용해보았고 해당 방법으로 env 를 관리한다고 했을때 docker file 을 작성해보았다.
      - 장점
        - 여러 서버에서 각각 Env 파일을 구성해 줄 필요가 없다.
        - 중요한 키의 경우 암호화를 시키는데 보안이 더욱 강화 돼 그럴 필요가 없어졌다.
          - 진짜 중요하고 주마다 변경시켜줘야되는 키의 경우 aws 에서 제공하는 서비스 시크릿 매니저 가 있는 그 서비스를 쓰면 된다.

### 2022-06-13

- msk 적용 하는 과정
  - 로컬에서 연동 하는 부분에서 sasl 을 통해 인증하는 부분에서 kafkajs로는 지원이 안되고있는 이슈가있다
    [msk cluster 가 같은 vpc 내부에 있지않으면 통신이 안되는 이슈]
    - 해결
      ec2를 msk cluster 와 동일한 vpc 로 설정하여 생성한다.
      openVPN 도커이미지로 띄워 실행
      이제 openVPN client 를 이용해 접속 가능
      그러면 이제 vpn 을 통해 현재 로컬에서도 vpc 내부에 있는 것처럼 작동하게 되어 접속이 가능해진다.
  - topic auto create err || replicatoin invaild err
    - 해결
      cluster configuration 에서 해당 설정들을 자신이 원하는 작업환경에 맞게 설정한다
      auto.topic... => ture
      default.replication... => broker 숫자에 맞게

### 2022-06-14

- typeORM Upsert 이슈
  - confict 조건을 걸어도 데이터가 계속 새로 생성됨
    - 아마 pk 를 조건으로 안걸어서 계속 생성된거같다. 한블로그에서도 update 조건에 pk 가 명시되어야 된다고 나와있다.

### 2022-06-15

- typeORM upsert 이슈
  - pk 로 조건을 걸어도 upsert 가 안돼고 자꾸 데이터를 생성하는 이슈 발생
    음.. 한 예제 에서는 find를 하고 난후 해당 데이터를 disable 걸고 upsert 하는 방식으로 하고있는데 이게 과연 Upsert 인가...
    upsert 는 데이터가 없으면 생성하고 있으면 수정인데 find를 하면 당연히 없을태고 해당 데이터를 upsert 에 넣는다???
    현재 아무런 방법으로도 되질않아 find => update 없으면 create 로 직접 upsert를 구현하였다.

### 2022-06-16

- 사내 프로젝트
  - 레거시에서 배송 서비스를 분리하여 서버 간 카프카 통신을 적용중이다. 현재 레거시에 배송 서비스 분리 후 카프카 통신 적용 및 rest API로 shipping 데이터 조회 기능을
    마침내 다 적용 시켰다... 일주일 정도 걸린것 같다.
    이벤트 구조에 대해 설계를 하였고 서버 가 스켈 아웃 됐을시 consumer가 동시에 메세지를 소비 하지않을까라는 고려 사항도 있었다.
    하지만 컨슈머가 파티션을 소모하는 원리인데 한 파티션을 두개이상의 컨슈머가 소비할수없다 라는 카프카 원칙이라는 이론이 있어 테스트를 해서 검증을 해야되는 상태이며
    그밖에는 작업이 잘 이루어 지고 있는 것 같다...
    백엔드 인원이 리더 , 나 포함 2명 이라서 상당히 헤비하고 빡세다...

### 2022-06-17

- 사내
  shipping 서버 구축

### 2022-06-18

- msa 기반 인프라 아키텍쳐 기술 검토
  - API gateway - ecs

### 2022-06-19

- ECR-ECS 구축

### 2022-06-20

- ECS ci/cd 자동화 배포 검토

### 2022-06-21 ~ 22

- ECS ci/cd dev 서버 자동화 배포 적용 [msa]
  - shipping
    이중 인증 => label 이 shipping 일떄 와 /apps/shipping/\*\* shipping 디렉토리 내에 수정사항이 있을때
    1차적으로 수정내역으로 보고 action 이 활성화 된다.
    2차적으로 label 이 어떤 서비스이냐에 따라 걸러진다.
  - github actions [ env 파일 생성 방법 ]
    ```
     steps:
    - name: "Say Hello Mona it's Monday"
      run: echo "AWS_SECRET_ID=$AWS_SECRET_ID_TEST" >> .env
      env:
        AWS_SECRET_ID_TEST: ${{ secrets.AWS_SECRET_ID_TEST }}
    ```
    github 에 저장되지 않는 env 파일은 secret 에 환경변수를 넣어 놓고 gitaction 런타임 상 이런식으로 env 파일을 생성할 수가 있다.

### 2022-06-23

- nodejs kafka 한 소비자 그룹내에 여러 토픽을 구독해서 이벤트를 받는 방법 검토

### 2022-06-24

- nodejs kafka 한 소비자 그룹내에 여러 토픽을 구독
  nestjs 에서 eventPattern 데코레이터로 각각 함수로 받을 수 있는 방법은 metadata 를 이용하는 방법이므로 가능하지만
  nodejs 에서는 해당 방법으로 구현이안된다
  run 함수내에 topic 별로 받는 방법을 객체 리터럴 패턴으로 구현하여 이벤트를 받게 되었다.
  https://github.com/tulios/kafkajs/issues/1040

### 2022-06-25

- typeORM transaction 관리 이슈
  - 시퀄라이즈의 경우 transaction 관리를 미들웨어로 만들어서 각 쿼리 인자에 transaction 을 넘겨주면 자동으로 transaction 이 수행된다. 롤백이나 커밋
  - typeORM 은 transaction 의 경우 queryRunner 로 쿼리를 짜서 트랜잭션을 수행하여야되기때문에 transaction 이 필요한 함수에만 queryRunner 를 써야되는 이슈가 있다.

### 2022-06-26

- nestjs decorater 란
  데코레이터를 잘 사용하면 횡단관심사를 분리하여 관점 지향 프로그래밍을 적용한 코드를 작성할 수 있습니다.
  [AOP]
  횡단관심사의 분리를 허용함으로써 모듈성을 증가시키는 것이 목적인 프로그래밍 패러다임

### 2022-06-27

- 사내 nestjs msa core 기능 통합
  logging, auth, exception , interceptor 를 여러 서비스에서 공통으로 쓰기에 기능 통합을 하였고 현재는 해당 core를 최상위 디렉토리에 올려 같이 쓰기에는 gitaction yml 파일도
  수정이 필요하기에 한 서비스에서 페어코딩으로 맞추고 복붙으로 갖다쓰기로 결정

### 2022-06-28

- typeorm left join count 이슈
  - join count 를 가져올시 Getmany 로 들고오지 못하고 getRowMany 로 들고와야 되는 이슈
    원시 데이터로 뽑아야 가져올수있다.

### 2022-06-29

- shipping server Prod, stage ECS CI/CD 구축 준비
  taskDefinition 작성 및 gitAction yml 파일 수정

## todo
