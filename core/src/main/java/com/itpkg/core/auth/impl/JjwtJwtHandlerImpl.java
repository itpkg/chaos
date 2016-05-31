package com.itpkg.core.auth.impl;

import com.itpkg.core.auth.JwtHandler;
import com.itpkg.core.services.SettingService;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.JwtException;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import org.apache.commons.lang3.RandomStringUtils;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.io.IOException;
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
public class JjwtJwtHandlerImpl implements JwtHandler {


    @Override
    public Map<String, Object> parse(String token) {
        Claims body = Jwts.parser().setSigningKey(secret).parseClaimsJws(token).getBody();
        Date now = new Date();
        if (now.before(body.getNotBefore()) || now.after(body.getExpiration())) {
            throw new JwtException("token invalid.");
        }

//        Map<String, String> map = new HashMap<>();
//        for (String k : body.keySet()) {
//            if (k.equals("nbf") || k.equals("exp")) {
//                continue;
//            }
//            map.put(k, body.get(k, String.class));
//        }
        return body;
    }

    @Override
    public String generate(String subject, Map<String, Object> data, long exp, TemporalUnit unit) {
        Claims claims = Jwts.claims().setSubject(subject);
        claims.putAll(data);
        claims.setNotBefore(new Date());
        claims.setExpiration(Date.from(LocalDateTime.now().plus(exp, unit).atZone(ZoneId.systemDefault()).toInstant()));

        return Jwts.builder()
                .setClaims(claims)
                .signWith(SignatureAlgorithm.HS512, secret)
                .compact();
    }


    @PostConstruct
    void init() throws IOException {

        secret = settingService.get(KEY, String.class);
        if (secret == null) {
            secret = RandomStringUtils.random(32);
            settingService.set(KEY, secret, true);
        }
    }

    @Resource
    SettingService settingService;
    private String secret;

    public final static String KEY = "jwt.secret";


}
