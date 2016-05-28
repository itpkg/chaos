package com.itpkg.core.models;

import javax.persistence.*;

/**
 * Created by flamen on 16-5-27.
 */
@Entity
@Table(name = "settings")
public class Setting extends Editable {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;
    @Column(nullable = false, unique = true)
    private String key;
    @Column(nullable = false)
    private String val;

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public String getKey() {
        return key;
    }

    public void setKey(String key) {
        this.key = key;
    }

    public String getVal() {
        return val;
    }

    public void setVal(String val) {
        this.val = val;
    }
}
