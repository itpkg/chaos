package com.itpkg.core.web;

import java.io.Serializable;

/**
 * Created by flamen on 16-5-30.
 */
public class Link implements Serializable {
    public Link(String href, String label) {
        this.href = href;
        this.label = label;
    }

    private String href;
    private String label;

    public String getHref() {
        return href;
    }

    public void setHref(String href) {
        this.href = href;
    }

    public String getLabel() {
        return label;
    }

    public void setLabel(String label) {
        this.label = label;
    }
}
