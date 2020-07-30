/**
 * 
 */
package com.mckesson.dataservice.bean;

import java.util.Date;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

/**
 * @author jakeermahamad
 *
 */
@Entity
@Table(name = "T_NCPDP_DATA")
public class NCPDPBean {
	
	@Id
    @GeneratedValue(strategy= GenerationType.AUTO)
    private int ID;
    private String pharmacy_Id;//PHARMACY_ID
    private String pharmacy_Name;//PHARMACY_NAME
    private String ncpdp;//
    private String npi_Nbr;//NPI_NBR
    private String phone_Nbr;//PHONE_NBR
    private String fax_Nbr;//FAX_NBR
    private Date createdDate=new Date();
    private String createdBy;
    private Date lastupdateDate=new Date();
    private String lastUpdateddBy;
    private String isActive="T";
    
    
    
	public NCPDPBean() {
		 
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
	 * @return the pharmacy_Id
	 */
	public String getPharmacy_Id() {
		return pharmacy_Id;
	}



	/**
	 * @param pharmacy_Id the pharmacy_Id to set
	 */
	public void setPharmacy_Id(String pharmacy_Id) {
		this.pharmacy_Id = pharmacy_Id;
	}



	/**
	 * @return the pharmacy_Name
	 */
	public String getPharmacy_Name() {
		return pharmacy_Name;
	}



	/**
	 * @param pharmacy_Name the pharmacy_Name to set
	 */
	public void setPharmacy_Name(String pharmacy_Name) {
		this.pharmacy_Name = pharmacy_Name;
	}



	/**
	 * @return the ncpdp
	 */
	public String getNcpdp() {
		return ncpdp;
	}



	/**
	 * @param ncpdp the ncpdp to set
	 */
	public void setNcpdp(String ncpdp) {
		this.ncpdp = ncpdp;
	}



	/**
	 * @return the npi_Nbr
	 */
	public String getNpi_Nbr() {
		return npi_Nbr;
	}



	/**
	 * @param npi_Nbr the npi_Nbr to set
	 */
	public void setNpi_Nbr(String npi_Nbr) {
		this.npi_Nbr = npi_Nbr;
	}



	/**
	 * @return the phone_Nbr
	 */
	public String getPhone_Nbr() {
		return phone_Nbr;
	}



	/**
	 * @param phone_Nbr the phone_Nbr to set
	 */
	public void setPhone_Nbr(String phone_Nbr) {
		this.phone_Nbr = phone_Nbr;
	}



	/**
	 * @return the fax_Nbr
	 */
	public String getFax_Nbr() {
		return fax_Nbr;
	}



	/**
	 * @param fax_Nbr the fax_Nbr to set
	 */
	public void setFax_Nbr(String fax_Nbr) {
		this.fax_Nbr = fax_Nbr;
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
