package com.itpkg.core.auth;

import com.itpkg.core.models.User;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import java.util.ArrayList;
import java.util.Collection;

/**
 * Created by flamen on 16-5-28.
 */
public class JwtUser implements UserDetails {
    private final User user;
    private final Collection<SimpleGrantedAuthority> authorities;

    public JwtUser(User user) {
        this.user = user;
        this.authorities = new ArrayList<>();
    }

    @Override
    public Collection<SimpleGrantedAuthority> getAuthorities() {
        return authorities;
    }

    @Override
    public String getPassword() {
        return user.getPassword();
    }

    @Override
    public String getUsername() {
        return user.getName();
    }

    @Override
    public boolean isAccountNonExpired() {
        return true;
    }

    @Override
    public boolean isAccountNonLocked() {
        return true;
    }

    @Override
    public boolean isCredentialsNonExpired() {
        return true;
    }

    @Override
    public boolean isEnabled() {
        return user.getConfirmedAt() != null && user.getLockedAt() == null;
    }

    public String getUid() {
        return user.getUid();
    }

    public String getEmail() {
        return user.getEmail();
    }
}
