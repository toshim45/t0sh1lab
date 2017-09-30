package com.talentmgr.api.domain.profile;

import java.util.Date;

/**
 * Created by artikow on 9/30/17.
 */

public class Profile {
    private String description;
    private Date born;

    public Profile(String description, Date born) {
        this.description = description;
        this.born = born;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public Date getBorn() {
        return born;
    }

    public void setBorn(Date born) {
        this.born = born;
    }
}
