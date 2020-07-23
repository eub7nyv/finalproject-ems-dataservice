package com.mckesson.dataservice.bean;

import java.util.Date;

import javax.persistence.*;


@Entity
@Table(name = "T_DRG_PAYER_DATA")
public class DRGPayerBean {

    @Id
    @GeneratedValue(strategy= GenerationType.AUTO)
    private int ID;
    private String payerId;
    private String parentid;//parentId
    private String payerName;
    private Date createdDate=new Date();
    private String createdBy;
    private Date lastupdateDate=new Date();
    private String lastUpdateddBy;
    private String isActive="T";




    public DRGPayerBean() {
		// TODO Auto-generated constructor stub
	}

	public int getID() {
        return ID;
    }

    public void setID(int ID) {
        this.ID = ID;
    }

    public String getPayerId() {
        return payerId;
    }

    public void setPayerId(String payerId) {
        this.payerId = payerId;
    }

    public String getParentid() {
        return parentid;
    }

    public void setParentid(String parentid) {
        this.parentid = parentid;
    }

    public String getPayerName() {
        return payerName;
    }

    public void setPayerName(String payerName) {
        this.payerName = payerName;
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
}
