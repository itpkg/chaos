package com.itpkg.core.controllers;

import com.itpkg.core.repositories.LocaleRepository;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.security.Principal;
import java.util.Date;
import java.util.HashMap;
import java.util.Locale;
import java.util.Map;

/**
 * Created by flamen on 16-5-28.
 */
@RestController
public class HomeController {

    @RequestMapping(value = "/locales/{lang}.json", method = RequestMethod.GET)
    public Map<String, String> locale(@PathVariable String lang) {
        Map<String, String> map = new HashMap<>();
        for (com.itpkg.core.models.Locale l : localeRepository.findByLang(lang)) {
            map.put(l.getCode(), l.getMessage());
        }

        return map;
    }

    @RequestMapping(value = "/info", method = RequestMethod.GET)
    public Map<String, Object> info(Locale locale) {
        Map<String, Object> map = new HashMap<>();

        map.put("title", "todo");
        map.put("locale", locale);
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

    @Resource
    LocaleRepository localeRepository;

}
