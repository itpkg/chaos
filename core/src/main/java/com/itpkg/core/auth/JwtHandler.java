package com.itpkg.core.auth;

import java.time.temporal.TemporalUnit;
import java.util.Map;

/**
 * Created by flamen on 16-5-28.
 */
public interface JwtHandler {
    String UID = "uid";
    String ROLES = "roles";
    String AUTHORIZATION = "Authorization";
    String BEARER = "Bearer ";


    Map<String, Object> parse(String token);

    String generate(String subject, Map<String, Object> data, long exp, TemporalUnit unit);
}
