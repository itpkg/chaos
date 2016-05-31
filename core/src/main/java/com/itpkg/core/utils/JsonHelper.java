package com.itpkg.core.utils;

import java.io.IOException;
import java.io.Serializable;
import java.util.List;

/**
 * Created by flamen on 16-5-31.
 */
public interface JsonHelper {
    String obj2str(Serializable obj) throws IOException;

    <T extends Serializable> T str2obj(String code, Class<T> clazz) throws IOException;

    <T extends Serializable> List<T> str2lst(String s, Class<T> clazz) throws IOException;
}
