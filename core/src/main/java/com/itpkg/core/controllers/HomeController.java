package com.itpkg.core.controllers;

import com.itpkg.core.repositories.LocaleRepository;
import com.itpkg.core.services.SettingService;
import com.itpkg.core.web.Link;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.io.IOException;
import java.security.Principal;
import java.util.*;

/**
 * Created by flamen on 16-5-28.
 */
@RestController
public class HomeController {

    @RequestMapping(value = "/info", method = RequestMethod.GET)
    public Map<String, Object> info(Locale locale) throws IOException {
        Map<String, Object> map = new HashMap<>();
        map.put("locale", locale.toString());
        for (String k : new String[]{"title", "subTitle", "keywords", "description", "copyright"}) {
            com.itpkg.core.models.Locale l = localeRepository.findByCodeAndLang("site." + k, locale.toString());
            map.put(k, l == null ? locale.toString() + "." + k : l.getMessage());
        }

        List<Link> links = settingService.getList("top.links", Link.class);
        if (links == null) {
            links = new ArrayList<>();
            links.add(new Link("index", "core.pages.index"));
        }
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
    @Resource
    SettingService settingService;


}
