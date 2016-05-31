package com.itpkg.core.utils;

/**
 * Created by flamen on 16-5-27.
 */
public interface Encryptor {
    String sum(String plain);

    boolean chk(String plain, String code);

    String encode(String plain);

    String decode(String code);

}
