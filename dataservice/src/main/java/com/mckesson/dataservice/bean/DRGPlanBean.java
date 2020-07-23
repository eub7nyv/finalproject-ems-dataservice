package com.mckesson.dataservice.bean;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "T_DRG_PLAN_DATA")
public class DRGPlanBean {

    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    private int ID;
    private String payerId;
    private String planId;
    private String benifitOfferringName;
    private String segment;
    private String planTypeName;
    private String hmo;
    private String ppo;
    private String pos;
    private Date createdDate=new Date();
    private String createdBy;
    private Date lastupdateDate=new Date();
    private String lastUpdateddBy;
    private String isActive="T";


    public DRGPlanBean() {
    }

    public int getID() {
        return ID;
    }

    public void setID(int ID) {
        this.ID = ID;
    }

   

    public String getPlanId() {
        return planId;
    }

    public void setPlanId(String planId) {
        this.planId = planId;
    }

    public String getBenifitOfferringName() {
        return benifitOfferringName;
    }

    public void setBenifitOfferringName(String benifitOfferringName) {
        this.benifitOfferringName = benifitOfferringName;
    }

    public String getSegment() {
        return segment;
    }

    public void setSegment(String segment) {
        this.segment = segment;
    }

    public String getPlanTypeName() {
        return planTypeName;
    }

    public void setPlanTypeName(String planTypeName) {
        this.planTypeName = planTypeName;
    }

    public String getHmo() {
        return hmo;
    }

    public void setHmo(String hmo) {
        this.hmo = hmo;
    }

    public String getPos() {
        return pos;
    }

    public void setPos(String pos) {
        this.pos = pos;
    }

    public Date getCreatedDate() {
        return createdDate;
    }

    public void setCreatedDate(Date createdDate) {
        this.createdDate = createdDate;
    }

    public String getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(String createdBy) {
        this.createdBy = createdBy;
    }

    public Date getLastupdateDate() {
        return lastupdateDate;
    }

    public void setLastupdateDate(Date lastupdateDate) {
        this.lastupdateDate = lastupdateDate;
    }

    public String getLastUpdateddBy() {
        return lastUpdateddBy;
    }

    public void setLastUpdateddBy(String lastUpdateddBy) {
        this.lastUpdateddBy = lastUpdateddBy;
    }

    public String getIsActive() {
        return isActive;
    }

    public void setIsActive(String isActive) {
        this.isActive = isActive;
    }

	public String getPayerId() {
		return payerId;
	}

	public void setPayerId(String payerId) {
		this.payerId = payerId;
	}

	public String getPpo() {
		return ppo;
	}

	public void setPpo(String ppo) {
		this.ppo = ppo;
	}
}
