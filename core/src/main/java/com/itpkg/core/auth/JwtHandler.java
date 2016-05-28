package com.itpkg.core.auth;

import java.time.temporal.TemporalUnit;
import java.util.Map;

/**
 * Created by flamen on 16-5-28.
 */
public interface JwtHandler {
    Map<String, String> parse(String token);

    String generate(String subject, Map<String, String> data, long exp, TemporalUnit unit);
}
