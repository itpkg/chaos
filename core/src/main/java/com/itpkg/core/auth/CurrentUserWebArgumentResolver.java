package com.itpkg.core.auth;

import com.itpkg.core.models.User;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.MethodParameter;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.support.WebDataBinderFactory;
import org.springframework.web.context.request.NativeWebRequest;
import org.springframework.web.method.support.HandlerMethodArgumentResolver;
import org.springframework.web.method.support.ModelAndViewContainer;

/**
 * Created by flamen on 16-5-28.
 */
@Component
public class CurrentUserWebArgumentResolver implements HandlerMethodArgumentResolver {

    @Override
    public boolean supportsParameter(MethodParameter parameter) {
        return parameter.getParameterAnnotation(CurrentUser.class) != null;
    }

    @Override
    public Object resolveArgument(MethodParameter parameter, ModelAndViewContainer mavContainer, NativeWebRequest webRequest, WebDataBinderFactory binderFactory) throws Exception {
        logger.debug("GET CURRENT USER");
        //todo
        User user = new User();
        user.setName("test user");
        return user;
        //WebArgumentResolver.UNRESOLVED
    }

    private final Logger logger = LoggerFactory.getLogger(CurrentUserWebArgumentResolver.class);
}
