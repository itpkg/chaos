package com.itpkg.core.models;

import javax.persistence.*;

/**
 * Created by flamen on 16-5-27.
 */

@Entity
@Table(name = "permissions")
public class Permission extends Editable {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;
    @ManyToOne
    @JoinColumn(nullable = false, name = "user_id")
    private User user;
    private String resourceType;
    private Long resourceId;

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

    public String getResourceType() {
        return resourceType;
    }

    public void setResourceType(String resourceType) {
        this.resourceType = resourceType;
    }

    public Long getResourceId() {
        return resourceId;
    }

    public void setResourceId(Long resourceId) {
        this.resourceId = resourceId;
    }
}
