package com.itpkg.core.auth;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.method.support.HandlerMethodArgumentResolver;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;

import javax.annotation.Resource;
import java.util.List;

/**
 * Created by flamen on 16-5-28.
 */
@Configuration
@EnableWebMvc
public class Config extends WebMvcConfigurerAdapter {
    @Override
    public void addArgumentResolvers(List<HandlerMethodArgumentResolver> argumentResolvers) {
        argumentResolvers.add(currentUserResolver);
    }

    @Resource
    CurrentUserWebArgumentResolver currentUserResolver;
}
