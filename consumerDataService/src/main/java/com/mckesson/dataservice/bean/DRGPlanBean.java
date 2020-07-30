package com.mckesson.dataservice.bean;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "T_DRG_PLAN_DATA")
public class DRGPlanBean {

    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    private int ID;
    private String payer_Id;
    private String plan_Id;
    private String plan_name;
    private String segment;
    private String plan_Type_Name;
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


	/**
	 * @return the iD
	 */
	public int getID() {
		return ID;
	}


	/**
	 * @param iD the iD to set
	 */
	public void setID(int iD) {
		ID = iD;
	}


	/**
	 * @return the payer_Id
	 */
	public String getPayer_Id() {
		return payer_Id;
	}


	/**
	 * @param payer_Id the payer_Id to set
	 */
	public void setPayer_Id(String payer_Id) {
		this.payer_Id = payer_Id;
	}


	/**
	 * @return the plan_Id
	 */
	public String getPlan_Id() {
		return plan_Id;
	}


	/**
	 * @param plan_Id the plan_Id to set
	 */
	public void setPlan_Id(String plan_Id) {
		this.plan_Id = plan_Id;
	}


	/**
	 * @return the plan_name
	 */
	public String getPlan_name() {
		return plan_name;
	}


	/**
	 * @param plan_name the plan_name to set
	 */
	public void setPlan_name(String plan_name) {
		this.plan_name = plan_name;
	}


	/**
	 * @return the segment
	 */
	public String getSegment() {
		return segment;
	}


	/**
	 * @param segment the segment to set
	 */
	public void setSegment(String segment) {
		this.segment = segment;
	}


	/**
	 * @return the plan_Type_Name
	 */
	public String getPlan_Type_Name() {
		return plan_Type_Name;
	}


	/**
	 * @param plan_Type_Name the plan_Type_Name to set
	 */
	public void setPlan_Type_Name(String plan_Type_Name) {
		this.plan_Type_Name = plan_Type_Name;
	}


	/**
	 * @return the hmo
	 */
	public String getHmo() {
		return hmo;
	}


	/**
	 * @param hmo the hmo to set
	 */
	public void setHmo(String hmo) {
		this.hmo = hmo;
	}


	/**
	 * @return the ppo
	 */
	public String getPpo() {
		return ppo;
	}


	/**
	 * @param ppo the ppo to set
	 */
	public void setPpo(String ppo) {
		this.ppo = ppo;
	}


	/**
	 * @return the pos
	 */
	public String getPos() {
		return pos;
	}


	/**
	 * @param pos the pos to set
	 */
	public void setPos(String pos) {
		this.pos = pos;
	}


	/**
	 * @return the createdDate
	 */
	public Date getCreatedDate() {
		return createdDate;
	}


	/**
	 * @param createdDate the createdDate to set
	 */
	public void setCreatedDate(Date createdDate) {
		this.createdDate = createdDate;
	}


	/**
	 * @return the createdBy
	 */
	public String getCreatedBy() {
		return createdBy;
	}


	/**
	 * @param createdBy the createdBy to set
	 */
	public void setCreatedBy(String createdBy) {
		this.createdBy = createdBy;
	}


	/**
	 * @return the lastupdateDate
	 */
	public Date getLastupdateDate() {
		return lastupdateDate;
	}


	/**
	 * @param lastupdateDate the lastupdateDate to set
	 */
	public void setLastupdateDate(Date lastupdateDate) {
		this.lastupdateDate = lastupdateDate;
	}


	/**
	 * @return the lastUpdateddBy
	 */
	public String getLastUpdateddBy() {
		return lastUpdateddBy;
	}


	/**
	 * @param lastUpdateddBy the lastUpdateddBy to set
	 */
	public void setLastUpdateddBy(String lastUpdateddBy) {
		this.lastUpdateddBy = lastUpdateddBy;
	}


	/**
	 * @return the isActive
	 */
	public String getIsActive() {
		return isActive;
	}


	/**
	 * @param isActive the isActive to set
	 */
	public void setIsActive(String isActive) {
		this.isActive = isActive;
	}

    
}
