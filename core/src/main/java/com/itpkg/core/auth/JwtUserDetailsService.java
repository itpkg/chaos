package com.itpkg.core.auth;

import com.itpkg.core.models.User;
import com.itpkg.core.repositories.PermissionRepository;
import com.itpkg.core.repositories.UserRepository;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.Date;

/**
 * Created by flamen on 16-5-28.
 */
@Service
public class JwtUserDetailsService implements UserDetailsService {
    @Override
    public UserDetails loadUserByUsername(String uid) throws UsernameNotFoundException {
        User user = userRepository.findByUid(uid);
        if (user == null) {
            throw new UsernameNotFoundException(String.format("No user found with uid '%s'.", uid));
        }
        JwtUser details = new JwtUser(user);
        Date now = new Date();
        permissionRepository.findRoles(user.getId()).forEach(p -> {
            if (now.after(p.getBegin()) && now.before(p.getEnd())) {
                details.getAuthorities().add(new SimpleGrantedAuthority(p.getOperation()));
            }
        });

        return details;
    }

    @Resource
    UserRepository userRepository;
    @Resource
    PermissionRepository permissionRepository;
}
