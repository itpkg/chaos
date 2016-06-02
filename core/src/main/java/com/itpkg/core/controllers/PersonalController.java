package com.itpkg.core.controllers;

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
import org.hibernate.validator.constraints.Email;
import org.springframework.context.MessageSource;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import javax.validation.Valid;
import java.util.Locale;

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
        if(u.getConfirmedAt()==null){
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.not_confirmed", null, l));
        }
        if(u.getLockedAt()!=null){
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.locked", null, l));
        }


        return jwtHelper.generate(u);

    }

    @RequestMapping(value = "/signUp", method = RequestMethod.POST)
    public String signUp(@Valid @ModelAttribute SignUpFm fm, Locale l) {
        if(!fm.getPassword().equals(fm.getPasswordConfirm())){
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.passwords_not_match", null, l));
        }

        if (userRepository.findByEmail(fm.getEmail()) != null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_already_exist", null, l));
        }
        User u = userService.add(fm.getName(), fm.getEmail(), fm.getPassword());
        userService.log(u, messageSource.getMessage("core.logs.user.signUp", null, l));
        sendConfirmEmail(u, l);
        return messageSource.getMessage("core.messages.user.confirm", null, l);
    }

    @RequestMapping(value = "/signOut", method = RequestMethod.POST)
    public void signOut() {
        //todo
    }

    @RequestMapping(value = "/confirm", method = RequestMethod.POST)
    public String confirm(@Valid @ModelAttribute SignUpFm fm, Locale l) {
        User u = userRepository.findByEmail(fm.getEmail());
        if (u == null) {
            throw new IllegalArgumentException(messageSource.getMessage("core.errors.user.email_not_exist", null, l));
        }
        sendConfirmEmail(u, l);
        return messageSource.getMessage("core.messages.user.confirm", null, l);
    }

    @RequestMapping(value = "/unlock", method = RequestMethod.POST)
    public void unlock(@Valid @ModelAttribute EmailFm fm, Locale l) {
        //todo
    }

    @RequestMapping(value = "/forgotPassword", method = RequestMethod.POST)
    public void forgotPassword(@Valid @ModelAttribute EmailFm fm, Locale l) {
        //todo
    }

    @RequestMapping(value = "/resetPassword", method = RequestMethod.POST)
    public void resetPassword(@Valid @ModelAttribute ResetPasswordFm fm, Locale l) {
        //todo
    }

    private void sendConfirmEmail(User u, Locale l) {
        //todo
        Mail m = new Mail();
        m.subject = "SSS";
        m.to = "to@test.com";
        m.body = "<h1>aaa</h1>";

        emailSender.send(m);

    }

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
