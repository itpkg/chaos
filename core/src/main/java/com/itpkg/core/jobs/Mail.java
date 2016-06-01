package com.itpkg.core.jobs;

import java.io.Serializable;
import java.util.Set;

/**
 * Created by flamen on 16-6-1.
 */
public class Mail implements Serializable {
    public String subject;
    public String to;
    public String body;
    public Set<String> files;
}
