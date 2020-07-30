package com.mckesson.producerdataservice.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * @author anoopunnikrishnan
 *
 */

public class MessageDTO {

    private Long id;
    @JsonProperty("application_name")
    private String applicationName;
    @JsonProperty("incoming_message")
    private String incomingMessage;
    @JsonProperty("send_status")
    private boolean sendStatus;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getIncomingMessage() {
        return incomingMessage;
    }

    public void setIncomingMessage(String incomingMessage) {
        this.incomingMessage = incomingMessage;
    }

    public String getApplicationName() {
        return applicationName;
    }

    public void setApplicationName(String applicationName) {
        this.applicationName = applicationName;
    }

    public boolean isSendStatus() {
        return sendStatus;
    }

    public void setSendStatus(boolean sendStatus) {
        this.sendStatus = sendStatus;
    }

}