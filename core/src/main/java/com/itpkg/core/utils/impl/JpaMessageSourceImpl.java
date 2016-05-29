package com.itpkg.core.utils.impl;

import com.itpkg.core.repositories.LocaleRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.support.AbstractMessageSource;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.text.MessageFormat;
import java.util.Locale;

/**
 * Created by flamen on 16-5-28.
 */
//@Component
public class JpaMessageSourceImpl extends AbstractMessageSource {
    @Override
    protected MessageFormat resolveCode(String code, Locale locale) {
        com.itpkg.core.models.Locale l = localeRepository.findByCodeAndLang(code, locale.toString());

        if (l == null) {
            return createMessageFormat(code, locale);
        }
        return createMessageFormat(l.getMessage(), locale);
    }

    @Resource
    LocaleRepository localeRepository;

    private final static Logger logger = LoggerFactory.getLogger(JpaMessageSourceImpl.class);
}
