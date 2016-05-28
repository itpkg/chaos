package com.itpkg.core.auth;

import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;
import org.springframework.security.web.authentication.WebAuthenticationDetailsSource;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;
import java.io.IOException;

/**
 * Created by flamen on 16-5-28.
 */
@Component
public class JwtAuthenticationFilter extends UsernamePasswordAuthenticationFilter {
    public final static String AUTHORIZATION = "Authorization";
    public final static String BEARER = "Bearer ";
    public final static String UID = "uid";

    @Override
    public void doFilter(ServletRequest req, ServletResponse res, FilterChain chain) throws IOException, ServletException {
        HttpServletRequest request = (HttpServletRequest) req;
        String header = request.getHeader(AUTHORIZATION);
        if (header == null || !header.startsWith(BEARER)) {
            throw new IOException("No JWT token found in request headers");
        }
        String uid = jwtHandler.parse(header.substring(BEARER.length())).get(UID);
        if (uid != null && SecurityContextHolder.getContext().getAuthentication() == null) {
            UserDetails user = userDetailsService.loadUserByUsername(uid);
            UsernamePasswordAuthenticationToken token = new UsernamePasswordAuthenticationToken(user, null, user.getAuthorities());
            token.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));
            SecurityContextHolder.getContext().setAuthentication(token);
        }
        chain.doFilter(req, res);

    }

    @PostConstruct
    void init() {
        setAuthenticationManager(authenticationManager);
    }

    @Resource
    AuthenticationManager authenticationManager;
    @Resource
    JwtHandler jwtHandler;
    @Resource
    UserDetailsService userDetailsService;
}
