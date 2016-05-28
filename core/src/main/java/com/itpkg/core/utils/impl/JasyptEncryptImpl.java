package com.itpkg.core.utils.impl;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.itpkg.core.utils.Encrypt;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;
import java.io.IOException;
import java.io.Serializable;

/**
 * Created by flamen on 16-5-27.
 */
@Component
public class JasyptEncryptImpl implements Encrypt {

    @Override
    public String sum(Object obj) {
        //todo
        return null;
    }

    @Override
    public boolean chk(Object obj, String code) {
        //todo
        return false;
    }

    @Override
    public String encode(Object obj) {
        //todo
        return null;
    }

    @Override
    public <T extends Serializable> T decode(String code, Class<T> clazz) {
        //todo
        return null;
    }


    @PostConstruct
    void init() {
        mapper = new ObjectMapper();
    }

    @PreDestroy
    void destroy() {

    }

    private String obj2str(Object obj) throws JsonProcessingException {
        return mapper.writeValueAsString(obj);
    }

    private <T extends Serializable> T str2obj(String code, Class<T> clazz) throws IOException {
        return mapper.readValue(code, clazz);
    }

    private ObjectMapper mapper;

    //StrongPasswordEncryptor passwordEncryptor = new StrongPasswordEncryptor();
}
