package com.itpkg.core.services;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Date;

/**
 * Created by flamen on 16-5-27.
 */
@Service
public class PermissionService {
    public void allow(long userId, String operation, Date begin, Date end) {
        allow(userId, operation, null, null, begin, end);
    }

    public void allow(long userId, String operation, String resourceType, Date begin, Date end) {
        allow(userId, operation, resourceType, null, begin, end);
    }

    public void allow(long userId, String operation, String resourceType, Long resourceId, Date begin, Date end) {
        set(userId, operation, resourceType, resourceId, begin, end, true);
    }

    public void deny(long userId, String operation) {
        set(userId, operation, null, null, null, null, false);
    }

    public void deny(long userId, String operation, String resourceType) {
        set(userId, operation, resourceType, null, null, null, false);
    }

    public void deny(long userId, String operation, String resourceType, Long resourceId) {
        set(userId, operation, resourceType, resourceId, null, null, false);
    }

    public boolean can(long userId, String operation) {
        return can(userId, operation, null, null);
    }

    public boolean can(long userId, String operation, String resourceType) {
        return can(userId, operation, resourceType, null);
    }

    public boolean can(long userId, String operation, String resourceType, Long resourceId) {
        //todo
        return false;
    }

    @Transactional
    private void set(long userId, String operation, String resourceType, Long resourceId, Date begin, Date end, boolean allow) {
        //todo
    }

}
