package com.itpkg.core.controllers;

import com.itpkg.core.auth.JwtHandler;
import com.itpkg.core.forms.SignInFm;
import com.itpkg.core.forms.SignUpFm;
import com.itpkg.core.models.Permission;
import com.itpkg.core.models.User;
import com.itpkg.core.repositories.PermissionRepository;
import com.itpkg.core.repositories.UserRepository;
import com.itpkg.core.services.UserService;
import com.itpkg.core.utils.Encryptor;
import com.itpkg.core.web.Response;
import org.springframework.context.MessageSource;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.time.temporal.ChronoUnit;
import java.util.*;

import static com.itpkg.core.auth.JwtHandler.ROLES;
import static com.itpkg.core.auth.JwtHandler.UID;

/**
 * Created by flamen on 16-5-27.
 */
@RestController
@RequestMapping(value = "/personal")
public class PersonalController {

    @RequestMapping(value = "/signIn", method = RequestMethod.POST)
    public String signIn(@ModelAttribute SignInFm fm, Locale l) {
        User u = userRepository.findByEmail(fm.getEmail());
        if (u == null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_not_exist", null, l));

        }
        if (u.getProviderType() != User.Type.EMAIL || !encryptor.chk(fm.getPassword(), u.getPassword())) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_password_not_match", null, l));
        }
        Map<String, Object> tkn = new HashMap<>();
        tkn.put(UID, u.getUid());
        List<String> roles = new ArrayList<>();
        for(Permission p : permissionRepository.findRoles(u.getId())){
            roles.add(p.getOperation());
        }
        tkn.put(ROLES, roles);

        return jwtHandler.generate(u.getName(), tkn, 1, ChronoUnit.WEEKS);

    }

    @RequestMapping(value = "/signUp", method = RequestMethod.POST)
    public Map<String, Object> signUp(@ModelAttribute SignUpFm fm, Locale l) {

        if (userRepository.findByEmail(fm.getEmail()) != null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_already_exist", null, l));
        }
        User u = userService.add(fm.getName(), fm.getEmail(), fm.getPassword());
        userService.log(u, messageSource.getMessage("core.logs.user.signUp", null, l));
        return u.toModel();
    }

    @RequestMapping(value = "/signOut", method = RequestMethod.POST)
    public void signOut() {
        //todo
    }

    @RequestMapping(value = "/confirm", method = RequestMethod.POST)
    public void confirm() {
        //todo
    }

    @RequestMapping(value = "/unlock", method = RequestMethod.POST)
    public void unlock() {
        //todo
    }

    @RequestMapping(value = "/forgotPassword", method = RequestMethod.POST)
    public void forgotPassword() {
        //todo
    }

    @RequestMapping(value = "/resetPassword", method = RequestMethod.POST)
    public void resetPassword() {
        //todo
    }

    @Resource
    UserRepository userRepository;
    @Resource
    MessageSource messageSource;
    @Resource
    Encryptor encryptor;
    @Resource
    JwtHandler jwtHandler;
    @Resource
    PermissionRepository permissionRepository;
    @Resource
    UserService userService;

}
