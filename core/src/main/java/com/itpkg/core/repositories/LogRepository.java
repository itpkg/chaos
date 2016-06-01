package com.itpkg.core.repositories;

import com.itpkg.core.models.Locale;
import com.itpkg.core.models.Log;
import org.springframework.data.repository.CrudRepository;

/**
 * Created by flamen on 16-6-1.
 */
public interface LogRepository extends CrudRepository<Log, Long> {
}
