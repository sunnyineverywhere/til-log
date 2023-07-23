# Authentication 객체(org.springframework.security.core.Authentication)에 대해 정리

### 1. Authentication과 Authorization
- Authentication == 인증
  - (특정 리소스에 접근하려고 하는) 사용자의 정보를 통해 본인이 맞는지 확인하는 절차
- Authorization == 인가
  - 인증된 사용자가 요청된 자원에 접근 가능한지 결정하는 절차
  - Authentication(인증) 후, 사용자를 식별 + 권한 부여 가능
 
### 2. 인증 방식
종류에 대해 먼저 기술해보면
1) credential 기반: username, password 기반
2) twofactor: 사용자가 입력한 개인정보를 인증 + 다른 인증 체계를 이용 => 두가지 조합으로 인증(신용카드 등)
3) hardward: 자동차 키와 같은 방식  

Spring Security는 credential 기반

### 3. SecurityContextHolder & SecurityContext
![image](https://github.com/sunnyineverywhere/til-log/assets/80109963/2d07bb30-e7ba-4382-8db6-9f923a997d2e)

**SecurityContextHolder**
- Authentication된 사용자의 상세 정보(SecurityContext)를 저장
- 값을 넣는 방법은 상관하지 않고, 값이 있다면 현재 인증한 사용자 정보로 취급
- 사용자가 인증되었음을 나타내는 가장 쉬운 방법은 SecurityContextHolder를 설정하는 것
- ThreadLocal을 사용해서 정보를 저장하기 때문에, 메소드에 직접 SecurityContext를 넘기지 않아도, 동일한 스레드라면 SecurityContext에 접근 가능
- 요청 처리 후 비워주는 것만 잊지 말 것 & Spring Security의 FilterChainProxy는 항상 SecurityContext를 비워줌

**ThreadLocal이 적합하지 않은 경우**
- `SecurityContextHolder.MODE_GLOBAL` = standalone application
- `SecurityContextHolder.MODE_INHERITABLETHREADLOCAL` = 인증 처리를 마친 스레드가 생성한 SecurityContext를 다른 스레드에서도 그래도 사용해야 하는 경우
- `SecurityContextHolder.MODE_THREADLOCAL` = DEFAULT

**SecurityContext**
- SecurityContextHolder로 접근
- Authentication 객체를 가짐

### 4. Authentication class
- `AuthenticationManager`의 input으로 사용되어, 인증에 사용할 `사용자의 credential`을 제공
  (이 경우, this.isAuthenticated() = false)
- 현재 인증된 사용자: Authentication은 SecurityContext에서 가져오는 것이 가능

**가져올 수 있는 정보**
- principal: 사용자 식별. 보통 UserDetails 인스턴스(사용자 이름/비밀번호)
- credentials: 비밀번호. 인증 후에 보통 비워둠.
- authorities: role, scope처럼 GrantedAuthority 클래스로 사용자에게 부여된 권한을 추상화
