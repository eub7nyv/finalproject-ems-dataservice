package com.mckesson.dataservice.bean;

import java.util.Date;

import javax.persistence.*;


@Entity
@Table(name = "T_DRG_PAYER_DATA")
public class DRGPayerBean {

    @Id
    @GeneratedValue(strategy= GenerationType.AUTO)
    private int ID;
    private String payer_Id;
    private String parent_Id;//parentId
    private String payer_Name;
    private Date createdDate=new Date();
    private String createdBy;
    private Date lastupdateDate=new Date();
    private String lastUpdateddBy;
    private String isActive="T";




    public DRGPayerBean() {
		// TODO Auto-generated constructor stub
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
	 * @return the parent_Id
	 */
	public String getParent_Id() {
		return parent_Id;
	}




	/**
	 * @param parent_Id the parent_Id to set
	 */
	public void setParent_Id(String parent_Id) {
		this.parent_Id = parent_Id;
	}




	/**
	 * @return the payer_Name
	 */
	public String getPayer_Name() {
		return payer_Name;
	}




	/**
	 * @param payer_Name the payer_Name to set
	 */
	public void setPayer_Name(String payer_Name) {
		this.payer_Name = payer_Name;
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
