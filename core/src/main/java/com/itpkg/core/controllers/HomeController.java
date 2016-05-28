package com.itpkg.core.controllers;

import com.itpkg.core.models.User;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;

/**
 * Created by flamen on 16-5-28.
 */
@RestController
public class HomeController {
    @RequestMapping(value = "/site", method = RequestMethod.GET)
    public User greeting(@RequestParam(value = "name", defaultValue = "World") String name) throws IOException {

        User u = new User();
        return u;
    }

}
