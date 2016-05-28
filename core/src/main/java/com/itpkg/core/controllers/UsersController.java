package com.itpkg.core.controllers;

import com.itpkg.core.models.User;
import com.itpkg.core.utils.Encryptor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.io.IOException;

/**
 * Created by flamen on 16-5-27.
 */
@RestController
public class UsersController {
    @RequestMapping(value = "/users", method = RequestMethod.GET)
    public User greeting(@RequestParam(value = "name", defaultValue = "World") String name) throws IOException{
        User u = new User();
        u.setName(secret);
        u.setName(encryptor.encode("hello"));
        u.setPassword(encryptor.sum("hello"));
        return u;
    }
    @Resource
    private Encryptor encryptor;

    @Value("${app.secret}")
    private String secret;

    public void setSecret(String secret) {
        this.secret = secret;
    }

    public void setEncryptor(Encryptor encryptor) {
        this.encryptor = encryptor;
    }
}
