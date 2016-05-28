package com.itpkg.core.services;

import com.itpkg.core.models.Setting;
import com.itpkg.core.repositories.SettingRepository;
import com.itpkg.core.utils.Encryptor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.Resource;
import java.io.IOException;
import java.io.Serializable;

/**
 * Created by flamen on 16-5-27.
 */
@Service
public class SettingService {
    @Transactional
    public void set(String key, Serializable val, boolean flag) throws IOException {
        Setting s = settingRepository.findByKey(key);
        if (s == null) {
            s = new Setting();
            s.setKey(key);
        }
        s.setFlag(flag);
        s.setVal(flag ? encryptor.encode(val) : encryptor.obj2str(val));
        settingRepository.save(s);
    }

    public <T extends Serializable> T get(String key, Class<T> clazz) throws IOException {
        Setting s = settingRepository.findByKey(key);
        if (s == null) {
            return null;
        }
        if (s.isFlag()) {
            return encryptor.decode(s.getVal(), clazz);
        }
        return encryptor.str2obj(s.getVal(), clazz);
    }

    @Resource
    Encryptor encryptor;
    @Resource
    SettingRepository settingRepository;
}
