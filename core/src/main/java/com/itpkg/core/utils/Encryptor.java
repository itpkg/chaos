package com.itpkg.core.utils;

import java.io.IOException;
import java.io.Serializable;

/**
 * Created by flamen on 16-5-27.
 */
public interface Encryptor {
    String sum(Serializable obj) throws IOException;

    boolean chk(Serializable obj, String code) throws IOException;

    String encode(Serializable obj) throws IOException;

    <T extends Serializable> T decode(String code, Class<T> clazz) throws IOException;

    String obj2str(Serializable obj) throws IOException;

    <T extends Serializable> T str2obj(String code, Class<T> clazz) throws IOException;
}
