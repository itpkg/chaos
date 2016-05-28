package com.itpkg.core.utils;

import java.io.Serializable;

/**
 * Created by flamen on 16-5-27.
 */
public interface Encrypt {
    String sum(Object obj);

    boolean chk(Object obj, String code);

    String encode(Object obj);

    <T extends Serializable> T decode(String code, Class<T> clazz);
}
