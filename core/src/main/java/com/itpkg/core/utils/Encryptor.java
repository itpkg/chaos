package com.itpkg.core.utils;

import java.io.IOException;
import java.io.Serializable;

/**
 * Created by flamen on 16-5-27.
 */
public interface Encryptor {
    String sum(Object obj) throws IOException;

    boolean chk(Object obj, String code) throws IOException;

    String encode(Object obj) throws IOException;

    <T extends Serializable> T decode(String code, Class<T> clazz) throws IOException;
}
