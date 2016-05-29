package com.itpkg.core.repositories;

import com.itpkg.core.models.Locale;
import org.springframework.data.repository.CrudRepository;

import java.util.List;

/**
 * Created by flamen on 16-5-28.
 */
public interface LocaleRepository extends CrudRepository<Locale, Long> {
    Locale findByCodeAndLang(String code, String lang);

    List<Locale> findByLang(String lang);
}
