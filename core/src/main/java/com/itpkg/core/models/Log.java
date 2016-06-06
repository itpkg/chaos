package com.itpkg.core.models;

import org.hibernate.annotations.DynamicInsert;

import javax.persistence.*;
import java.io.Serializable;
import java.util.Date;

/**
 * Created by flamen on 16-5-27.
 */
@Entity
@Table(name = "logs")
@DynamicInsert
public class Log implements Serializable {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;
    @ManyToOne
    @JoinColumn(nullable = false, name = "user_id")
    private User user;
    @Column(nullable = false)
    private String message;
    @Column(nullable = false)
    private Date createdAt;

    @PrePersist
    void createdAt() {
        this.createdAt = new Date();
    }


    public Date getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Date createdAt) {
        this.createdAt = createdAt;
    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

}
