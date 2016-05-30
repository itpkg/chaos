package com.itpkg.core.controllers;

import com.itpkg.core.models.*;
import com.itpkg.core.repositories.LocaleRepository;
import com.itpkg.core.web.Link;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.security.Principal;
import java.util.*;
import java.util.Locale;

/**
 * Created by flamen on 16-5-28.
 */
@RestController
public class HomeController {

    @RequestMapping(value = "/info", method = RequestMethod.GET)
    public Map<String, Object> info(Locale locale) {
        Map<String, Object> map = new HashMap<>();
        for (String k : new String[]{"title", "subTitle", "keywords", "description", "copyright"}) {
            com.itpkg.core.models.Locale l = localeRepository.findByCodeAndLang("site." + k, locale.toString());
            map.put(k, l == null ? k : l.getMessage());
        }
        map.put("languages", new String[]{Locale.ENGLISH.toString(), Locale.SIMPLIFIED_CHINESE.toString()});
        List<Link> links = new ArrayList<>();
        //todo
        links.add(new Link("index", "Home"));
        links.add(new Link("reading.notes", "Notes"));
        map.put("links", links);

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
