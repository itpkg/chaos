package com.itpkg.core.web;

import java.io.Serializable;

/**
 * Created by flamen on 16-5-31.
 */
public class StringResponse implements Serializable {
    public StringResponse(String response) {
        this.response = response;
    }

    private String response;

    public String getResponse() {
        return response;
    }

    public void setResponse(String response) {
        this.response = response;
    }
}
