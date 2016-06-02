package com.itpkg.core.oauth2;

import com.google.api.services.oauth2.model.Userinfoplus;
import com.itpkg.core.models.User;
import com.itpkg.core.repositories.UserRepository;
import com.itpkg.core.services.UserService;
import com.itpkg.core.utils.JwtHelper;
import org.apache.http.client.utils.URLEncodedUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.MessageSource;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.HashMap;
import java.util.Locale;
import java.util.Map;

/**
 * Created by flamen on 16-6-2.
 */
@RestController
@RequestMapping(value = "/oauth2")
public class Oauth2Controller {
    @RequestMapping(value = "/callback", method = RequestMethod.POST)
    public String callback(@RequestParam(value = "href") String href, Locale l) throws URISyntaxException, IOException {
        Map<String, String> params = new HashMap<>();

        URLEncodedUtils.parse(new URI(href), "UTF-8").forEach(vp -> {
            params.put(vp.getName(), vp.getValue());
        });

        logger.debug(href);
        logger.debug("PARAMS: {}", params);

        String state = params.get("state");

        if (state != null && state.startsWith(Oauth2GoogleHelper.ID)) {
            //google oauth2
            Userinfoplus gu = googleHelper.getUser(params.get("code"));
            User u = userRepository.findByEmail(gu.getEmail());
            if (u == null) {
                u = userService.add(User.Type.GOOGLE, (String) gu.get("sub"), gu.getName(), gu.getEmail());
                userService.log(u, messageSource.getMessage("core.logs.user.signUp", null, l));
            }
            userService.signIn(u);
            userService.log(u, messageSource.getMessage("core.logs.user.signIn", null, l));
            return jwtHelper.generate(u);
        }

        throw new IllegalArgumentException();
    }

    @Resource
    JwtHelper jwtHelper;
    @Resource
    UserRepository userRepository;
    @Resource
    Oauth2GoogleHelper googleHelper;
    @Resource
    UserService userService;
    @Resource
    MessageSource messageSource;

    private final static Logger logger = LoggerFactory.getLogger(Oauth2Controller.class);

}
