package com.itpkg.core.config;

import com.itpkg.core.jobs.EmailReceiver;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cache.CacheManager;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.cache.RedisCacheManager;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.connection.jedis.JedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.listener.RedisMessageListenerContainer;
import org.springframework.data.redis.listener.adapter.MessageListenerAdapter;
import org.springframework.data.redis.serializer.Jackson2JsonRedisSerializer;
import org.springframework.data.redis.serializer.StringRedisSerializer;

import java.util.concurrent.CountDownLatch;

/**
 * Created by flamen on 16-6-1.
 */

@Configuration
public class RedisConfig {
    @Bean
    CacheManager cacheManager(RedisTemplate redisTemplate) {
        RedisCacheManager manager = new RedisCacheManager(redisTemplate);
        manager.setDefaultExpiration(cacheExpiration);
        return manager;
    }

    @Bean
    JedisConnectionFactory redisConnectionFactory() {
        JedisConnectionFactory factory = new JedisConnectionFactory();

        factory.setHostName(redisHostname);
        factory.setDatabase(redisDatabase);
        factory.setPort(redisPort);
        factory.setUsePool(true);
        return factory;
    }

    @Bean
    RedisMessageListenerContainer container(RedisConnectionFactory connectionFactory) {

        RedisMessageListenerContainer container = new RedisMessageListenerContainer();
        container.setConnectionFactory(connectionFactory);

        return container;
    }

    @Bean(name = "core.emailMessageListenerAdapter")
    MessageListenerAdapter messageListenerAdapter(EmailReceiver emailReceiver) {
        //return new MessageListenerAdapter(emailReceiver, "receiveMessage");
        return new MessageListenerAdapter(emailReceiver);
    }

    @Bean
    CountDownLatch latch() {
        return new CountDownLatch(1);
    }

    @Bean
    RedisTemplate redisTemplate(RedisConnectionFactory connectionFactory) {
        RedisTemplate<String, Object> t = new RedisTemplate<>();
        t.setConnectionFactory(connectionFactory);
        t.setKeySerializer(new StringRedisSerializer());
        t.setHashKeySerializer(new StringRedisSerializer());
        t.setValueSerializer(new Jackson2JsonRedisSerializer<>(Object.class));
        t.setHashValueSerializer(new Jackson2JsonRedisSerializer<>(Object.class));
        return t;
    }

    @Value("${redis.hostname}")
    String redisHostname;
    @Value("${redis.port}")
    int redisPort;
    @Value("${redis.database}")
    int redisDatabase;
    @Value("${cache.expiration}")
    long cacheExpiration;
}
