package com.itpkg.core.utils;

import com.itpkg.core.auth.JwtHandler;
import com.itpkg.core.models.User;
import com.itpkg.core.repositories.PermissionRepository;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.time.temporal.ChronoUnit;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static com.itpkg.core.auth.JwtHandler.ROLES;
import static com.itpkg.core.auth.JwtHandler.UID;

/**
 * Created by flamen on 16-6-2.
 */
@Component
public class JwtHelper {

    public String generate(User u) {
        Map<String, Object> data = new HashMap<>();
        data.put(UID, u.getUid());
        List<String> roles = new ArrayList<>();
        permissionRepository.findRoles(u.getId()).forEach(p -> {
            roles.add(p.getOperation());
        });
        data.put(ROLES, roles);
        return jwtHandler.generate(u.getName(), data, 1, ChronoUnit.WEEKS);
    }

    @Resource
    JwtHandler jwtHandler;
    @Resource
    PermissionRepository permissionRepository;
}
