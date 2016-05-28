package com.itpkg.core.controllers;

import com.itpkg.core.models.User;
import com.itpkg.core.utils.Jwt;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.io.IOException;
import java.time.temporal.ChronoUnit;
import java.util.HashMap;
import java.util.Map;

/**
 * Created by flamen on 16-5-27.
 */
@RestController
public class UsersController {
    @RequestMapping(value = "/users", method = RequestMethod.GET)
    public User greeting(@RequestParam(value = "name", defaultValue = "World") String name) throws IOException {

        Map<String, String> map = new HashMap<>();
        map.put("msg", "你好!!!");
        String token = jwt.generate("hello", map, 7, ChronoUnit.DAYS);

        User u = new User();
        u.setProviderId(token);
        u.setName(jwt.parse(token).get("msg"));
        return u;
    }

    @Resource
    private Jwt jwt;


    public void setJwt(Jwt jwt) {
        this.jwt = jwt;
    }

}
