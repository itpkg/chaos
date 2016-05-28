package com.itpkg.core.controllers;

import com.itpkg.core.models.User;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

/**
 * Created by flamen on 16-5-27.
 */
@RestController
public class UsersController {
    @RequestMapping(value = "/users", method = RequestMethod.GET)
    public User greeting(@RequestParam(value = "name", defaultValue = "World") String name) {
        return new User();
    }
}
