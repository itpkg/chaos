package com.itpkg.core.utils.impl;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.itpkg.core.utils.Encryptor;
import org.jasypt.util.password.PasswordEncryptor;
import org.jasypt.util.password.StrongPasswordEncryptor;
import org.jasypt.util.text.StrongTextEncryptor;
import org.jasypt.util.text.TextEncryptor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.io.IOException;
import java.io.Serializable;

/**
 * Created by flamen on 16-5-27.
 */
@Component
public class JasyptEncryptorImpl implements Encryptor {

    @Override
    public String sum(Serializable obj) throws IOException {
        return password.encryptPassword(obj2str(obj));
    }

    @Override
    public boolean chk(Serializable obj, String code) throws IOException {
        return password.checkPassword(obj2str(obj), code);
    }

    @Override
    public String encode(Serializable obj) throws IOException {
        return text.encrypt(obj2str(obj));
    }

    @Override
    public <T extends Serializable> T decode(String code, Class<T> clazz) throws IOException {
        return str2obj(text.decrypt(code), clazz);
    }


    @Override
    public String obj2str(Serializable obj) throws IOException {
        return mapper.writeValueAsString(obj);
    }

    @Override
    public <T extends Serializable> T str2obj(String code, Class<T> clazz) throws IOException {
        return mapper.readValue(code, clazz);
    }


    @PostConstruct
    void init() {
        mapper = new ObjectMapper();
        password = new StrongPasswordEncryptor();
        StrongTextEncryptor ste = new StrongTextEncryptor();
        ste.setPassword(secret);
        this.text = ste;
    }

    private ObjectMapper mapper;

    private PasswordEncryptor password;
    private TextEncryptor text;
    @Value("${app.secret}")
    private String secret;

    public void setSecret(String secret) {
        this.secret = secret;
    }
}
