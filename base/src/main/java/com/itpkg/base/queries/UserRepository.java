package com.itpkg.base.queries;

import com.itpkg.base.models.User;
import org.springframework.data.repository.CrudRepository;

import java.util.List;

/**
 * Created by flamen on 16-5-27.
 */
public interface UserRepository extends CrudRepository<User, Long> {
    User findByEmail(String email);

    List<User> findByProviderType(String providerType);
}
