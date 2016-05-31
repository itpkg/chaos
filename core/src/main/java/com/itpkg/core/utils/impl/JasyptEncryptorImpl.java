package com.itpkg.core.utils.impl;

import com.itpkg.core.utils.Encryptor;
import org.jasypt.util.password.PasswordEncryptor;
import org.jasypt.util.password.StrongPasswordEncryptor;
import org.jasypt.util.text.StrongTextEncryptor;
import org.jasypt.util.text.TextEncryptor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;

/**
 * Created by flamen on 16-5-27.
 */
@Component
public class JasyptEncryptorImpl implements Encryptor {

    @Override
    public String sum(String plain) {
        return password.encryptPassword(plain);
    }

    @Override
    public boolean chk(String plain, String code) {
        return password.checkPassword(plain, code);
    }

    @Override
    public String encode(String plain) {
        return text.encrypt(plain);
    }

    @Override
    public String decode(String code) {
        return text.decrypt(code);
    }


    @PostConstruct
    void init() {
        password = new StrongPasswordEncryptor();
        StrongTextEncryptor ste = new StrongTextEncryptor();
        ste.setPassword(secret);
        this.text = ste;
    }


    private PasswordEncryptor password;
    private TextEncryptor text;
    @Value("${app.secret}")
    String secret;

}
