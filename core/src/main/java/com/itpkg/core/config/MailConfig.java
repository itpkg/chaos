package com.itpkg.core.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.JavaMailSenderImpl;

/**
 * Created by flamen on 16-6-1.
 */
@Configuration
public class MailConfig {
    @Bean
    JavaMailSender mailSender() {
        JavaMailSenderImpl sender = new JavaMailSenderImpl();
        sender.setHost(host);
        sender.setPort(port);
        sender.setUsername(username);
        sender.setPassword(password);
        return sender;
    }

    @Value("${mail.host}")
    String host;
    @Value("${mail.port}")
    int port;
    @Value("${mail.username}")
    String username;
    @Value("${mail.password}")
    String password;

}
