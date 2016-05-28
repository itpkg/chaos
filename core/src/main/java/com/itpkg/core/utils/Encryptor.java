package com.itpkg.core.utils;

import java.io.IOException;
import java.io.Serializable;

/**
 * Created by flamen on 16-5-27.
 */
public interface Encryptor {
    <T extends Serializable> String sum(T obj) throws IOException;

    <T extends Serializable> boolean chk(T obj, String code) throws IOException;

    <T extends Serializable> String encode(T obj) throws IOException;

    <T extends Serializable> T decode(String code, Class<T> clazz) throws IOException;

    <T extends Serializable> String obj2str(T obj) throws IOException;

    <T extends Serializable> T str2obj(String code, Class<T> clazz) throws IOException;
}
