package com.mckesson.dataservice.entity;


import org.springframework.boot.autoconfigure.domain.EntityScan;



@EntityScan
public class Message {

	private String appName;
    private String incomingMessage;
    
    
	public String getAppName() {
		return appName;
	}
	public void setAppName(String appName) {
		this.appName = appName;
	}
	public String getIncomingMessage() {
		return incomingMessage;
	}
	public void setIncomingMessage(String incomingMessage) {
		this.incomingMessage = incomingMessage;
	}
    
    
    
}
