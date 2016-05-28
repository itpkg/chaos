package com.itpkg.core.auth;

import java.lang.annotation.*;

/**
 * Created by flamen on 16-5-28.
 */
@Target(ElementType.PARAMETER)
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface CurrentUser {
}
