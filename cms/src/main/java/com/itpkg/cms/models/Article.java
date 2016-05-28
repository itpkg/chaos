package com.itpkg.cms.models;

import com.itpkg.core.models.User;

import javax.persistence.*;
import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by flamen on 16-5-27.
 */

@Entity
@Table(name = "cms_articles")
public class Article implements Serializable {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;
    @Column(nullable = false)
    private String title;
    @Column(nullable = false)
    private String summary;
    @Column(nullable = false)
    @Lob
    private String body;
    @ManyToMany(mappedBy = "articles")
    private List<Tag> tags;
    @OneToMany(cascade = {CascadeType.ALL}, mappedBy = "article")
    private List<Comment> comments;
    @ManyToOne
    @JoinColumn(nullable = false, name = "user_id")
    private User user;

    public Article() {
        this.tags = new ArrayList<>();
        this.comments = new ArrayList<>();
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    public List<Comment> getComments() {
        return comments;
    }

    public void setComments(List<Comment> comments) {
        this.comments = comments;
    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getSummary() {
        return summary;
    }

    public void setSummary(String summary) {
        this.summary = summary;
    }

    public String getBody() {
        return body;
    }

    public void setBody(String body) {
        this.body = body;
    }

    public List<Tag> getTags() {
        return tags;
    }

    public void setTags(List<Tag> tags) {
        this.tags = tags;
    }
}
