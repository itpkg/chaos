package com.itpkg.core.controllers;

import com.itpkg.core.forms.SignInFm;
import com.itpkg.core.models.User;
import com.itpkg.core.repositories.UserRepository;
import com.itpkg.core.utils.Encryptor;
import com.itpkg.core.web.Response;
import org.springframework.context.MessageSource;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.Locale;

/**
 * Created by flamen on 16-5-27.
 */
@RestController
@RequestMapping(value = "/personal")
public class PersonalController {
//    public class SignInFm {
//    }

    public class SignUpFm {

    }

    @RequestMapping(value = "/signIn", method = RequestMethod.POST)
    public Response signIn(@ModelAttribute SignInFm fm, Locale l) {
        Response rst = new Response();
        rst.setData(fm);
        User u = userRepository.findByEmail(fm.getEmail());
        if (u == null) {
            rst.setError(messageSource.getMessage("core.errors.user.email_not_exist", null, l));

        } else {
            if (u.getProviderType() == User.Type.EMAIL && encryptor.chk(fm.getPassword(), u.getPassword())) {
                rst.setOk(true);
                //todo
            } else {
                rst.setError(messageSource.getMessage("core.errors.user.email_password_not_match", null, l));
            }
        }
        return rst;
    }

    @RequestMapping(value = "/signUp", method = RequestMethod.POST)
    public void signUp() {
        //todo
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

}
