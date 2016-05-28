package com.itpkg.core.repositories;

import com.itpkg.core.models.Setting;
import org.springframework.data.repository.CrudRepository;

/**
 * Created by flamen on 16-5-28.
 */
public interface SettingRepository extends CrudRepository<Setting, Long> {
    Setting findByKey(String key);
}
