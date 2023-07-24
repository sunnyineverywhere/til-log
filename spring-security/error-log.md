## A granted authority textual representation is required
: Authentication 내에 권한을 알 수 없어서 나타나는 에러

코드 발췌
```java
/* jwtFilter */
if(StringUtils.hasText(jwt) && jwtUtil.validateToken(jwt)) {
            Authentication authentication = jwtUtil.getAuthentication(jwt);
            SecurityContext context = SecurityContextHolder.createEmptyContext();
            context.setAuthentication(authentication);
            SecurityContextHolder.setContext(context);
        }
```

```java
/* jwtUtil */
Collection<? extends GrantedAuthority> authorities =
                Arrays.stream(claims.get(AUTHORITIES_KEY).toString().split(","))
                        .map(SimpleGrantedAuthority::new).toList();
```

## There is no PasswordEncoder mapped for the id "null"
: 스프링 시큐리티에서 패스워드 저장 시, 암호화를 하지 않아 발생하는 에러
(아 진짜 까다롭게 구네,,, 해준다 해줘)
