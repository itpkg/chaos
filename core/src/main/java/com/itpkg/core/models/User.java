package com.itpkg.core.models;

import javax.persistence.*;
import java.util.*;

/**
 * Created by flamen on 16-5-27.
 */
@Entity
@Table(name = "users",
        indexes = {
                @Index(columnList = "providerType,providerId", unique = true),
                @Index(columnList = "providerType"),
                @Index(columnList = "name")
        }
)
public class User extends Editable implements ToModel {
    @Override
    public Map<String, Object> toModel() {
        Map<String, Object> map = new HashMap<>();
        map.put("uid", uid);
        map.put("name", name);
        map.put("email", email);
        return map;
    }

    public enum Type {
        EMAIL, GOOGLE, QQ, WE_CHAT, FACEBOOK
    }

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;
    @Column(unique = true, nullable = false)
    private String uid;
    @Column(nullable = false)
    private String name;
    @Column(nullable = false, unique = true)
    private String email;
    private String password;
    @Column(nullable = false)
    private String providerId;
    @Column(nullable = false)
    @Enumerated(EnumType.STRING)
    private Type providerType;
    private Date confirmedAt;
    private Date lockedAt;

    @OneToMany(cascade = {CascadeType.ALL}, mappedBy = "user", fetch = FetchType.LAZY)
    private List<Log> logs;
    @OneToMany(mappedBy = "user", fetch = FetchType.LAZY)
    private List<Permission> permissions;



    public Date getConfirmedAt() {
        return confirmedAt;
    }

    public void setConfirmedAt(Date confirmedAt) {
        this.confirmedAt = confirmedAt;
    }

    public Date getLockedAt() {
        return lockedAt;
    }

    public void setLockedAt(Date lockedAt) {
        this.lockedAt = lockedAt;
    }

    public User() {
        this.logs = new ArrayList<>();
        this.permissions = new ArrayList<>();
    }

    public List<Permission> getPermissions() {
        return permissions;
    }

    public void setPermissions(List<Permission> permissions) {
        this.permissions = permissions;
    }

    public List<Log> getLogs() {
        return logs;
    }

    public void setLogs(List<Log> logs) {
        this.logs = logs;
    }

    public String getProviderId() {
        return providerId;
    }

    public void setProviderId(String providerId) {
        this.providerId = providerId;
    }

    public Type getProviderType() {
        return providerType;
    }

    public void setProviderType(Type providerType) {
        this.providerType = providerType;
    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getUid() {
        return uid;
    }

    public void setUid(String uid) {
        this.uid = uid;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}
