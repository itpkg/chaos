package com.itpkg.core.services;

import com.itpkg.core.models.Setting;
import com.itpkg.core.repositories.SettingRepository;
import com.itpkg.core.utils.Encryptor;
import com.itpkg.core.utils.JsonHelper;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.Resource;
import java.io.IOException;
import java.io.Serializable;
import java.util.List;

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
        String jsn = jsonHelper.obj2str(val);
        s.setVal(flag ? encryptor.encode(jsn) : jsn);
        settingRepository.save(s);
    }

    public <T extends Serializable> List<T> getList(String key, Class<T> clazz) throws IOException {
        Setting s = settingRepository.findByKey(key);
        if (s == null) {
            return null;
        }
        return jsonHelper.str2lst(s.isFlag() ? encryptor.decode(s.getVal()) : s.getVal(), clazz);
    }

    public <T extends Serializable> T get(String key, Class<T> clazz) throws IOException {
        Setting s = settingRepository.findByKey(key);
        if (s == null) {
            return null;
        }

        return jsonHelper.str2obj(s.isFlag() ? encryptor.decode(s.getVal()) : s.getVal(), clazz);
    }

    @Resource
    Encryptor encryptor;
    @Resource
    JsonHelper jsonHelper;
    @Resource
    SettingRepository settingRepository;
}
