package com.itpkg.core.repositories;

import com.itpkg.core.models.Permission;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;

import java.util.List;

/**
 * Created by flamen on 16-5-28.
 */
public interface PermissionRepository extends CrudRepository<Permission, Long> {
    @Query("SELECT o FROM Permission o WHERE o.user.id = ?1 AND o.resourceType IS NULL AND o.resourceId IS NULL")
    List<Permission> findRoles(long userId);
}
