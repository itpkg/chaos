package com.itpkg.core.utils.impl;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.itpkg.core.utils.JsonHelper;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.io.IOException;
import java.io.Serializable;
import java.util.List;

/**
 * Created by flamen on 16-5-31.
 */
@Component
public class JacksonJsonHelperImpl implements JsonHelper {
    @Override
    public String obj2str(Serializable obj) throws IOException {
        return mapper.writeValueAsString(obj);
    }

    @Override
    public <T extends Serializable> T str2obj(String code, Class<T> clazz) throws IOException {
        return mapper.readValue(code, clazz);
    }

    @Override
    public <T extends Serializable> List<T> str2lst(String s, Class<T> clazz) throws IOException {
        return mapper.readValue(s, new TypeReference<List<T>>() {
        });
    }

    @PostConstruct
    void init() {
        mapper = new ObjectMapper();
    }

    private ObjectMapper mapper;
}
