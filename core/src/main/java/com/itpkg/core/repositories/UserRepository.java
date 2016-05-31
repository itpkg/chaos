package com.itpkg.core.repositories;

import com.itpkg.core.models.User;
import org.springframework.data.repository.CrudRepository;

/**
 * Created by flamen on 16-5-28.
 */
public interface UserRepository extends CrudRepository<User, Long> {
    User findByUid(String uid);

    User findByEmail(String email);
}
