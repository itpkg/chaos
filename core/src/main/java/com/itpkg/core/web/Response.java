package com.itpkg.core.web;

import java.io.Serializable;

/**
 * Created by flamen on 16-5-31.
 */
public class Response implements Serializable {

    private boolean ok;
    private String error;
    private Object data;

    public Object getData() {
        return data;
    }

    public void setData(Object data) {
        this.data = data;
    }

    public boolean isOk() {
        return ok;
    }

    public void setOk(boolean ok) {
        this.ok = ok;
    }

    public String getError() {
        return error;
    }

    public void setError(String error) {
        this.error = error;
    }
}
