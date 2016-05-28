package com.itpkg.core.services;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.io.Serializable;

/**
 * Created by flamen on 16-5-27.
 */
@Service
public class SettingService {
    @Transactional
    public void set(String key, Object val, boolean type) {

    }

    public <T extends Serializable> T get(String key, Class<T> clazz) {
        return null;
    }
}
