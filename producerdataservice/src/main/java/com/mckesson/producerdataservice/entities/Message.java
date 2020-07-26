package com.mckesson.producerdataservive.entities;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

/**
 * @author anoopunnikrishnan
 *
 */

@Entity
@Table(name = "T_INCOMING_MESSAGES")
public class Message {
    
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;
    private String applicationName;
    private String incomingMessage;
    private boolean sendStatus;

    public Message(String applicationName, String incomingMessage, boolean sendStatus) {
        this.applicationName = applicationName;
        this.incomingMessage = incomingMessage;
        this.sendStatus = sendStatus;
    }

    public Message() {
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getApplicationName() {
        return applicationName;
    }

    public void setApplicationName(String applicationName) {
        this.applicationName = applicationName;
    }

    public String getIncomingMessage() {
        return incomingMessage;
    }

    public void setIncomingMessage(String incomingMessage) {
        this.incomingMessage = incomingMessage;
    }

    public boolean isSendStatus() {
        return sendStatus;
    }

    public void setSendStatus(boolean sendStatus) {
        this.sendStatus = sendStatus;
    }

}