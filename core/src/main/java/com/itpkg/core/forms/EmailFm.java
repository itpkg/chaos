package com.itpkg.core.forms;

import java.io.Serializable;

/**
 * Created by flamen on 16-6-1.
 */
public class EmailFm implements Serializable {
    private String email;

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
}
