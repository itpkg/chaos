package com.itpkg.core.services;

import com.itpkg.core.models.Log;
import com.itpkg.core.models.User;
import com.itpkg.core.repositories.LogRepository;
import com.itpkg.core.repositories.UserRepository;
import com.itpkg.core.utils.Encryptor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.Resource;
import java.util.UUID;

/**
 * Created by flamen on 16-5-27.
 */
@Service
public class UserService {
    @Transactional
    public User add(String name, String email, String password) {
        User u = new User();
        u.setName(name);
        u.setEmail(email);
        u.setPassword(encryptor.sum(password));

        String uid = UUID.randomUUID().toString();
        u.setUid(uid);
        u.setProviderId(uid);
        u.setProviderType(User.Type.EMAIL);
        userRepository.save(u);
        return u;
    }

    public void log(User user, String message) {
        Log l = new Log();
        l.setUser(user);
        l.setMessage(message);
        logRepository.save(l);
    }

    @Resource
    Encryptor encryptor;
    @Resource
    UserRepository userRepository;
    @Resource
    LogRepository logRepository;

}
