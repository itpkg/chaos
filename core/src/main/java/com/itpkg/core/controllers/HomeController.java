package com.itpkg.core.controllers;

import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.security.Principal;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

/**
 * Created by flamen on 16-5-28.
 */
@RestController
public class HomeController {

    @RequestMapping(value = "/info", method = RequestMethod.GET)
    @PreAuthorize("permitAll")
    public Map<String, Object> info() {
        Map<String, Object> map = new HashMap<>();

        map.put("title", "todo");
        map.put("created", new Date());
        return map;
    }


    @RequestMapping(value = "/status", method = RequestMethod.GET)
    @PreAuthorize("hasRole('ADMIN')")
    public Map<String, Object> status(Principal principal) {
        Map<String, Object> map = new HashMap<>();

        map.put("principal", principal);
        map.put("created", new Date());
        return map;
    }

}
