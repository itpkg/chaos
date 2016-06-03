package com.itpkg.core.controllers;

import com.itpkg.core.auth.JwtHandler;
import com.itpkg.core.forms.EmailFm;
import com.itpkg.core.forms.ResetPasswordFm;
import com.itpkg.core.forms.SignInFm;
import com.itpkg.core.forms.SignUpFm;
import com.itpkg.core.jobs.EmailSender;
import com.itpkg.core.jobs.Mail;
import com.itpkg.core.models.User;
import com.itpkg.core.repositories.UserRepository;
import com.itpkg.core.services.UserService;
import com.itpkg.core.utils.Encryptor;
import com.itpkg.core.utils.JwtHelper;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.MessageSource;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import javax.validation.Valid;
import java.time.temporal.ChronoUnit;
import java.util.HashMap;
import java.util.Locale;
import java.util.Map;

/**
 * Created by flamen on 16-5-27.
 */
@RestController
@RequestMapping(value = "/personal")
public class PersonalController {

    @RequestMapping(value = "/signIn", method = RequestMethod.POST)
    public String signIn(@Valid @ModelAttribute SignInFm fm, Locale l) {
        User u = userRepository.findByEmail(fm.getEmail());
        if (u == null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_not_exist", null, l));
        }
        if (u.getProviderType() != User.Type.EMAIL || !encryptor.chk(fm.getPassword(), u.getPassword())) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_password_not_match", null, l));
        }
        if (u.getConfirmedAt() == null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.not_confirmed", null, l));
        }
        if (u.getLockedAt() != null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.locked", null, l));
        }


        return jwtHelper.generate(u);

    }

    @RequestMapping(value = "/signUp", method = RequestMethod.POST)
    public String signUp(@Valid @ModelAttribute SignUpFm fm, Locale l) {
        if (!fm.getPassword().equals(fm.getPasswordConfirm())) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.passwords_not_match", null, l));
        }

        if (userRepository.findByEmail(fm.getEmail()) != null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_already_exist", null, l));
        }
        User u = userService.add(fm.getName(), fm.getEmail(), fm.getPassword());
        userService.log(u, messageSource.getMessage("core.logs.user.signUp", null, l));
        sendEmail(u, "confirm", l);
        return messageSource.getMessage("core.messages.user.confirm", null, l);
    }

    @RequestMapping(value = "/signOut", method = RequestMethod.POST)
    public void signOut() {
        //todo
    }

    @RequestMapping(value = "/confirm", method = RequestMethod.POST)
    public String confirm(@Valid @ModelAttribute EmailFm fm, Locale l) {
        User u = userRepository.findByEmail(fm.getEmail());
        if (u == null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_not_exist", null, l));
        }
        if (u.getConfirmedAt() != null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.already_confirm", null, l));
        }
        sendEmail(u, "confirm", l);
        return messageSource.getMessage("core.messages.user.confirm", null, l);
    }

    @RequestMapping(value = "/unlock", method = RequestMethod.POST)
    public String unlock(@Valid @ModelAttribute EmailFm fm, Locale l) {
        User u = userRepository.findByEmail(fm.getEmail());
        if (u == null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_not_exist", null, l));
        }
        if (u.getLockedAt() == null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.not_locked", null, l));
        }
        sendEmail(u, "unlock", l);
        return messageSource.getMessage("core.messages.user.unlock", null, l);
    }

    @RequestMapping(value = "/forgotPassword", method = RequestMethod.POST)
    public String forgotPassword(@Valid @ModelAttribute EmailFm fm, Locale l) {
        User u = userRepository.findByEmail(fm.getEmail());
        if (u == null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_not_exist", null, l));
        }
        sendEmail(u, "resetPassword", l);
        return messageSource.getMessage("core.messages.user.resetPassword", null, l);
    }

    @RequestMapping(value = "/resetPassword", method = RequestMethod.POST)
    public void resetPassword(@Valid @ModelAttribute ResetPasswordFm fm, Locale l) {
        //todo
    }

    private String tokenUrl(User u, String act, Locale l) {
        Map<String, Object> map = new HashMap<>();
        map.put("uid", u.getUid());
        map.put("act", act);
        String tkn = jwtHandler.generate(u.getName(), map, 30, ChronoUnit.MINUTES);
        return String.format("%s/personal/%s?token=%s&locale=%s", home, act, tkn, l.toString());
    }

    private void sendEmail(User u, String act, Locale l) {

        Mail m = new Mail();
        m.subject = messageSource.getMessage("core.emails." + act + ".title", null, l);
        m.to = u.getEmail();

        switch (act) {
            case "confirm":
            case "unlock":
            case "resetPassword":
                m.body = messageSource.getMessage("core.emails." + act + ".body", new Object[]{u.getName(), tokenUrl(u, act, l)}, l);
                break;
            case "passwordChange":
                m.body = messageSource.getMessage("core.emails." + act + ".body", new Object[]{u.getName()}, l);
                break;
            default:
                throw new IllegalArgumentException(act);
        }


        emailSender.send(m);

    }

    @Resource
    JwtHandler jwtHandler;
    @Value("${app.home}")
    String home;

    @Resource
    UserRepository userRepository;
    @Resource
    MessageSource messageSource;
    @Resource
    Encryptor encryptor;
    @Resource
    JwtHelper jwtHelper;

    @Resource
    UserService userService;
    @Resource
    EmailSender emailSender;

}
