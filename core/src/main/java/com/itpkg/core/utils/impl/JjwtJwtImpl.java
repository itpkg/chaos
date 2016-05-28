package com.itpkg.core.utils.impl;

import com.itpkg.core.utils.Jwt;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.time.LocalDateTime;
import java.time.ZoneId;
import java.time.temporal.TemporalUnit;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

/**
 * Created by flamen on 16-5-28.
 */
@Component
public class JjwtJwtImpl implements Jwt {

    @Override
    public Map<String, String> parse(String token) {
        Claims body = Jwts.parser().setSigningKey(secret).parseClaimsJws(token).getBody();
        Map<String, String> map = new HashMap<>();
        for (String k : body.keySet()) {
            if (k.equals("nbf") || k.equals("exp")) {
                continue;
            }
            map.put(k, body.get(k, String.class));
        }
        return map;
    }

    @Override
    public String generate(String subject, Map<String, String> data, long exp, TemporalUnit unit) {
        Claims claims = Jwts.claims().setSubject(subject);
        claims.putAll(data);
        claims.setNotBefore(new Date());
        claims.setExpiration(Date.from(LocalDateTime.now().plus(exp, unit).atZone(ZoneId.systemDefault()).toInstant()));

        return Jwts.builder()
                .setClaims(claims)
                .signWith(SignatureAlgorithm.HS512, secret)
                .compact();
    }

    @Value("${jwt.secret}")
    private String secret;

    public void setSecret(String secret) {
        this.secret = secret;
    }

}
