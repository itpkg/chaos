package com.itpkg.core.jobs;

import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.listener.PatternTopic;
import org.springframework.data.redis.listener.RedisMessageListenerContainer;
import org.springframework.data.redis.listener.adapter.MessageListenerAdapter;
import org.springframework.data.redis.serializer.Jackson2JsonRedisSerializer;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;

/**
 * Created by flamen on 16-6-1.
 */
@Component
public class EmailSender {
    public final static String CHANNEL = "emails";

    public void send(Mail mail) {
        template.convertAndSend(CHANNEL, mail);
    }

    @PostConstruct
    void init() {
        emailAdapter.setSerializer(new Jackson2JsonRedisSerializer<>(Mail.class));
        container.addMessageListener(emailAdapter, new PatternTopic(CHANNEL));
    }

    @Resource
    RedisMessageListenerContainer container;
    @Resource
    RedisTemplate<String, Mail> template;
    @Resource(name = "core.emailMessageListenerAdapter")
    MessageListenerAdapter emailAdapter;

}
