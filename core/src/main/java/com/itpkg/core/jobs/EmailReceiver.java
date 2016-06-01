package com.itpkg.core.jobs;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.io.FileSystemResource;
import org.springframework.mail.MailSender;
import org.springframework.mail.javamail.JavaMailSenderImpl;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import javax.mail.MessagingException;
import javax.mail.internet.MimeMessage;
import java.io.File;
import java.util.concurrent.CountDownLatch;

/**
 * Created by flamen on 16-6-1.
 */
@Component
public class EmailReceiver {
    public void handleMessage(Mail mail) throws MessagingException {


//            SimpleMailMessage msg = new SimpleMailMessage();
//            msg.setFrom(from);
//            msg.setBcc(bcc);
//
//            msg.setTo(mail.to);
//            msg.setSubject(mail.subject);
//            msg.setText(mail.body);
//
//            mailSender.send(msg);

        JavaMailSenderImpl sender = (JavaMailSenderImpl) mailSender;
        MimeMessage msg = sender.createMimeMessage();
        MimeMessageHelper helper = new MimeMessageHelper(msg, true);
        helper.setFrom(from);
        helper.setBcc(bcc);

        helper.setTo(mail.to);
        helper.setSubject(mail.subject);
        helper.setText(mail.body, true);

        if (mail.files != null) {
            for (String f : mail.files) {
                helper.addInline(f, new FileSystemResource(new File(f)));
            }
        }

        if (debug) {
            logger.debug("=== SEND MAIL({}) {} ===\n{}\n{}", mail.to, mail.subject, mail.body);
        } else {
            sender.send(msg);
        }


        latch.countDown();
    }

    @Resource
    CountDownLatch latch;
    @Resource
    MailSender mailSender;
    @Value("${mail.from}")
    String from;
    @Value("${mail.bcc}")
    String bcc;
    @Value("${mail.debug}")
    boolean debug;
    private final static Logger logger = LoggerFactory.getLogger(EmailReceiver.class);
}
