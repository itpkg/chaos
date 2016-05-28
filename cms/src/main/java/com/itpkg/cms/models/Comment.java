package com.itpkg.cms.models;

import javax.persistence.*;
import java.io.Serializable;

/**
 * Created by flamen on 16-5-27.
 */

@Entity
@Table(name = "cms_comments")
public class Comment implements Serializable {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;
    @Column(nullable = false)
    private String body;
    @ManyToOne
    @JoinColumn(nullable = false, name = "article_id")
    private Article article;

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public String getBody() {
        return body;
    }

    public void setBody(String body) {
        this.body = body;
    }

    public Article getArticle() {
        return article;
    }

    public void setArticle(Article article) {
        this.article = article;
    }
}
