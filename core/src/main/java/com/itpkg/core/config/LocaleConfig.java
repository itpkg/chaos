package com.itpkg.core.config;

import com.itpkg.core.i18n.MultipleMessageSource;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.LocaleResolver;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;
import org.springframework.web.servlet.i18n.CookieLocaleResolver;
import org.springframework.web.servlet.i18n.LocaleChangeInterceptor;

import java.util.Locale;

/**
 * Created by flamen on 16-5-28.
 */
@Configuration
@EnableWebMvc
public class LocaleConfig extends WebMvcConfigurerAdapter {

    @Bean
    MultipleMessageSource messageSource() {
        MultipleMessageSource messageSource = new MultipleMessageSource();
        messageSource.setBasename("classpath*:/messages");
        messageSource.setDefaultEncoding("UTF-8");
        //messageSource().setUseCodeAsDefaultMessage(true);
        return messageSource;
    }


    @Bean
    LocaleResolver localeResolver() {
        CookieLocaleResolver resolver = new CookieLocaleResolver();
        resolver.setDefaultLocale(Locale.ENGLISH);
        resolver.setCookieName("locale");
        resolver.setCookieMaxAge(60 * 60 * 24 * 365 * 10);
        return resolver;
    }

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        LocaleChangeInterceptor interceptor = new LocaleChangeInterceptor();
        interceptor.setParamName("locale");
        registry.addInterceptor(interceptor);
    }
}
